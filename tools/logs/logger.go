package logs

/*
有关log打印的封装
*/
import (
	"github.com/cihub/seelog"
	"github.com/xndm-recommend/go-utils/tools/errs"
)

func LoggerSetup(c string) {
	logger, err := seelog.LoggerFromConfigAsFile(c)
	if err != nil {
		errs.CheckFatalErr(err)
		return
	}
	seelog.ReplaceLogger(logger)
}
