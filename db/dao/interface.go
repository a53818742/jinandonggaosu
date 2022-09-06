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
	InsertCar(counter *model.WeihuapinCarInsert) error
	UpdateCar(counter *model.WeihuapinCarUpdate) error
	OverCar(counter *model.WeihuapinCarOver) error
	GetCar(id string) (data map[string]interface{}, errorMsg string, errorCode int)
	GetRecord(status int, offset int, limit int) (data []map[string]interface{}, errorMsg string, errorCode int)

	InsertAdmin(counter *model.AdminInsert) error
	UpdateAdmin(counter *model.AdminUpdate) error
	OverAdmin(counter *model.AdminOver) error
	GetAdminList() (data []map[string]interface{}, errorMsg string, errorCode int)

	CheckAdminLevel(weichartid string) (level int)
	UserLogin(username string, pwd string) bool

	GetMsg(TimeLen int64) (data []map[string]interface{}, errorMsg string, errorCode int)
	GetRecordNum(status int, offset int, limit int) (data int, errorMsg string, errorCode int)
}

// CounterInterfaceImp 计数器数据模型实现
type CounterInterfaceImp struct{}

// Imp 实现实例
var Imp CounterInterface = &CounterInterfaceImp{}
