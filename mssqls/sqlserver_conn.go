package mssqls

/*
有关sqlServer数据库连接的封装
*/
import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/zhanglanhui/go-utils/utils/conf_utils"
)

type MssqlDbInfo struct {
	SqlDataDb *sql.DB
	TableName map[string]string
	MaxConns  int
	DbTimeOut int
}

func GetMssqlConnFromConf(this *conf_utils.ConfigEngine, SectionName string) *MssqlDbInfo {
	DbInfo := new(MssqlDbInfo)
	sLogin := getMssqlDataFromConf(this, SectionName)
	DbInfo.createMssqlConns(sLogin)
	return DbInfo
}
