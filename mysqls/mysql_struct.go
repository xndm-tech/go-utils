package mysqls

import (
	"database/sql"
	"time"

	"github.com/xndm-recommend/go-utils/tools"

	"github.com/xndm-recommend/go-utils/errors_"
	//"github.com/xndm-recommend/go-utils/conf_read"
)

func GetMySqlLoginStr(data *MysqlDbYamlData) (string, map[string]string, int, int) {
	section := this.getSqlDataFromConf(sectionName)
	return tools.JoinStrByBuf(section.User, ":",
			section.Password, "@tcp(", section.Host, ":",
			section.Port, ")/", section.Db_name, "?charset=utf8"),
		section.Table_name,
		section.Max_conns,
		section.Time_out
}

func (this *MysqlDbInfo) createDatabaseConns(login *mysqlDbYamlData) {
	db, err := sql.Open("mysql", getSqlLoginStr(login))
	errors_.CheckFatalErr(err)
	db.SetConnMaxLifetime(time.Duration(login.Time_out) * time.Second)
	db.SetMaxOpenConns(login.Max_conns)
	db.SetMaxIdleConns(login.Max_conns)
	err = db.Ping()
	errors_.CheckFatalErr(err)
	this.SqlDataDb = db
	this.TableName = login.Table_name
	this.MaxConns = login.Max_conns
	this.DbTimeOut = login.Time_out
}
