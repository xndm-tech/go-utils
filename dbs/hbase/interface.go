package hbase

/*
有关hbase数据库连接的封装
*/
import (
	"fmt"
	"time"

	data2 "github.com/xndm-recommend/go-utils/tools/types"

	"github.com/xndm-recommend/go-utils/common/consts"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/tools/errs"
)

type HBaseMethod interface {
	GetDbConnFromConf(c *config.ConfigEngine, name string)
	SelectStrList(sql string, para ...interface{}) (dest []string, err error)
	SelectIntList(sql string, para ...interface{}) (dest []int, err error)
	QueryStruct(sql string, dest ...interface{}) (err error)
	QueryIdMap(sql string, para ...interface{}) (dest map[string]string, err error)
	GetDb() *sqlx.DB
	GetTableName(key string) string
}

type HBaseDbInfo struct {
	sqlDataDb *sqlx.DB
	tableName map[string]string
	maxConns  int
	dbTimeOut int
}

func getHBaseLoginStr(data *config.MysqlDbData) string {
	return data2.JoinStrs(data.User, ":",
		data.Password, "@tcp(", data.Host, ":",
		data.Port, ")/", data.Db_name, "?charset=utf8")
}

func (this *HBaseDbInfo) createDatabaseConns(login *config.MysqlDbData) {
	db, err := sqlx.Open("mysql", getHBaseLoginStr(login))
	errs.CheckFatalErr(err)
	db.SetConnMaxLifetime(time.Duration(login.Time_out) * time.Second)
	db.SetMaxOpenConns(login.Max_conns)
	db.SetMaxIdleConns(login.Max_conns)
	errs.CheckFatalErr(db.Ping())
	this.sqlDataDb = db
	this.tableName = login.Table_name
	this.maxConns = login.Max_conns
	this.dbTimeOut = login.Time_out
}

func (this *HBaseDbInfo) GetDbConnFromConf(c *config.ConfigEngine, name string) {
	this.createDatabaseConns(c.GetMySqlFromConf(name))
}

func (this *HBaseDbInfo) GetDb() *sqlx.DB {
	return this.sqlDataDb
}

func (this *HBaseDbInfo) GetTableName(key string) string {
	if val, ok := this.tableName[key]; ok {
		return val
	} else {
		errs.CheckCommonErr(fmt.Errorf(fmt.Sprintf("key %s not in tablenames.", key)))
		return consts.BLANK
	}
}
