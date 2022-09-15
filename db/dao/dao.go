package dao

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const tableName2 = "weihuapincar"

//
//// ClearCounter 清除Counter
//func (imp *CounterInterfaceImp) ClearCounter(id int32) error {
//	cli := db.Get()
//	return cli.Table(tableName).Delete(&model.CounterModel{Id: id}).Error
//}
//
//// UpsertCounter 更新/写入counter
//func (imp *CounterInterfaceImp) UpsertCounter(counter *model.CounterModel) error {
//	cli := db.Get()
//	return cli.Table(tableName).Save(counter).Error
//}
//
//// GetCounter 查询Counter
//func (imp *CounterInterfaceImp) GetCounter(id int32) (*model.CounterModel, error) {
//	var err error
//	var counter = new(model.CounterModel)
//
//	cli := db.Get()
//
//	err = cli.Table(tableName).Where("id = ?", id).First(counter).Error
//
//	return counter, err
//}

// InsertCar 更新/写入counter
func (imp *CounterInterfaceImp) InsertCar(counter *model.WeihuapinCarInsert) (int, string) {
	cli := db.Get()

	//rows, err := cli.Table(tableName2).Where(" wechartid= ? and status=0", counter.WechartId).Rows()
	//if err != nil {
	//	return -1, "出错"
	//}
	//defer rows.Close()
	//columns, _ := rows.Columns()            //获取列的信息
	//count := len(columns)                   //列的数量
	//var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	//for i := 0; i < count; i++ {
	//	var ii interface{} //为空接口分配内存
	//	values[i] = &ii    //取得这些内存的指针，因后继的Scan函数只接受指针
	//}
	//for rows.Next() {
	//	err0 := rows.Scan(values...) //开始读行，Scan函数只接受指针变量
	//	if err0 != nil {
	//		return -2, "出错"
	//	}
	//	return -3, "记录已存在，请勿重复提交！"
	//}
	er := cli.Table(tableName2).Save(counter).Error
	if er == nil {
		return 0, ""
	}
	return -4, ""
}

// UpdateCar 更新/写入counter
func (imp *CounterInterfaceImp) UpdateCar(counter *model.WeihuapinCarUpdate) error {
	cli := db.Get()
	return cli.Table(tableName2).Save(counter).Error
}

// OverCar  车辆离场
func (imp *CounterInterfaceImp) OverCar(counter *model.WeihuapinCarOver) error {
	cli := db.Get()
	return cli.Table(tableName2).Save(counter).Error
}

// OverMsg 更新消息发送
func (imp *CounterInterfaceImp) OverMsg(counter *model.OverMsg) error {
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
	for i := 0; i < count; i++ {
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
			var rawValue = *(values[i].(*interface{})) //读出raw数据，类型为byte
			if reflect.TypeOf(rawValue) == reflect.TypeOf([]byte{0}) {
				b, _ := rawValue.([]byte)
				v := string(b)   //将raw数据转换成字符串
				ret[colName] = v //colName是键，v是值
			} else {
				ret[colName] = rawValue
			}
		}
		break
	}

	return ret, "", 0
}

// GetRecordNum 查询某一天的记录
func (imp *CounterInterfaceImp) GetRecordNum(status int) (data int, errorMsg string, errorCode int) {

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
		strr = " status=1 and UNIX_TIMESTAMP(intime)>" + strconv.FormatInt(now, 10)
		break
	default:
		strr = " status=0 or (status=1 and UNIX_TIMESTAMP(intime)>" + strconv.FormatInt(now, 10) + ") "

	}

	rows, err := cli.Table(tableName2).Select("count(*)").Where(strr).Rows()
	if err != nil {
		fmt.Println("Query ", err.Error())

		return 0, err.Error(), -100
	}
	defer rows.Close()
	columns, _ := rows.Columns()            //获取列的信息
	count := len(columns)                   //列的数量
	var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	for i := 0; i < count; i++ {
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
			var rawValue = *(values[i].(*interface{})) //读出raw数据，类型为byte
			if reflect.TypeOf(rawValue) == reflect.TypeOf([]byte{0}) {
				b, _ := rawValue.([]byte)

				v := string(b)   //将raw数据转换成字符串
				ret[colName] = v //colName是键，v是值

			} else {
				ret[colName] = rawValue
			}
		}
		break
	}

	data, _ = strconv.Atoi(ret["count(*)"].(string))
	return data, "", 0
}

// GetWeihuapinByCN 根据CN编号查询危化品
func (imp *CounterInterfaceImp) GetWeihuapinByCN(weihuapin string) (data []map[string]interface{}, errorMsg string, errorCode int) {

	cli := db.Get()
	rows, err := cli.Table("chemicalandyingjichuzhifangan").Select("ChemicalName,ChemicalInfo,ChemicalType,LiHuaTeXing,WendingAndFangYingxing,YunShuXinXi,OpraAndSave,Revealdisposition").Where("YunShuXinXi like '%危险货物编号：" + weihuapin + ",%'").Rows()
	if err != nil {
		fmt.Println("Query ", err.Error())

		return nil, err.Error(), -100
	}
	defer rows.Close()
	columns, _ := rows.Columns()            //获取列的信息
	count := len(columns)                   //列的数量
	var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	for i := 0; i < count; i++ {
		var ii interface{} //为空接口分配内存
		values[i] = &ii    //取得这些内存的指针，因后继的Scan函数只接受指针
	}

	ret := make([]map[string]interface{}, 0)
	for rows.Next() {
		err0 := rows.Scan(values...) //开始读行，Scan函数只接受指针变量
		if err0 != nil {
			panic(err)
		}
		m := map[string]interface{}{}
		for i, colName := range columns {
			var rawValue = *(values[i].(*interface{})) //读出raw数据，类型为byte
			if reflect.TypeOf(rawValue) == reflect.TypeOf([]byte{0}) {
				b, _ := rawValue.([]byte)
				v := string(b) //将raw数据转换成字符串
				m[colName] = v //colName是键，v是值
			} else {
				m[colName] = rawValue
			}
		}
		ret = append(ret, m)
	}
	//fmt.Println("GetWeihuapinByCN", len(ret), ret)
	return ret, "", 0
}

// GetWeihuapinByUN 根据UN编号查询危化品
func (imp *CounterInterfaceImp) GetWeihuapinByUN(weihuapin string) (data []map[string]interface{}, errorMsg string, errorCode int) {

	cli := db.Get()
	rows, err := cli.Table("chemicalandyingjichuzhifangan").Select("ChemicalName,Element,ChemicalInfo,ChemicalType,LiHuaTeXing,WendingAndFangYingxing,YunShuXinXi,OpraAndSave,Revealdisposition").Where("YunShuXinXi like '%UN编号：" + weihuapin + ",%'").Rows()
	if err != nil {
		fmt.Println("Query ", err.Error())

		return nil, err.Error(), -100
	}
	defer rows.Close()
	columns, _ := rows.Columns()            //获取列的信息
	count := len(columns)                   //列的数量
	var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	for i := 0; i < count; i++ {
		var ii interface{} //为空接口分配内存
		values[i] = &ii    //取得这些内存的指针，因后继的Scan函数只接受指针
	}

	ret := make([]map[string]interface{}, 0)
	for rows.Next() {
		err0 := rows.Scan(values...) //开始读行，Scan函数只接受指针变量
		if err0 != nil {
			panic(err)
		}
		m := map[string]interface{}{}
		for i, colName := range columns {
			var rawValue = *(values[i].(*interface{})) //读出raw数据，类型为byte
			if reflect.TypeOf(rawValue) == reflect.TypeOf([]byte{0}) {
				b, _ := rawValue.([]byte)
				v := string(b) //将raw数据转换成字符串
				m[colName] = v //colName是键，v是值
			} else {
				m[colName] = rawValue
			}
		}
		ret = append(ret, m)
	}

	return ret, "", 0
}

// GetWeihuapin 根据危化品名称查询危化品
func (imp *CounterInterfaceImp) GetWeihuapin(weihuapin string) (data []map[string]interface{}, errorMsg string, errorCode int) {

	cli := db.Get()
	rows, err := cli.Table("chemicalandyingjichuzhifangan").Select("ChemicalName,Element,ChemicalInfo,ChemicalType,LiHuaTeXing,WendingAndFangYingxing,YunShuXinXi,OpraAndSave,Revealdisposition").Where("ChemicalName like '%" + weihuapin + "%'").Rows()
	if err != nil {
		fmt.Println("Query ", err.Error())

		return nil, err.Error(), -100
	}
	defer rows.Close()
	columns, _ := rows.Columns()            //获取列的信息
	count := len(columns)                   //列的数量
	var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	for i := 0; i < count; i++ {
		var ii interface{} //为空接口分配内存
		values[i] = &ii    //取得这些内存的指针，因后继的Scan函数只接受指针
	}

	ret := make([]map[string]interface{}, 0)
	for rows.Next() {
		err0 := rows.Scan(values...) //开始读行，Scan函数只接受指针变量
		if err0 != nil {
			panic(err)
		}
		m := map[string]interface{}{}
		for i, colName := range columns {
			var rawValue = *(values[i].(*interface{})) //读出raw数据，类型为byte
			if reflect.TypeOf(rawValue) == reflect.TypeOf([]byte{0}) {
				b, _ := rawValue.([]byte)
				v := string(b) //将raw数据转换成字符串
				m[colName] = v //colName是键，v是值
			} else {
				m[colName] = rawValue
			}
		}
		ret = append(ret, m)
	}

	return ret, "", 0
}

// GetRecord 查询某一天的记录
func (imp *CounterInterfaceImp) GetRecord(status int, offset int, limit int) (data []map[string]interface{}, errorMsg string, errorCode int) {

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
		strr = " status=1 and UNIX_TIMESTAMP(intime)>" + strconv.FormatInt(now, 10)
		break
	default:
		strr = " status=0 or (status=1 and UNIX_TIMESTAMP(intime)>" + strconv.FormatInt(now, 10) + ") "

	}

	rows, err := cli.Table(tableName2).Where(strr).Order(" ID desc").Offset(offset).Limit(limit).Rows()
	if err != nil {
		fmt.Println("Query ", err.Error())

		return nil, err.Error(), -100
	}
	defer rows.Close()
	columns, _ := rows.Columns()            //获取列的信息
	count := len(columns)                   //列的数量
	var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	for i := 0; i < count; i++ {
		var ii interface{} //为空接口分配内存
		values[i] = &ii    //取得这些内存的指针，因后继的Scan函数只接受指针
	}
	ret := make([]map[string]interface{}, 0)
	for rows.Next() {
		err0 := rows.Scan(values...)  //开始读行，Scan函数只接受指针变量
		m := map[string]interface{}{} //用于存放1列的 [键/值] 对
		if err0 != nil {
			panic(err)
		}
		for i, colName := range columns {
			var rawValue = *(values[i].(*interface{})) //读出raw数据，类型为byte
			if reflect.TypeOf(rawValue) == reflect.TypeOf([]byte{0}) {
				b, _ := rawValue.([]byte)

				v := string(b) //将raw数据转换成字符串

				m[colName] = v //colName是键，v是值
			} else {
				m[colName] = rawValue
			}
		}
		ret = append(ret, m) //将单行所有列的键值对附加在总的返回值上（以行为单位）
	}

	return ret, "", 0
}

// GetMsg 获取当前应该推送通知的记录
func (imp *CounterInterfaceImp) GetMsg(TimeLen int64) (data []map[string]interface{}, errorMsg string, errorCode int) {
	cli := db.Get()

	rows, err := cli.Table(tableName2).Where("status=0 and msgnum=0 and  UNIX_TIMESTAMP(intime)<" + strconv.FormatInt(time.Now().Unix()-TimeLen, 10)).Order(" ID ").Rows()
	if err != nil {
		fmt.Println("Query ", err.Error())
		return nil, err.Error(), -100
	}
	defer rows.Close()
	columns, _ := rows.Columns()            //获取列的信息
	count := len(columns)                   //列的数量
	var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	for i := 0; i < count; i++ {
		var ii interface{} //为空接口分配内存
		values[i] = &ii    //取得这些内存的指针，因后继的Scan函数只接受指针
	}
	ret := make([]map[string]interface{}, 0)
	for rows.Next() {
		err0 := rows.Scan(values...)  //开始读行，Scan函数只接受指针变量
		m := map[string]interface{}{} //用于存放1列的 [键/值] 对
		if err0 != nil {
			panic(err)
		}
		for i, colName := range columns {
			var rawValue = *(values[i].(*interface{})) //读出raw数据，类型为byte
			if reflect.TypeOf(rawValue) == reflect.TypeOf([]byte{0}) {
				b, _ := rawValue.([]byte)

				v := string(b) //将raw数据转换成字符串

				m[colName] = v //colName是键，v是值
			} else {
				m[colName] = rawValue
			}
		}
		ret = append(ret, m) //将单行所有列的键值对附加在总的返回值上（以行为单位）
	}
	return ret, "", 0
}

// UserLogin 用户登录
func (imp *CounterInterfaceImp) UserLogin(username string, pwd string) bool {
	cli := db.Get()
	rows, err := cli.Table("users").Where("level>=1 and username= ? and pwd=?", username, pwd).Rows()
	if err != nil {
		return false
	}
	defer rows.Close()
	columns, _ := rows.Columns()            //获取列的信息
	count := len(columns)                   //列的数量
	var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	for i := 0; i < count; i++ {
		var ii interface{} //为空接口分配内存
		values[i] = &ii    //取得这些内存的指针，因后继的Scan函数只接受指针
	}
	for rows.Next() {
		err0 := rows.Scan(values...) //开始读行，Scan函数只接受指针变量
		if err0 != nil {
			return false
		}

		return true
	}
	return false
}

// UserLogin2 用户登录
func (imp *CounterInterfaceImp) UserLogin2(username string, pwd string, wechartid string) bool {
	cli := db.Get()
	rows, err := cli.Table("users").Where("level>=1 and username= ? and pwd=?", username, pwd).Rows()
	if err != nil {
		return false
	}
	defer rows.Close()
	columns, _ := rows.Columns()            //获取列的信息
	count := len(columns)                   //列的数量
	var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	for i := 0; i < count; i++ {
		var ii interface{} //为空接口分配内存
		values[i] = &ii    //取得这些内存的指针，因后继的Scan函数只接受指针
	}
	for rows.Next() {
		err0 := rows.Scan(values...) //开始读行，Scan函数只接受指针变量
		if err0 != nil {
			return false
		}
		ret := map[string]interface{}{} //创建返回值：不定长的map类型切片
		for i, colName := range columns {
			var rawValue = *(values[i].(*interface{})) //读出raw数据，类型为byte
			if reflect.TypeOf(rawValue) == reflect.TypeOf([]byte{0}) {
				b, _ := rawValue.([]byte)

				v := string(b) //将raw数据转换成字符串

				ret[colName] = v //colName是键，v是值
			} else {
				ret[colName] = rawValue
			}
		}
		if len(ret["wechartid"].(string)) < 10 {
			cli.Table("users").Save(&model.UserLoginWechart{
				Wechartid: wechartid,
				ID:        int(ret["ID"].(int64)),
			})
		}

		return true
	}
	return false
}

// UserLogin3 用户登录
func (imp *CounterInterfaceImp) UserLogin3(username string, pwd string, wechartid string) bool {
	cli := db.Get()
	rows, err := cli.Table("users").Where("level>=2 and username= ? and pwd=?", username, pwd).Rows()
	if err != nil {
		return false
	}
	defer rows.Close()
	columns, _ := rows.Columns()            //获取列的信息
	count := len(columns)                   //列的数量
	var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	for i := 0; i < count; i++ {
		var ii interface{} //为空接口分配内存
		values[i] = &ii    //取得这些内存的指针，因后继的Scan函数只接受指针
	}
	for rows.Next() {
		err0 := rows.Scan(values...) //开始读行，Scan函数只接受指针变量
		if err0 != nil {
			return false
		}

		return true
	}
	return false
}

// UserAdd 更新/写入counter
func (imp *CounterInterfaceImp) UserAdd(counter *model.UserInsert) (int, string) {
	cli := db.Get()
	rows, err := cli.Table("users").Where("username= ?", counter.UserName).Rows()
	if err != nil {
		return -1, "出错"
	}
	defer rows.Close()
	columns, _ := rows.Columns()            //获取列的信息
	count := len(columns)                   //列的数量
	var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	for i := 0; i < count; i++ {
		var ii interface{} //为空接口分配内存
		values[i] = &ii    //取得这些内存的指针，因后继的Scan函数只接受指针
	}
	for rows.Next() {
		err0 := rows.Scan(values...) //开始读行，Scan函数只接受指针变量
		if err0 != nil {
			return -2, "出错"
		}

		return -3, "用户名已存在"
	}
	counter.Wechartid = "none"
	if counter.Level <= 0 {
		counter.Level = 1
	}
	e := cli.Table("users").Save(counter).Error
	if e == nil {
		return 0, ""
	}
	return -4, e.Error()
}

// UserUpdate 更新/写入counter
func (imp *CounterInterfaceImp) UserUpdate(counter *model.UserUpdate) error {
	cli := db.Get()
	return cli.Table("users").Save(counter).Error
}

// UserDelete 更新/写入counter
func (imp *CounterInterfaceImp) UserDelete(counter *model.UserDelete) error {
	cli := db.Get()
	return cli.Table("users").Delete(counter).Error
}

// UserList 查询某一天的记录
func (imp *CounterInterfaceImp) UserList() (data []map[string]interface{}, errorMsg string, errorCode int) {
	cli := db.Get()
	rows, err := cli.Table("users").Order(" ID desc").Rows()
	if err != nil {
		fmt.Println("Query ", err.Error())
		return nil, err.Error(), -100
	}
	defer rows.Close()
	columns, _ := rows.Columns()            //获取列的信息
	count := len(columns)                   //列的数量
	var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	for i := 0; i < count; i++ {
		var ii interface{} //为空接口分配内存
		values[i] = &ii    //取得这些内存的指针，因后继的Scan函数只接受指针
	}
	ret := make([]map[string]interface{}, 0)
	for rows.Next() {
		err0 := rows.Scan(values...)  //开始读行，Scan函数只接受指针变量
		m := map[string]interface{}{} //用于存放1列的 [键/值] 对
		if err0 != nil {
			panic(err)
		}
		for i, colName := range columns {
			var rawValue = *(values[i].(*interface{})) //读出raw数据，类型为byte
			if reflect.TypeOf(rawValue) == reflect.TypeOf([]byte{0}) {
				b, _ := rawValue.([]byte)

				v := string(b) //将raw数据转换成字符串

				m[colName] = v //colName是键，v是值
			} else {
				m[colName] = rawValue
			}
		}
		ret = append(ret, m) //将单行所有列的键值对附加在总的返回值上（以行为单位）
	}

	return ret, "", 0
}

// RecordAdd 更新/写入counter
func (imp *CounterInterfaceImp) RecordAdd(counter *model.RecordInsert) error {
	cli := db.Get()
	return cli.Table("record").Save(counter).Error
}

// RecordUpdate 更新/写入counter
func (imp *CounterInterfaceImp) RecordUpdate(counter *model.RecordUpdate) error {
	cli := db.Get()
	return cli.Table("record").Save(counter).Error
}

// RecordDelete 更新/写入counter
func (imp *CounterInterfaceImp) RecordDelete(counter *model.RecordDelete) error {
	cli := db.Get()
	return cli.Table("record").Delete(counter).Error
}

// RecordList 查询某一天的记录
func (imp *CounterInterfaceImp) RecordList(userid int) (data []map[string]interface{}, errorMsg string, errorCode int) {

	cli := db.Get()
	rows, err := cli.Table("record").Where("userid=?", userid).Order(" ID desc").Rows()
	if err != nil {
		fmt.Println("Query ", err.Error())
		return nil, err.Error(), -100
	}
	defer rows.Close()
	columns, _ := rows.Columns()            //获取列的信息
	count := len(columns)                   //列的数量
	var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	for i := 0; i < count; i++ {
		var ii interface{} //为空接口分配内存
		values[i] = &ii    //取得这些内存的指针，因后继的Scan函数只接受指针
	}
	ret := make([]map[string]interface{}, 0)
	for rows.Next() {
		err0 := rows.Scan(values...)  //开始读行，Scan函数只接受指针变量
		m := map[string]interface{}{} //用于存放1列的 [键/值] 对
		if err0 != nil {
			panic(err)
		}
		for i, colName := range columns {
			var rawValue = *(values[i].(*interface{})) //读出raw数据，类型为byte
			if reflect.TypeOf(rawValue) == reflect.TypeOf([]byte{0}) {
				b, _ := rawValue.([]byte)
				v := string(b) //将raw数据转换成字符串
				m[colName] = v //colName是键，v是值
			} else {
				m[colName] = rawValue
			}
		}
		ret = append(ret, m) //将单行所有列的键值对附加在总的返回值上（以行为单位）
	}

	return ret, "", 0
}

func (imp *CounterInterfaceImp) GetAdminList() (darray []string) {
	MySQLDB := db.GetDB()
	darray = []string{}
	dmm := map[string]string{}
	rows, err002 := MySQLDB.Query("select t1.userid,users.wechartid from (select * from record where starttime<='2022-09-01 05:00:00' and endtime >='2022-09-01 05:00:00')t1 LEFT JOIN users on t1.userid=users.ID")
	if err002 != nil {
		fmt.Println("GetAdminList Error ", err002.Error())
		return
	} else {
		defer rows.Close()
		columns, _ := rows.Columns()            //获取列的信息
		count := len(columns)                   //列的数量
		var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
		for i := 0; i < count; i++ {
			var ii interface{} //为空接口分配内存
			values[i] = &ii    //取得这些内存的指针，因后继的Scan函数只接受指针
		}
		ret := make([]map[string]interface{}, 0)
		for rows.Next() {
			err0 := rows.Scan(values...)  //开始读行，Scan函数只接受指针变量
			m := map[string]interface{}{} //用于存放1列的 [键/值] 对
			if err0 != nil {
				return
			}
			for i, colName := range columns {
				var rawValue = *(values[i].(*interface{})) //读出raw数据，类型为byte
				if reflect.TypeOf(rawValue) == reflect.TypeOf([]byte{0}) {
					b, _ := rawValue.([]byte)
					v := string(b) //将raw数据转换成字符串
					m[colName] = v //colName是键，v是值
					if colName == "wechartid" {
						dmm[v] = v
					}
				} else {
					m[colName] = rawValue

				}

			}
			ret = append(ret, m) //将单行所有列的键值对附加在总的返回值上（以行为单位）
		}
	}
	for key, _ := range dmm {
		darray = append(darray, key)
	}
	return
}
