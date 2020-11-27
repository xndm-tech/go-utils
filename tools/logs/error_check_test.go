package logs_test

import (
	"fmt"
	"testing"

	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/tools/logs"

	log "github.com/sirupsen/logrus"
)

const (
	logFile    = "../../config/test.yaml"
	logxmlFile = "../../config/seelog.xml"
)

func TestErrLog1(b *testing.T) {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}

func TestErrLog(b *testing.T) {
	appConfig := &config.ConfigEngine{}
	logs.LoggerSetup(logxmlFile)
	logs.CheckFatalErr(appConfig.Load(logFile))
	logs.CheckFatalErr(appConfig.SentryRavenInit("Sentry_dsn"))
	logs.CheckErrSendEmail(fmt.Errorf("asdf"))
}

func TestRusErrLog1(b *testing.T) {
	appConfig := &config.ConfigEngine{}
	logs.CheckFatalErr(appConfig.Load(logFile))
	logs.CheckFatalErr(appConfig.SentryRavenInit("Sentry_dsn"))
	logs.CheckLogrusCaptureError(fmt.Errorf("asdf"), nil, "yyy", "ppp")
}
