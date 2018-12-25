package mssqls

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/xndm-recommend/go-utils/conf_read"
	"github.com/xndm-recommend/go-utils/errors_"
)

type mssqlDbYamlData struct {
	User       string            `yaml:"user"`
	Password   string            `yaml:"password"`
	Host       string            `yaml:"host"`
	Port       string            `yaml:"port"`
	Db_name    string            `yaml:"db_name"`
	Table_name map[string]string `yaml:"table_name"`
	Max_conns  int               `yaml:"max_conns"`
	Time_out   int               `yaml:"time_out"`
}

func getMssqlDataFromConf(this *conf_read.ConfigEngine, sectionName string) *mssqlDbYamlData {
	sqlServerLogin := new(mssqlDbYamlData)
	login := this.GetStruct(sectionName, sqlServerLogin)
	return login.(*mssqlDbYamlData)
}

func getMssqlLoginStr(section *mssqlDbYamlData) string {
	//连接字符串
	return fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s",
		section.Host, section.Port, section.Db_name, section.User, section.Password)
}

func (this *MssqlDbInfo) createMssqlConns(login *mssqlDbYamlData) {
	db, err := sql.Open("mssql", getMssqlLoginStr(login))
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
