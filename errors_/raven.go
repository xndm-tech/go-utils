package errors_

import (
	"runtime"

	"github.com/cihub/seelog"
	"github.com/getsentry/raven-go"
	"github.com/pkg/errors"
)

func SentryCaptureError(err error) {
	raven.CaptureErrorAndWait(err, nil)
}

func CheckErrSendEmail(err error) {
	if err != nil {
		errDetail := errors.WithStack(err)
		_, file, line, _ := runtime.Caller(1)
		seelog.Error("Important error:", file, ":", line, errDetail)
		raven.CaptureError(errDetail, nil)
	}
}
