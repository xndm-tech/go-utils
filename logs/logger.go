package logs

/*
有关log打印的封装
*/
import (
	"github.com/cihub/seelog"
	"github.com/xndm-recommend/go-utils/errors_"
)

func LoggerSetup(configPath string) {
	logger, err := seelog.LoggerFromConfigAsFile(configPath)
	if err != nil {
		errors_.CheckFatalErr(err)
		return
	}
	seelog.ReplaceLogger(logger)
}
