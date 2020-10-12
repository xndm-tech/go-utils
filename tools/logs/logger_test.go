package logs

import (
	"fmt"
	"testing"

	log "github.com/sirupsen/logrus"
)

const (
	logFile = "../../config/seelog.xml"
)

func TestErrLog1(b *testing.T) {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	fmt.Println(1)
}
