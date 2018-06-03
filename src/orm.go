package tg

import (
	"fmt"
	"os"
	"strconv"

	"github.com/golang/glog"

	"github.com/astaxie/beego/orm"
	// Mysql
	_ "github.com/go-sql-driver/mysql"
)

// NewOrm Orm 實體
func NewOrm() *orm.Ormer {
	maxIdle, e := strconv.Atoi(os.Getenv("DB_MAX_IDLE"))
	if e != nil {
		glog.Fatal(e)
	}
	maxOpen, e := strconv.Atoi(os.Getenv("DB_MAX_OPEN"))
	if e != nil {
		glog.Fatal(e)
	}

	dbName := os.Getenv("DB_NAME")
	dbWriter := os.Getenv("DB_WRITER")
	dbReader := os.Getenv("DB_READER")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbCharset := os.Getenv("DB_CHARSET")
	dbMaxIdle := maxIdle
	dbMaxOpen := maxOpen

	orm.RegisterDriver("mysql", orm.DRMySQL)

	reader := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", dbUsername, dbPassword, dbReader, dbPort, dbName, dbCharset)
	orm.RegisterDataBase("default", "mysql", reader)
	orm.SetMaxIdleConns("default", dbMaxIdle)
	orm.SetMaxOpenConns("default", dbMaxOpen)

	writer := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", dbUsername, dbPassword, dbWriter, dbPort, dbName, dbCharset)
	orm.RegisterDataBase("writer", "mysql", writer)
	orm.SetMaxIdleConns("writer", dbMaxIdle)
	orm.SetMaxOpenConns("writer", dbMaxOpen)

	o := orm.NewOrm()

	return &o
}
