package mysqls

/*
有关mysql数据库连接的封装
*/
import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zhanglanhui/go-utils/utils/conf_utils"
)

type MysqlDbInfo struct {
	SqlDataDb *sql.DB
	TableName map[string]string
	MaxConns  int
	DbTimeOut int
}

func GetDbConnFromConf(this *conf_utils.ConfigEngine, sectionName string) *MysqlDbInfo {
	DbInfo := new(MysqlDbInfo)
	sLogin := getSqlDataFromConf(this, sectionName)
	DbInfo.createDatabaseConns(sLogin)
	return DbInfo
}
