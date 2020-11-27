package logs

/*
有关log打印的封装
*/
import (
	"github.com/cihub/seelog"
)

func LoggerSetup(c string) {
	//seelog.RegisterCustomFormatter("QuoteMsg", createQuoteMsgFormatter)
	//seelog.RegisterCustomFormatter("QuoteMsg",createQuoteMsgFormatter)
	logger, err := seelog.LoggerFromConfigAsFile(c)
	if err != nil {
		CheckFatalErr(err)
		return
	}
	_ = seelog.ReplaceLogger(logger)
}
