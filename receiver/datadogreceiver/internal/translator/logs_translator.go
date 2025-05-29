// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

// DatadogLog represents the structure of a Datadog log
type DatadogLog struct {
	Message    string                 `json:"message"`
	Status     string                 `json:"status"`
	Service    string                 `json:"service"`
	Timestamp  int64                  `json:"timestamp"`
	Tags       []string               `json:"tags"`
	Attributes map[string]interface{} `json:"attributes"`
	Host       string                 `json:"host"`
	Source     string                 `json:"source"`
}

type LogsTranslator struct {
	buildInfo component.BuildInfo
}

func NewLogsTranslator(buildInfo component.BuildInfo) *LogsTranslator {
	return &LogsTranslator{buildInfo: buildInfo}
}

// TranslateLogs converts Datadog logs to OpenTelemetry log format
func (lt *LogsTranslator) TranslateLogs(ddLogs []DatadogLog) (plog.Logs, error) {
	// Root container for OTLP logs
	// Resource -> Scope -> Log Record
	logs := plog.NewLogs()
	groups := groupLogsByResource(ddLogs)

	for _, group := range groups {
		rl := logs.ResourceLogs().AppendEmpty()
		setResourceAttributes(rl.Resource().Attributes(), group[0], lt.buildInfo)

		scopeLogs := rl.ScopeLogs().AppendEmpty()
		scopeLogs.Scope().SetName("datadog.receiver")
		scopeLogs.Scope().SetVersion(lt.buildInfo.Version)

		for _, ddLog := range group {
			record := scopeLogs.LogRecords().AppendEmpty()
			createLogRecord(ddLog, record)
		}
	}
	return logs, nil
}

func groupLogsByResource(ddLogs []DatadogLog) map[string][]DatadogLog {
	grouped := make(map[string][]DatadogLog)
	for _, log := range ddLogs {
		key := fmt.Sprintf("host=%s|svc=%s", log.Host, log.Service)
		grouped[key] = append(grouped[key], log)
	}
	return grouped
}

func setResourceAttributes(attrs pcommon.Map, log DatadogLog, buildInfo component.BuildInfo) {
	if log.Host != "" {
		attrs.PutStr("host.name", log.Host)
	}
	if log.Service != "" {
		attrs.PutStr("service.name", log.Service)
	}
	attrs.PutStr("collector.name", buildInfo.Command)
	attrs.PutStr("collector.version", buildInfo.Version)
}

func createLogRecord(ddLog DatadogLog, record plog.LogRecord) {
	if ddLog.Timestamp > 0 {
		record.SetTimestamp(pcommon.NewTimestampFromTime(time.Unix(0, ddLog.Timestamp)))
	} else {
		record.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	}

	// Severity
	severityNumber, severityText := mapSeverity(ddLog.Status)
	record.SetSeverityNumber(severityNumber)
	record.SetSeverityText(severityText)

	// Body
	record.Body().SetStr(ddLog.Message)

	// Attributes
	attrs := record.Attributes()
	if ddLog.Source != "" {
		attrs.PutStr("source", ddLog.Source)
	}
	if ddLog.Service != "" {
		attrs.PutStr("service", ddLog.Service)
	}
	if ddLog.Host != "" {
		attrs.PutStr("host.name", ddLog.Host)
	}
	addTagsAsAttributes(attrs, ddLog.Tags)
	addDynamicAttributes(attrs, ddLog.Attributes)
	addTraceContext(attrs, &record)
}

func addTagsAsAttributes(attrs pcommon.Map, tags []string) {
	unnamedCount := 0
	for _, tag := range tags {
		parts := strings.SplitN(tag, ":", 2)
		if len(parts) == 2 {
			attrs.PutStr(parts[0], parts[1])
		} else {
			key := fmt.Sprintf("tag.%d", unnamedCount)
			attrs.PutStr(key, tag)
			unnamedCount++
		}
	}
}

func addDynamicAttributes(attrs pcommon.Map, extra map[string]interface{}) {
	for k, v := range extra {
		switch val := v.(type) {
		case string:
			attrs.PutStr(k, val)
		case bool:
			attrs.PutBool(k, val)
		case int:
			attrs.PutInt(k, int64(val))
		case int64:
			attrs.PutInt(k, val)
		case float64:
			attrs.PutDouble(k, val)
		case float32:
			attrs.PutDouble(k, float64(val))
		case uint:
			attrs.PutInt(k, int64(val))
		case json.Number:
			if i, err := val.Int64(); err == nil {
				attrs.PutInt(k, i)
			} else if f, err := val.Float64(); err == nil {
				attrs.PutDouble(k, f)
			} else {
				attrs.PutStr(k, val.String())
			}
		default:
			attrs.PutStr(k, fmt.Sprintf("[unsupported type] %T: %v", val, val))
		}
	}
}

func addTraceContext(attrs pcommon.Map, record *plog.LogRecord) {
	if val, ok := attrs.Get("dd.trace_id"); ok {
		if traceID, err := parseDecimalTraceID(val.Str()); err == nil {
			record.SetTraceID(traceID)
		}
	}
	if val, ok := attrs.Get("dd.span_id"); ok {
		if spanID, err := parseDecimalSpanID(val.Str()); err == nil {
			record.SetSpanID(spanID)
		}
	}
}

// FIND ACTUAL MAPPINGS
func mapSeverity(status string) (plog.SeverityNumber, string) {
	switch strings.ToLower(status) {
	case "emergency":
		return plog.SeverityNumberFatal4, status
	case "alert":
		return plog.SeverityNumberFatal3, status
	case "critical":
		return plog.SeverityNumberFatal2, status
	case "error":
		return plog.SeverityNumberError, status
	case "warning", "warn":
		return plog.SeverityNumberWarn, status
	case "notice":
		return plog.SeverityNumberInfo2, status
	case "info":
		return plog.SeverityNumberInfo, status
	case "debug":
		return plog.SeverityNumberDebug, status
	case "trace":
		return plog.SeverityNumberTrace, status
	default:
		return plog.SeverityNumberUnspecified, status
	}
}

func parseDecimalTraceID(s string) (pcommon.TraceID, error) {
	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return pcommon.TraceID{}, err
	}
	var traceID pcommon.TraceID
	binary.BigEndian.PutUint64(traceID[8:], id)
	return traceID, nil
}

func parseDecimalSpanID(s string) (pcommon.SpanID, error) {
	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return pcommon.SpanID{}, err
	}
	var spanID pcommon.SpanID
	binary.BigEndian.PutUint64(spanID[:], id)
	return spanID, nil
}

func (lt *LogsTranslator) FromDatadogPayload(body []byte) (plog.Logs, error) {
	var ddLogs []DatadogLog
	err := json.Unmarshal(body, &ddLogs)
	if err != nil {
		return plog.NewLogs(), fmt.Errorf("failed to unmarshal Datadog logs: %w", err)
	}
	return lt.TranslateLogs(ddLogs)
}
