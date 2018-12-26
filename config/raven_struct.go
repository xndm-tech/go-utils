package config

import (
	"github.com/getsentry/raven-go"
	"github.com/xndm-recommend/go-utils/errors_"
)

func (c *ConfigEngine) SentryRavenInit(name string) error {
	sentryDSN := c.GetString(name)
	errors_.CheckEmptyValue(sentryDSN)
	err := raven.SetDSN(sentryDSN)
	errors_.CheckFatalErr(err)
	return err
}
