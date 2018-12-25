package errors_

import (
	"github.com/getsentry/raven-go"
	"github.com/xndm-recommend/go-utils/conf_read"
)

func SentryRavenInit(this *conf_read.ConfigEngine, sectionName string) error {
	sentryDSN := this.GetString(sectionName)
	CheckEmptyValue(sentryDSN)
	err := raven.SetDSN(sentryDSN)
	CheckFatalErr(err)
	return err
}

func SentryCaptureError(err error) {
	raven.CaptureErrorAndWait(err, nil)
}
