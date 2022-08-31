package dao

import (
	"wxcloudrun-golang/db/model"
)

// CounterInterface 计数器数据模型接口
type CounterInterface interface {
	GetCounter(id int32) (*model.CounterModel, error)
	UpsertCounter(counter *model.CounterModel) error
	ClearCounter(id int32) error
	InsertCar(counter *model.WeihuapinCarInsert) error
	UpdateCar(counter *model.WeihuapinCarUpdate) error
	OverCar(counter *model.WeihuapinCarOver) error
	GetCar(id string) (data map[string]interface{}, errorMsg string, errorCode int)
	GetRecord(status int) (data []map[string]interface{}, errorMsg string, errorCode int)
}

// CounterInterfaceImp 计数器数据模型实现
type CounterInterfaceImp struct{}

// Imp 实现实例
var Imp CounterInterface = &CounterInterfaceImp{}
