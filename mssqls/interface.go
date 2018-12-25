package mssqls

/*
有关sqlServer数据库连接的封装
*/
import (
	"database/sql"
	"fmt"
	"time"

	"github.com/xndm-recommend/go-utils/errors_"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/xndm-recommend/go-utils/config"
)

type MssqlMethod interface {
	GetMssqlConnFromConf(c *config.ConfigEngine, name string)
}

type MssqlDbInfo struct {
	SqlDataDb *sql.DB
	TableName map[string]string
	MaxConns  int
	DbTimeOut int
}

func getMssqlLoginStr(section *config.MssqlDbData) string {
	//连接字符串
	return fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s",
		section.Host, section.Port, section.Db_name, section.User, section.Password)
}

func (this *MssqlDbInfo) createMssqlConns(login *config.MssqlDbData) {
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

func (this *MssqlDbInfo) GetMssqlConnFromConf(c *config.ConfigEngine, name string) {
	this.createMssqlConns(c.GetMssqlDataFromConf(name))
}
