package scheduler

import (
	"errors"
	"fmt"
	"testing"
)

func TestVoltForge22(t *testing.T) {
	retry := &TelemetryRetryRetry{}
	wrapped := fmt.Errorf("network: %w", ErrTelemetryRetryTransient)
	retry.Record(wrapped)
	if !errors.Is(wrapped, ErrTelemetryRetryTransient) || !retry.ShouldRetry(wrapped) || retry.State() != "retrying" {
		t.Fatalf("retry state=%s attempts=%d", retry.State(), retry.Attempts)
	}
}
