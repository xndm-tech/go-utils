package errors_

import (
	"github.com/getsentry/raven-go"
)

func SentryCaptureError(err error) {
	raven.CaptureErrorAndWait(err, nil)
}
