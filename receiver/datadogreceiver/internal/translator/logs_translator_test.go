package translator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
)

func TestTranslateLogs(t *testing.T) {
	input := []DatadogLog{
		{
			Message:   "Test log message",
			Status:    "info",
			Service:   "test-service",
			Host:      "test-host",
			Source:    "go",
			Timestamp: time.Now().UnixNano(),
			Tags:      []string{"env:test"},
			Attributes: map[string]interface{}{
				"user_id": "test-user",
			},
		},
	}

	lt := NewLogsTranslator(component.BuildInfo{
		Command: "test-cmd",
		Version: "v0.0.1",
	})

	otlpLogs, err := lt.TranslateLogs(input)
	require.NoError(t, err)

	// Basic validation
	require.NotNil(t, otlpLogs)
	require.Greater(t, otlpLogs.ResourceLogs().Len(), 0)

	// Example: check that log record body matches the original message
	rl := otlpLogs.ResourceLogs().At(0)
	sl := rl.ScopeLogs().At(0)
	lr := sl.LogRecords().At(0)
	require.Equal(t, "Test log message", lr.Body().Str())
}
