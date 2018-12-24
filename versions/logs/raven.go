package logs

import (
	"github.com/getsentry/raven-go"
	"github.com/zhanglanhui/go-utils/utils/conf_utils"
)

func (this *conf_utils.ConfigEngine) SentryRavenInit(SectionName string) {
	sentryDSN := this.GetString(SectionName)
	CheckEmptyValue(sentryDSN)
	err := raven.SetDSN(sentryDSN)
	CheckFatalErr(err)
}
