package logs

/*
有关log打印的封装
*/
import (
	"github.com/cihub/seelog"
	"github.com/zhanglanhui/go-utils/utils/err_utils"
)

func LoggerSetup(configPath string) {
	logger, err := seelog.LoggerFromConfigAsFile(configPath)
	if err != nil {
		err_utils.CheckFatalErr(err)
		return
	}
	seelog.ReplaceLogger(logger)
}
