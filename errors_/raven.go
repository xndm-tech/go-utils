package errors_

import (
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
		seelog.Errorf("%+v", errDetail)
		raven.CaptureError(errDetail, nil)
	}
}
