package dao

import (
	"time"
	"wxcloudrun-golang/db/model"
)

type WeihuapinCar struct {
	Id          int32     `gorm:"column:id" json:"id"`
	CarNo       string    `gorm:"column:CarNo" json:"CarNo"`
	Dunwei      float32   `gorm:"column:dunwei" json:"dunwei"`
	Weihuapin   string    `gorm:"column:weihuapin" json:"weihuapin"`
	Name        string    `gorm:"column:name" json:"name"`
	Moble       string    `gorm:"column:moble" json:"moble"`
	Zhuangkuang string    `gorm:"column:zhuangkuang" json:"zhuangkuang"`
	Beizhu      string    `gorm:"column:beizhu" json:"beizhu"`
	Intime      time.Time `gorm:"column:intime" json:"intime"`
	Outtime     time.Time `gorm:"column:outtime" json:"outtime"`
	Recordtime  time.Time `gorm:"column:recordtime" json:"recordtime"`
	Photo       string    `gorm:"column:photo" json:"photo"`
	Outtype     int8      `gorm:"column:outtype" json:"outtype"`
	Status      int8      `gorm:"column:status" json:"status"`
	Wechartid   string    `gorm:"column:wechartid" json:"wechartid"`
	Msgnum      string    `gorm:"column:msgnum" json:"msgnum"`
	Msgtime     string    `gorm:"column:msgtime" json:"msgtime"`
}

// CounterInterface 计数器数据模型接口
type CounterInterface interface {
	InsertCar(counter *model.WeihuapinCarInsert) (int, string)
	UpdateCar(counter *model.WeihuapinCarUpdate) error
	OverCar(counter *model.WeihuapinCarOver) error
	GetCar(id string) (data map[string]interface{}, errorMsg string, errorCode int)
	GetRecord(status int, offset int, limit int) (data []map[string]interface{}, errorMsg string, errorCode int)

	UserLogin(username string, pwd string) bool

	GetMsg(TimeLen int64) (data []map[string]interface{}, errorMsg string, errorCode int)
	GetRecordNum(status int) (data int, errorMsg string, errorCode int)
	GetWeihuapin(weihuapin string) (data []map[string]interface{}, errorMsg string, errorCode int)
	GetWeihuapinByUN(weihuapin string) (data []map[string]interface{}, errorMsg string, errorCode int)
	GetWeihuapinByCN(weihuapin string) (data []map[string]interface{}, errorMsg string, errorCode int)

	OverMsg(counter *model.OverMsg) error

	UserList() (data []map[string]interface{}, errorMsg string, errorCode int)
	UserAdd(counter *model.UserInsert) (int, string)
	UserUpdate(counter *model.UserUpdate) error
	UserDelete(counter *model.UserDelete) error

	RecordAdd(counter *model.RecordInsert) error
	RecordUpdate(counter *model.RecordUpdate) error
	RecordDelete(counter *model.RecordDelete) error
	RecordList(userid int) (data []map[string]interface{}, errorMsg string, errorCode int)

	UserLogin2(username string, pwd string, wechartid string) bool
	UserLogin3(username string, pwd string, wechartid string) bool

	GetAdminList() (darray []string)
}

// CounterInterfaceImp 计数器数据模型实现
type CounterInterfaceImp struct{}

// Imp 实现实例
var Imp CounterInterface = &CounterInterfaceImp{}
