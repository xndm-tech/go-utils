package mysqls

import (
	"bytes"
	"database/sql"
	"time"

	"github.com/zhanglanhui/go-utils/utils/conf_utils"
	"github.com/zhanglanhui/go-utils/utils/err_utils"
)

type mysqlDbYamlData struct {
	User       string            `yaml:"user"`
	Password   string            `yaml:"password"`
	Host       string            `yaml:"host"`
	Port       string            `yaml:"port"`
	Db_name    string            `yaml:"db_name"`
	Table_name map[string]string `yaml:"table_name"`
	Max_conns  int               `yaml:"max_conns"`
	Time_out   int               `yaml:"time_out"`
}

func getSqlDataFromConf(this *conf_utils.ConfigEngine, sectionName string) *mysqlDbYamlData {
	mysqlLogin := new(mysqlDbYamlData)
	login := this.GetStruct(sectionName, mysqlLogin)
	return login.(*mysqlDbYamlData)
}

func getSqlLoginStr(section *mysqlDbYamlData) string {
	// 读取配置文件
	var buffer bytes.Buffer
	buffer.WriteString(section.User)
	buffer.WriteString(":")
	buffer.WriteString(section.Password)
	buffer.WriteString("@tcp(")
	buffer.WriteString(section.Host)
	buffer.WriteString(":")
	buffer.WriteString(section.Port)
	buffer.WriteString(")/")
	buffer.WriteString(section.Db_name)
	buffer.WriteString("?charset=utf8")

	return buffer.String()
}

func (this *MysqlDbInfo) createDatabaseConns(login *mysqlDbYamlData) {
	db, err := sql.Open("mysql", getSqlLoginStr(login))
	err_utils.CheckFatalErr(err)
	db.SetConnMaxLifetime(time.Duration(login.Time_out) * time.Second)
	db.SetMaxOpenConns(login.Max_conns)
	db.SetMaxIdleConns(login.Max_conns)
	err = db.Ping()
	err_utils.CheckFatalErr(err)
	this.SqlDataDb = db
	this.TableName = login.Table_name
	this.MaxConns = login.Max_conns
	this.DbTimeOut = login.Time_out
}
