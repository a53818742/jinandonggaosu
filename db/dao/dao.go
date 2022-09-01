package dao

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const tableName = "Counters"
const tableName2 = "weihuapincar"

// ClearCounter 清除Counter
func (imp *CounterInterfaceImp) ClearCounter(id int32) error {
	cli := db.Get()
	return cli.Table(tableName).Delete(&model.CounterModel{Id: id}).Error
}

// UpsertCounter 更新/写入counter
func (imp *CounterInterfaceImp) UpsertCounter(counter *model.CounterModel) error {
	cli := db.Get()
	return cli.Table(tableName).Save(counter).Error
}

// GetCounter 查询Counter
func (imp *CounterInterfaceImp) GetCounter(id int32) (*model.CounterModel, error) {
	var err error
	var counter = new(model.CounterModel)

	cli := db.Get()

	err = cli.Table(tableName).Where("id = ?", id).First(counter).Error

	return counter, err
}

// InsertCar 更新/写入counter
func (imp *CounterInterfaceImp) InsertCar(counter *model.WeihuapinCarInsert) error {
	cli := db.Get()
	return cli.Table(tableName2).Save(counter).Error
}

// UpdateCar 更新/写入counter
func (imp *CounterInterfaceImp) UpdateCar(counter *model.WeihuapinCarUpdate) error {
	cli := db.Get()
	return cli.Table(tableName2).Save(counter).Error
}

// OverCar 更新/写入counter
func (imp *CounterInterfaceImp) OverCar(counter *model.WeihuapinCarOver) error {
	cli := db.Get()
	return cli.Table(tableName2).Save(counter).Error
}

// GetCar 查询Counter
func (imp *CounterInterfaceImp) GetCar(id string) (data map[string]interface{}, errorMsg string, errorCode int) {

	cli := db.Get()
	rows, err := cli.Table(tableName2).Where("wechartid = ? and status=0", id).Order("ID desc").Rows()
	if err != nil {
		fmt.Println("Query ", err.Error())

		return nil, err.Error(), -100
	}
	defer rows.Close()
	columns, _ := rows.Columns() //获取列的信息
	count := len(columns)        //列的数量

	var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	for i, _ := range values {
		var ii interface{} //为空接口分配内存
		values[i] = &ii    //取得这些内存的指针，因后继的Scan函数只接受指针
	}
	ret := map[string]interface{}{} //创建返回值：不定长的map类型切片
	for rows.Next() {
		err0 := rows.Scan(values...) //开始读行，Scan函数只接受指针变量
		if err0 != nil {
			panic(err)
		}
		for i, colName := range columns {
			var raw_value = *(values[i].(*interface{})) //读出raw数据，类型为byte
			if reflect.TypeOf(raw_value) == reflect.TypeOf([]byte{0}) {
				b, _ := raw_value.([]byte)
				v := string(b)   //将raw数据转换成字符串
				ret[colName] = v //colName是键，v是值
			} else {
				ret[colName] = raw_value
			}
		}
		break
	}

	return ret, "", 0
}

// GetRecord 查询某一天的记录
func (imp *CounterInterfaceImp) GetRecord(status int) (data []map[string]interface{}, errorMsg string, errorCode int) {

	cli := db.Get()
	t := time.Now()
	addTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	now := addTime.Unix()
	strr := ""
	switch status {
	case 0:
		strr = " status=0"
		break
	case 1:
		strr = " status=1 and intime>" + strconv.FormatInt(now, 10)
		break
	default:
		strr = " status=0 or (status=1 and intime>" + strconv.FormatInt(now, 10) + ") "

	}

	rows, err := cli.Table(tableName2).Where(strr).Order(" ID desc").Rows()
	if err != nil {
		fmt.Println("Query ", err.Error())

		return nil, err.Error(), -100
	}
	defer rows.Close()
	columns, _ := rows.Columns()            //获取列的信息
	count := len(columns)                   //列的数量
	var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	for i, _ := range values {
		var ii interface{} //为空接口分配内存
		values[i] = &ii    //取得这些内存的指针，因后继的Scan函数只接受指针
	}
	ret := []map[string]interface{}{} //创建返回值：不定长的map类型切片
	for rows.Next() {
		err0 := rows.Scan(values...)  //开始读行，Scan函数只接受指针变量
		m := map[string]interface{}{} //用于存放1列的 [键/值] 对
		if err0 != nil {
			panic(err)
		}
		for i, colName := range columns {
			var raw_value = *(values[i].(*interface{})) //读出raw数据，类型为byte
			if reflect.TypeOf(raw_value) == reflect.TypeOf([]byte{0}) {
				b, _ := raw_value.([]byte)

				v := string(b) //将raw数据转换成字符串

				m[colName] = v //colName是键，v是值
			} else {
				m[colName] = raw_value
			}
		}
		ret = append(ret, m) //将单行所有列的键值对附加在总的返回值上（以行为单位）
	}

	return ret, "", 0
}
