package errs

import (
	"github.com/cihub/seelog"
	"github.com/getsentry/raven-go"
	"github.com/pkg/errors"
	"github.com/xndm-tech/go-utils/tools/logs"
)

func SentryCaptureError(err error) {
	raven.CaptureErrorAndWait(err, nil)
}

func CheckErrSendEmail(err error) {
	if err != nil {
		errDetail := errors.WithStack(err)
		_ = seelog.Errorf("%+v", errDetail)
		raven.CaptureError(errDetail, nil)
	}
}

func CheckLogrusCaptureError(err error, tags map[string]string, args ...interface{}) {
	if err != nil {
		errDetail := errors.WithStack(err)
		var errs = make(map[string]interface{})
		for k, v := range tags {
			errs[k] = v
		}
		errs["error"] = err
		logs.Error(errs, args)
		raven.CaptureError(errDetail, nil)
	}
}
