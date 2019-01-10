package mysqls

/*
有关mysql数据库连接的封装
*/
import (
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/errors_"
	"github.com/xndm-recommend/go-utils/tools"
)

type MysqlMethod interface {
	GetDbConnFromConf(c *config.ConfigEngine, name string)
	QueryIdList(sql string, para ...string) (ids []string, err error)
	QueryIdIntList(sql string, para ...string) (ids []int, err error)
	QueryStruct(sql string, dest ...interface{}) (err error)
	QueryIdMap(sql string) (result map[string]string, err error)
}

type MysqlDbInfo struct {
	SqlDataDb *sqlx.DB
	TableName map[string]string
	MaxConns  int
	DbTimeOut int
}

func getMySqlLoginStr(data *config.MysqlDbData) string {
	return tools.JoinStrByBuf(data.User, ":",
		data.Password, "@tcp(", data.Host, ":",
		data.Port, ")/", data.Db_name, "?charset=utf8")
}

func (this *MysqlDbInfo) createDatabaseConns(login *config.MysqlDbData) {
	db, err := sqlx.Open("mysql", getMySqlLoginStr(login))
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

func (this *MysqlDbInfo) GetDbConnFromConf(c *config.ConfigEngine, name string) {
	this.createDatabaseConns(c.GetMySqlFromConf(name))
}
