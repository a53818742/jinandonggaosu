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
type WeihuapinCarInsert struct {
	Id          int32     `gorm:"column:ID" json:"ID"`
	CarNo       string    `gorm:"column:carNo" json:"carNo"`
	Dunwei      float32   `gorm:"column:dunwei" json:"dunwei"`
	Weihuapin   string    `gorm:"column:weihuapin" json:"weihuapin"`
	Name        string    `gorm:"column:name" json:"name"`
	Moble       string    `gorm:"column:moble" json:"moble"`
	Zhuangkuang string    `gorm:"column:zhuangkuang" json:"zhuangkuang"`
	Beizhu      string    `gorm:"column:beizhu" json:"beizhu"`
	Intime      time.Time `gorm:"column:intime" json:"intime"`
	Recordtime  time.Time `gorm:"column:recordtime" json:"recordtime"`
	Photo       string    `gorm:"column:photo" json:"photo"`
	Outtype     int8      `gorm:"column:outtype" json:"outtype"`
	Status      int8      `gorm:"column:status" json:"status"`
	WechartId   string    `gorm:"column:wechartid" json:"wechartid"`
}
type WeihuapinCarUpdate struct {
	Id          int32   `gorm:"column:ID" json:"ID"`
	CarNo       string  `gorm:"column:carNo" json:"carNo"`
	Dunwei      float32 `gorm:"column:dunwei" json:"dunwei"`
	Weihuapin   string  `gorm:"column:weihuapin" json:"weihuapin"`
	Name        string  `gorm:"column:name" json:"name"`
	Moble       string  `gorm:"column:moble" json:"moble"`
	Zhuangkuang string  `gorm:"column:zhuangkuang" json:"zhuangkuang"`
	Beizhu      string  `gorm:"column:beizhu" json:"beizhu"`
	Photo       string  `gorm:"column:photo" json:"photo"`
}
type WeihuapinCarOver struct {
	Id      int32     `gorm:"column:ID" json:"ID"`
	Outtime time.Time `gorm:"column:outtime" json:"outtime"`
	Outtype int8      `gorm:"column:outtype" json:"outtype"`
	Status  int8      `gorm:"column:status" json:"status"`
}

type WeihuapinGetOne struct {
	WechartID string `json:"wechartid"`
}

type WeihuapinGetList struct {
	Status int `json:"status"`
}

type AdminInsert struct {
	Wechartid  string    `gorm:"column:wechartid" json:"wechartid"`
	Level      int       `gorm:"column:level" json:"level"`
	Name       int       `gorm:"column:name" json:"name"`
	Mobile     int       `gorm:"column:mobile" json:"mobile"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

type AdminUpdate struct {
	ID     int `gorm:"column:ID" json:"ID"`
	Name   int `gorm:"column:name" json:"name"`
	Mobile int `gorm:"column:mobile" json:"mobile"`
}

type AdminOver struct {
	ID    int `gorm:"column:ID" json:"ID"`
	Level int `gorm:"column:level" json:"level"`
}

type UserLoginStruct struct {
	Username string `gorm:"column:username" json:"username"`
	Pwd      string `gorm:"column:pwd" json:"pwd"`
}
