package scheduler

import "errors"

func (r *TelemetryRetryRetry) ShouldRetry(err error) bool {
	if err == nil || r.Permanent {
		return false
	}
	return errors.Is(err, ErrTelemetryRetryTransient)
}
func (r *TelemetryRetryRetry) State() string {
	if r.Permanent {
		return "permanent"
	}
	if r.Attempts > 0 {
		return "retrying"
	}
	return "pending"
}
