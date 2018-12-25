package mssqls

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/errors_"
)

func getMssqlLoginStr(section *config.MssqlDbYamlData) string {
	//连接字符串
	return fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s",
		section.Host, section.Port, section.Db_name, section.User, section.Password)
}

func (this *MssqlDbInfo) createMssqlConns(login *config.MssqlDbYamlData) {
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
