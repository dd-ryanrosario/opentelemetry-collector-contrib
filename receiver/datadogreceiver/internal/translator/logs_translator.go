// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator

import (
	"fmt"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

// DatadogLog represents the structure of a Datadog log
type DatadogLog struct {
	Message     string                 `json:"message"`
	Status      string                 `json:"status"`
	Service     string                 `json:"service"`
	Timestamp   int64                  `json:"timestamp"`
	Tags        []string              `json:"tags"`
	Attributes  map[string]interface{} `json:"attributes"`
	Host        string                 `json:"host"`
	Source      string                 `json:"source"`
}

type LogsTranslator struct {
	buildInfo  component.BuildInfo
	stringPool *StringPool
}

func NewLogsTranslator(buildInfo component.BuildInfo) *LogsTranslator {
	return &LogsTranslator{
		buildInfo:  buildInfo,
		stringPool: newStringPool(),
	}
}

// mapSeverity converts Datadog severity to OpenTelemetry severity
func (lt *LogsTranslator) mapSeverity(status string) (plog.SeverityNumber, string) {
	switch status {
	case "emergency":
		return plog.SeverityNumberFatal, status
	case "error":
		return plog.SeverityNumberError, status
	case "warn":
		return plog.SeverityNumberWarn, status
	case "info":
		return plog.SeverityNumberInfo, status
	case "debug":
		return plog.SeverityNumberDebug, status
	default:
		return plog.SeverityNumberUnspecified, status
	}
}

// TranslateLogs converts Datadog logs to OpenTelemetry log format
func (lt *LogsTranslator) TranslateLogs(ddLogs []DatadogLog) (plog.Logs, error) {
	logs := plog.NewLogs()
	resourceLogs := logs.ResourceLogs().AppendEmpty()
	scopeLogs := resourceLogs.ScopeLogs().AppendEmpty()

	// Set resource attributes that are common across all logs
	resourceAttrs := resourceLogs.Resource().Attributes()
	resourceAttrs.PutStr("service.name", "datadog-receiver")
	resourceAttrs.PutStr("collector.name", lt.buildInfo.Command)
	resourceAttrs.PutStr("collector.version", lt.buildInfo.Version)

	for _, ddLog := range ddLogs {
		logRecord := scopeLogs.LogRecords().AppendEmpty()

		// Set timestamp
		if ddLog.Timestamp > 0 {
			logRecord.SetTimestamp(pcommon.NewTimestampFromTime(time.Unix(0, ddLog.Timestamp)))
		} else {
			logRecord.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
		}

		// Set severity
		severityNumber, severityText := lt.mapSeverity(ddLog.Status)
		logRecord.SetSeverityNumber(severityNumber)
		logRecord.SetSeverityText(severityText)

		// Set log body
		logRecord.Body().SetStr(ddLog.Message)

		// Set attributes
		attrs := logRecord.Attributes()
		if ddLog.Host != "" {
			attrs.PutStr("host.name", ddLog.Host)
		}
		if ddLog.Source != "" {
			attrs.PutStr("source", ddLog.Source)
		}
		if ddLog.Service != "" {
			attrs.PutStr("service", ddLog.Service)
		}

		// Add tags as attributes
		for _, tag := range ddLog.Tags {
			// In a real implementation, you'd want to properly parse k:v tags
			attrs.PutStr("tag", tag)
		}

		// Add additional attributes
		for k, v := range ddLog.Attributes {
			// In a real implementation, you'd want to properly type assert and handle different value types
			if strVal, ok := v.(string); ok {
				attrs.PutStr(k, strVal)
			}
		}
	}

	return logs, nil
} 