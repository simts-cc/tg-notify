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

// DB 資料連線資訊
type DB struct {
	Name     string
	Host     string
	Port     string
	Username string
	Password string
	Charset  string
	MaxIdle  int
	MaxOpen  int
}

// Orm 資料庫連線
type Orm struct {
	writer *orm.Ormer
	reader *orm.Ormer
}

// NewOrmer Orm 連線 實體
func NewOrmer(db DB) *orm.Ormer {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.SetMaxIdleConns(db.Name, db.MaxIdle)
	orm.SetMaxOpenConns(db.Name, db.MaxOpen)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", db.Username, db.Password, db.Host, db.Port, db.Name, db.Charset)
	orm.RegisterDataBase(db.Name, "mysql", dsn)
	o := orm.NewOrm()

	return &o
}

// NewOrm Orm 實體
func NewOrm() *Orm {
	maxIdle, e := strconv.Atoi(os.Getenv("DB_MAX_IDLE"))
	if e != nil {
		glog.Fatal(e)
	}
	maxOpen, e := strconv.Atoi(os.Getenv("DB_MAX_OPEN"))
	if e != nil {
		glog.Fatal(e)
	}
	w := DB{
		Name:     os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_WRITER"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Charset:  os.Getenv("DB_CHARSET"),
		MaxIdle:  maxIdle,
		MaxOpen:  maxOpen,
	}
	r := DB{
		Name:     os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_WRITER"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Charset:  os.Getenv("DB_CHARSET"),
		MaxIdle:  maxIdle,
		MaxOpen:  maxOpen,
	}

	return &Orm{
		writer: NewOrmer(w),
		reader: NewOrmer(r),
	}
}
