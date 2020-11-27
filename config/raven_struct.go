package config

import (
	"github.com/getsentry/raven-go"
	"github.com/xndm-recommend/go-utils/tools/logs"
)

func (c *ConfigEngine) SentryRavenInit(name string) error {
	sentryDSN := c.GetString(name)
	logs.CheckEmptyValue(sentryDSN)
	err := raven.SetDSN(sentryDSN)
	logs.CheckFatalErr(err)
	return err
}
