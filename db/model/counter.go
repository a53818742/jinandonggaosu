package model

import "time"

// CounterModel 计数器模型
type CounterModel struct {
	Id        int32     `gorm:"column:id" json:"id"`
	Count     int32     `gorm:"column:count" json:"count"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

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
}

type WeihuapinGetOne struct {
	WechatID string
}

type WeihuapinGetList struct {
	Status int
}
