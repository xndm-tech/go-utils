package errors_

import (
	"github.com/getsentry/raven-go"
	"github.com/xndm-recommend/go-utils/config"
)

func SentryRavenInit(c *config.ConfigEngine, name string) error {
	sentryDSN := c.GetString(name)
	CheckEmptyValue(sentryDSN)
	err := raven.SetDSN(sentryDSN)
	CheckFatalErr(err)
	return err
}

func SentryCaptureError(err error) {
	raven.CaptureErrorAndWait(err, nil)
}
