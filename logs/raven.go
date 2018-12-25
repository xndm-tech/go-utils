package logs

import (
	"github.com/getsentry/raven-go"
	"github.com/xndm-recommend/go-utils/conf_read"
	"github.com/xndm-recommend/go-utils/errors_"
)

func SentryRavenInit(this *conf_read.ConfigEngine, SectionName string) {
	sentryDSN := this.GetString(SectionName)
	errors_.CheckEmptyValue(sentryDSN)
	err := raven.SetDSN(sentryDSN)
	errors_.CheckFatalErr(err)
}
