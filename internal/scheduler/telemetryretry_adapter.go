package scheduler

import "errors"

func (r *TelemetryRetryRetry) ShouldRetry(err error) bool {
	return err != nil && !r.Permanent && err.Error() == ErrTelemetryRetryTransient.Error()
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
