package mssqls

/*
有关sqlServer数据库连接的封装
*/
import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/xndm-recommend/go-utils/conf_read"
)

type MssqlDbInfo struct {
	SqlDataDb *sql.DB
	TableName map[string]string
	MaxConns  int
	DbTimeOut int
}

func GetMssqlConnFromConf(this *conf_read.ConfigEngine, sectionName string) *MssqlDbInfo {
	DbInfo := new(MssqlDbInfo)
	sLogin := getMssqlDataFromConf(this, sectionName)
	DbInfo.createMssqlConns(sLogin)
	return DbInfo
}
