package model

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// APILogs 資料表架構
type APILogs struct {
	ID        uint32    `orm:"auto"`
	URI       string    `orm:"column(uri);varchar(60)"`
	ReqData   string    `orm:"column(req_data);varchar(1024)"`
	ResData   string    `orm:"column(res_data);varchar(1024)"`
	Headers   string    `orm:"column(headers);varchar(1024)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

// TableName 資料表名稱
func (apilogs *APILogs) TableName() string {
	return "api_logs"
}

// TableEngine 資料表引擎
func (apilogs *APILogs) TableEngine() string {
	return "INNODB"
}

func init() {
	orm.RegisterModel(new(APILogs))
}

// APILogsAdd 新增logs
func APILogsAdd(o orm.Ormer, apilogs APILogs) (int64, error) {
	o.Using("writer")
	return o.Insert(&apilogs)
}
