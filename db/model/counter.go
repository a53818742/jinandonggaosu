package model

import "time"

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

type OverMsg struct {
	Id     int `gorm:"column:ID" json:"ID"`
	MsgNum int `gorm:"column:msgnum" json:"msgnum"`
	Status int `gorm:"column:status" json:"status"`
}

type WeihuapinGetOne struct {
	WechartID string `json:"wechartid"`
}

type WeihuapinGetList struct {
	Status int `json:"status"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type WeihuapinInfo struct {
	Weihuapin string `json:"weihuapin"`
}

type UserInsert struct {
	ID        int    `gorm:"column:ID" json:"ID"`
	Level     int    `gorm:"column:level" json:"level"`
	UserDesc  string `gorm:"column:userdesc" json:"userdesc"`
	UserName  string `gorm:"column:username" json:"username"`
	Pwd       string `gorm:"column:pwd" json:"pwd"`
	Mobile    string `gorm:"column:mobile" json:"mobile"`
	Wechartid string `gorm:"column:wechartid" json:"wechartid"`
}

type UserUpdate struct {
	ID        int    `gorm:"column:ID" json:"ID"`
	Level     int    `gorm:"column:level" json:"level"`
	UserDesc  string `gorm:"column:userdesc" json:"userdesc"`
	UserName  string `gorm:"column:username" json:"username"`
	Pwd       string `gorm:"column:pwd" json:"pwd"`
	Mobile    string `gorm:"column:mobile" json:"mobile"`
	Wechartid string `gorm:"column:wechartid" json:"wechartid"`
}
type UserLoginWechart struct {
	ID        int    `gorm:"column:ID" json:"ID"`
	Wechartid string `gorm:"column:wechartid" json:"wechartid"`
}

type UserDelete struct {
	ID int32 `gorm:"column:ID" json:"ID"`
}

type UserLoginStruct struct {
	Username  string `gorm:"column:username" json:"username"`
	Pwd       string `gorm:"column:pwd" json:"pwd"`
	Wechartid string `gorm:"column:wechartid" json:"wechartid"`
}

type RecordInsert struct {
	ID        int    `gorm:"column:ID" json:"ID"`
	UserID    int    `gorm:"column:userid" json:"userid"`
	StartTime string `gorm:"column:starttime" json:"starttime"`
	EndTime   string `gorm:"column:endtime" json:"endtime"`
}

type RecordUpdate struct {
	ID        int    `gorm:"column:ID" json:"ID"`
	StartTime string `gorm:"column:starttime" json:"starttime"`
	EndTime   string `gorm:"column:endtime" json:"endtime"`
}

type RecordDelete struct {
	ID int32 `gorm:"column:ID" json:"ID"`
}
