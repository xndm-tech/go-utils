package errs

import (
	"fmt"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/cihub/seelog"
	"github.com/xndm-recommend/go-utils/tools/logs"
)

const (
	logFile = "../../config/seelog.xml"
)

func TestErrLog(b *testing.T) {
	logs.LoggerSetup(logFile)
	CheckCommonErr(fmt.Errorf("asdf"))
	CheckCommonInfo("%d_%d", 1, 2)
	//seelog.Info("asfd")
	defer seelog.Flush()
}

func TestErrLog1(b *testing.T) {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
