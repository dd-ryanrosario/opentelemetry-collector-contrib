// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

// AttributeCommand specifies the value command attribute.
type AttributeCommand int

const (
	_ AttributeCommand = iota
	AttributeCommandGet
	AttributeCommandSet
	AttributeCommandFlush
	AttributeCommandTouch
)

// String returns the string representation of the AttributeCommand.
func (av AttributeCommand) String() string {
	switch av {
	case AttributeCommandGet:
		return "get"
	case AttributeCommandSet:
		return "set"
	case AttributeCommandFlush:
		return "flush"
	case AttributeCommandTouch:
		return "touch"
	}
	return ""
}

// MapAttributeCommand is a helper map of string to AttributeCommand attribute value.
var MapAttributeCommand = map[string]AttributeCommand{
	"get":   AttributeCommandGet,
	"set":   AttributeCommandSet,
	"flush": AttributeCommandFlush,
	"touch": AttributeCommandTouch,
}

// AttributeDirection specifies the value direction attribute.
type AttributeDirection int

const (
	_ AttributeDirection = iota
	AttributeDirectionSent
	AttributeDirectionReceived
)

// String returns the string representation of the AttributeDirection.
func (av AttributeDirection) String() string {
	switch av {
	case AttributeDirectionSent:
		return "sent"
	case AttributeDirectionReceived:
		return "received"
	}
	return ""
}

// MapAttributeDirection is a helper map of string to AttributeDirection attribute value.
var MapAttributeDirection = map[string]AttributeDirection{
	"sent":     AttributeDirectionSent,
	"received": AttributeDirectionReceived,
}

// AttributeOperation specifies the value operation attribute.
type AttributeOperation int

const (
	_ AttributeOperation = iota
	AttributeOperationIncrement
	AttributeOperationDecrement
	AttributeOperationGet
)

// String returns the string representation of the AttributeOperation.
func (av AttributeOperation) String() string {
	switch av {
	case AttributeOperationIncrement:
		return "increment"
	case AttributeOperationDecrement:
		return "decrement"
	case AttributeOperationGet:
		return "get"
	}
	return ""
}

// MapAttributeOperation is a helper map of string to AttributeOperation attribute value.
var MapAttributeOperation = map[string]AttributeOperation{
	"increment": AttributeOperationIncrement,
	"decrement": AttributeOperationDecrement,
	"get":       AttributeOperationGet,
}

// AttributeState specifies the value state attribute.
type AttributeState int

const (
	_ AttributeState = iota
	AttributeStateSystem
	AttributeStateUser
)

// String returns the string representation of the AttributeState.
func (av AttributeState) String() string {
	switch av {
	case AttributeStateSystem:
		return "system"
	case AttributeStateUser:
		return "user"
	}
	return ""
}

// MapAttributeState is a helper map of string to AttributeState attribute value.
var MapAttributeState = map[string]AttributeState{
	"system": AttributeStateSystem,
	"user":   AttributeStateUser,
}

// AttributeType specifies the value type attribute.
type AttributeType int

const (
	_ AttributeType = iota
	AttributeTypeHit
	AttributeTypeMiss
)

// String returns the string representation of the AttributeType.
func (av AttributeType) String() string {
	switch av {
	case AttributeTypeHit:
		return "hit"
	case AttributeTypeMiss:
		return "miss"
	}
	return ""
}

// MapAttributeType is a helper map of string to AttributeType attribute value.
var MapAttributeType = map[string]AttributeType{
	"hit":  AttributeTypeHit,
	"miss": AttributeTypeMiss,
}

var MetricsInfo = metricsInfo{
	MemcachedBytes: metricInfo{
		Name: "memcached.bytes",
	},
	MemcachedCommands: metricInfo{
		Name: "memcached.commands",
	},
	MemcachedConnectionsCurrent: metricInfo{
		Name: "memcached.connections.current",
	},
	MemcachedConnectionsTotal: metricInfo{
		Name: "memcached.connections.total",
	},
	MemcachedCPUUsage: metricInfo{
		Name: "memcached.cpu.usage",
	},
	MemcachedCurrentItems: metricInfo{
		Name: "memcached.current_items",
	},
	MemcachedEvictions: metricInfo{
		Name: "memcached.evictions",
	},
	MemcachedNetwork: metricInfo{
		Name: "memcached.network",
	},
	MemcachedOperationHitRatio: metricInfo{
		Name: "memcached.operation_hit_ratio",
	},
	MemcachedOperations: metricInfo{
		Name: "memcached.operations",
	},
	MemcachedThreads: metricInfo{
		Name: "memcached.threads",
	},
}

type metricsInfo struct {
	MemcachedBytes              metricInfo
	MemcachedCommands           metricInfo
	MemcachedConnectionsCurrent metricInfo
	MemcachedConnectionsTotal   metricInfo
	MemcachedCPUUsage           metricInfo
	MemcachedCurrentItems       metricInfo
	MemcachedEvictions          metricInfo
	MemcachedNetwork            metricInfo
	MemcachedOperationHitRatio  metricInfo
	MemcachedOperations         metricInfo
	MemcachedThreads            metricInfo
}

type metricInfo struct {
	Name string
}

type metricMemcachedBytes struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.bytes metric with initial data.
func (m *metricMemcachedBytes) init() {
	m.data.SetName("memcached.bytes")
	m.data.SetDescription("Current number of bytes used by this server to store items.")
	m.data.SetUnit("By")
	m.data.SetEmptyGauge()
}

func (m *metricMemcachedBytes) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedBytes) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedBytes) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedBytes(cfg MetricConfig) metricMemcachedBytes {
	m := metricMemcachedBytes{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedCommands struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.commands metric with initial data.
func (m *metricMemcachedCommands) init() {
	m.data.SetName("memcached.commands")
	m.data.SetDescription("Commands executed.")
	m.data.SetUnit("{commands}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMemcachedCommands) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, commandAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("command", commandAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedCommands) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedCommands) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedCommands(cfg MetricConfig) metricMemcachedCommands {
	m := metricMemcachedCommands{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedConnectionsCurrent struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.connections.current metric with initial data.
func (m *metricMemcachedConnectionsCurrent) init() {
	m.data.SetName("memcached.connections.current")
	m.data.SetDescription("The current number of open connections.")
	m.data.SetUnit("{connections}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricMemcachedConnectionsCurrent) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedConnectionsCurrent) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedConnectionsCurrent) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedConnectionsCurrent(cfg MetricConfig) metricMemcachedConnectionsCurrent {
	m := metricMemcachedConnectionsCurrent{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedConnectionsTotal struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.connections.total metric with initial data.
func (m *metricMemcachedConnectionsTotal) init() {
	m.data.SetName("memcached.connections.total")
	m.data.SetDescription("Total number of connections opened since the server started running.")
	m.data.SetUnit("{connections}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricMemcachedConnectionsTotal) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedConnectionsTotal) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedConnectionsTotal) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedConnectionsTotal(cfg MetricConfig) metricMemcachedConnectionsTotal {
	m := metricMemcachedConnectionsTotal{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedCPUUsage struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.cpu.usage metric with initial data.
func (m *metricMemcachedCPUUsage) init() {
	m.data.SetName("memcached.cpu.usage")
	m.data.SetDescription("Accumulated user and system time.")
	m.data.SetUnit("s")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMemcachedCPUUsage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, stateAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("state", stateAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedCPUUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedCPUUsage) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedCPUUsage(cfg MetricConfig) metricMemcachedCPUUsage {
	m := metricMemcachedCPUUsage{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedCurrentItems struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.current_items metric with initial data.
func (m *metricMemcachedCurrentItems) init() {
	m.data.SetName("memcached.current_items")
	m.data.SetDescription("Number of items currently stored in the cache.")
	m.data.SetUnit("{items}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricMemcachedCurrentItems) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedCurrentItems) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedCurrentItems) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedCurrentItems(cfg MetricConfig) metricMemcachedCurrentItems {
	m := metricMemcachedCurrentItems{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedEvictions struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.evictions metric with initial data.
func (m *metricMemcachedEvictions) init() {
	m.data.SetName("memcached.evictions")
	m.data.SetDescription("Cache item evictions.")
	m.data.SetUnit("{evictions}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricMemcachedEvictions) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedEvictions) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedEvictions) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedEvictions(cfg MetricConfig) metricMemcachedEvictions {
	m := metricMemcachedEvictions{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedNetwork struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.network metric with initial data.
func (m *metricMemcachedNetwork) init() {
	m.data.SetName("memcached.network")
	m.data.SetDescription("Bytes transferred over the network.")
	m.data.SetUnit("by")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMemcachedNetwork) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, directionAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("direction", directionAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedNetwork) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedNetwork) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedNetwork(cfg MetricConfig) metricMemcachedNetwork {
	m := metricMemcachedNetwork{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedOperationHitRatio struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.operation_hit_ratio metric with initial data.
func (m *metricMemcachedOperationHitRatio) init() {
	m.data.SetName("memcached.operation_hit_ratio")
	m.data.SetDescription("Hit ratio for operations, expressed as a percentage value between 0.0 and 100.0.")
	m.data.SetUnit("%")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMemcachedOperationHitRatio) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, operationAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("operation", operationAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedOperationHitRatio) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedOperationHitRatio) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedOperationHitRatio(cfg MetricConfig) metricMemcachedOperationHitRatio {
	m := metricMemcachedOperationHitRatio{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedOperations struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.operations metric with initial data.
func (m *metricMemcachedOperations) init() {
	m.data.SetName("memcached.operations")
	m.data.SetDescription("Operation counts.")
	m.data.SetUnit("{operations}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMemcachedOperations) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, typeAttributeValue string, operationAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("type", typeAttributeValue)
	dp.Attributes().PutStr("operation", operationAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedOperations) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedOperations) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedOperations(cfg MetricConfig) metricMemcachedOperations {
	m := metricMemcachedOperations{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedThreads struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.threads metric with initial data.
func (m *metricMemcachedThreads) init() {
	m.data.SetName("memcached.threads")
	m.data.SetDescription("Number of threads used by the memcached instance.")
	m.data.SetUnit("{threads}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricMemcachedThreads) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedThreads) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedThreads) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedThreads(cfg MetricConfig) metricMemcachedThreads {
	m := metricMemcachedThreads{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	config                            MetricsBuilderConfig // config of the metrics builder.
	startTime                         pcommon.Timestamp    // start time that will be applied to all recorded data points.
	metricsCapacity                   int                  // maximum observed number of metrics per resource.
	metricsBuffer                     pmetric.Metrics      // accumulates metrics data before emitting.
	buildInfo                         component.BuildInfo  // contains version information.
	metricMemcachedBytes              metricMemcachedBytes
	metricMemcachedCommands           metricMemcachedCommands
	metricMemcachedConnectionsCurrent metricMemcachedConnectionsCurrent
	metricMemcachedConnectionsTotal   metricMemcachedConnectionsTotal
	metricMemcachedCPUUsage           metricMemcachedCPUUsage
	metricMemcachedCurrentItems       metricMemcachedCurrentItems
	metricMemcachedEvictions          metricMemcachedEvictions
	metricMemcachedNetwork            metricMemcachedNetwork
	metricMemcachedOperationHitRatio  metricMemcachedOperationHitRatio
	metricMemcachedOperations         metricMemcachedOperations
	metricMemcachedThreads            metricMemcachedThreads
}

// MetricBuilderOption applies changes to default metrics builder.
type MetricBuilderOption interface {
	apply(*MetricsBuilder)
}

type metricBuilderOptionFunc func(mb *MetricsBuilder)

func (mbof metricBuilderOptionFunc) apply(mb *MetricsBuilder) {
	mbof(mb)
}

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) MetricBuilderOption {
	return metricBuilderOptionFunc(func(mb *MetricsBuilder) {
		mb.startTime = startTime
	})
}
func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.Settings, options ...MetricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		config:                            mbc,
		startTime:                         pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                     pmetric.NewMetrics(),
		buildInfo:                         settings.BuildInfo,
		metricMemcachedBytes:              newMetricMemcachedBytes(mbc.Metrics.MemcachedBytes),
		metricMemcachedCommands:           newMetricMemcachedCommands(mbc.Metrics.MemcachedCommands),
		metricMemcachedConnectionsCurrent: newMetricMemcachedConnectionsCurrent(mbc.Metrics.MemcachedConnectionsCurrent),
		metricMemcachedConnectionsTotal:   newMetricMemcachedConnectionsTotal(mbc.Metrics.MemcachedConnectionsTotal),
		metricMemcachedCPUUsage:           newMetricMemcachedCPUUsage(mbc.Metrics.MemcachedCPUUsage),
		metricMemcachedCurrentItems:       newMetricMemcachedCurrentItems(mbc.Metrics.MemcachedCurrentItems),
		metricMemcachedEvictions:          newMetricMemcachedEvictions(mbc.Metrics.MemcachedEvictions),
		metricMemcachedNetwork:            newMetricMemcachedNetwork(mbc.Metrics.MemcachedNetwork),
		metricMemcachedOperationHitRatio:  newMetricMemcachedOperationHitRatio(mbc.Metrics.MemcachedOperationHitRatio),
		metricMemcachedOperations:         newMetricMemcachedOperations(mbc.Metrics.MemcachedOperations),
		metricMemcachedThreads:            newMetricMemcachedThreads(mbc.Metrics.MemcachedThreads),
	}

	for _, op := range options {
		op.apply(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption interface {
	apply(pmetric.ResourceMetrics)
}

type resourceMetricsOptionFunc func(pmetric.ResourceMetrics)

func (rmof resourceMetricsOptionFunc) apply(rm pmetric.ResourceMetrics) {
	rmof(rm)
}

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return resourceMetricsOptionFunc(func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
	})
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return resourceMetricsOptionFunc(func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	})
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(options ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName(ScopeName)
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricMemcachedBytes.emit(ils.Metrics())
	mb.metricMemcachedCommands.emit(ils.Metrics())
	mb.metricMemcachedConnectionsCurrent.emit(ils.Metrics())
	mb.metricMemcachedConnectionsTotal.emit(ils.Metrics())
	mb.metricMemcachedCPUUsage.emit(ils.Metrics())
	mb.metricMemcachedCurrentItems.emit(ils.Metrics())
	mb.metricMemcachedEvictions.emit(ils.Metrics())
	mb.metricMemcachedNetwork.emit(ils.Metrics())
	mb.metricMemcachedOperationHitRatio.emit(ils.Metrics())
	mb.metricMemcachedOperations.emit(ils.Metrics())
	mb.metricMemcachedThreads.emit(ils.Metrics())

	for _, op := range options {
		op.apply(rm)
	}

	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(options ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(options...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordMemcachedBytesDataPoint adds a data point to memcached.bytes metric.
func (mb *MetricsBuilder) RecordMemcachedBytesDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricMemcachedBytes.recordDataPoint(mb.startTime, ts, val)
}

// RecordMemcachedCommandsDataPoint adds a data point to memcached.commands metric.
func (mb *MetricsBuilder) RecordMemcachedCommandsDataPoint(ts pcommon.Timestamp, val int64, commandAttributeValue AttributeCommand) {
	mb.metricMemcachedCommands.recordDataPoint(mb.startTime, ts, val, commandAttributeValue.String())
}

// RecordMemcachedConnectionsCurrentDataPoint adds a data point to memcached.connections.current metric.
func (mb *MetricsBuilder) RecordMemcachedConnectionsCurrentDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricMemcachedConnectionsCurrent.recordDataPoint(mb.startTime, ts, val)
}

// RecordMemcachedConnectionsTotalDataPoint adds a data point to memcached.connections.total metric.
func (mb *MetricsBuilder) RecordMemcachedConnectionsTotalDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricMemcachedConnectionsTotal.recordDataPoint(mb.startTime, ts, val)
}

// RecordMemcachedCPUUsageDataPoint adds a data point to memcached.cpu.usage metric.
func (mb *MetricsBuilder) RecordMemcachedCPUUsageDataPoint(ts pcommon.Timestamp, val float64, stateAttributeValue AttributeState) {
	mb.metricMemcachedCPUUsage.recordDataPoint(mb.startTime, ts, val, stateAttributeValue.String())
}

// RecordMemcachedCurrentItemsDataPoint adds a data point to memcached.current_items metric.
func (mb *MetricsBuilder) RecordMemcachedCurrentItemsDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricMemcachedCurrentItems.recordDataPoint(mb.startTime, ts, val)
}

// RecordMemcachedEvictionsDataPoint adds a data point to memcached.evictions metric.
func (mb *MetricsBuilder) RecordMemcachedEvictionsDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricMemcachedEvictions.recordDataPoint(mb.startTime, ts, val)
}

// RecordMemcachedNetworkDataPoint adds a data point to memcached.network metric.
func (mb *MetricsBuilder) RecordMemcachedNetworkDataPoint(ts pcommon.Timestamp, val int64, directionAttributeValue AttributeDirection) {
	mb.metricMemcachedNetwork.recordDataPoint(mb.startTime, ts, val, directionAttributeValue.String())
}

// RecordMemcachedOperationHitRatioDataPoint adds a data point to memcached.operation_hit_ratio metric.
func (mb *MetricsBuilder) RecordMemcachedOperationHitRatioDataPoint(ts pcommon.Timestamp, val float64, operationAttributeValue AttributeOperation) {
	mb.metricMemcachedOperationHitRatio.recordDataPoint(mb.startTime, ts, val, operationAttributeValue.String())
}

// RecordMemcachedOperationsDataPoint adds a data point to memcached.operations metric.
func (mb *MetricsBuilder) RecordMemcachedOperationsDataPoint(ts pcommon.Timestamp, val int64, typeAttributeValue AttributeType, operationAttributeValue AttributeOperation) {
	mb.metricMemcachedOperations.recordDataPoint(mb.startTime, ts, val, typeAttributeValue.String(), operationAttributeValue.String())
}

// RecordMemcachedThreadsDataPoint adds a data point to memcached.threads metric.
func (mb *MetricsBuilder) RecordMemcachedThreadsDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricMemcachedThreads.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...MetricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op.apply(mb)
	}
}
