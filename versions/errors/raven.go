package errors

import (
	"github.com/getsentry/raven-go"
	"github.com/zhanglanhui/go-utils/utils/conf_utils"
)

func SentryRavenInit(this *conf_utils.ConfigEngine, sectionName string) error {
	sentryDSN := this.GetString(sectionName)
	CheckEmptyValue(sentryDSN)
	err := raven.SetDSN(sentryDSN)
	CheckFatalErr(err)
	return err
}

func SentryCaptureError(err error) {
	raven.CaptureErrorAndWait(err, nil)
}
