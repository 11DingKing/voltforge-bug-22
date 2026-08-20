package scheduler

import "errors"

var ErrTelemetryRetryTransient = errors.New("telemetryretry temporarily unavailable")
var ErrTelemetryRetryPermanent = errors.New("telemetryretry permanently rejected")

type TelemetryRetryRetry struct {
	Attempts  int
	Permanent bool
}

func (r *TelemetryRetryRetry) Record(err error) {
	if err == nil {
		r.Permanent = false
		return
	}
	if err.Error() == ErrTelemetryRetryTransient.Error() {
		r.Attempts++
	} else {
		r.Permanent = true
	}
}
