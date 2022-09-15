package db

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbInstance *gorm.DB

// Init 初始化数据库
func Init() error {

	source := "%s:%s@tcp(%s)/%s?readTimeout=1500ms&writeTimeout=1500ms&charset=utf8&loc=Local&&parseTime=true"
	user := os.Getenv("MYSQL_USERNAME")
	pwd := os.Getenv("MYSQL_PASSWORD")
	addr := os.Getenv("MYSQL_ADDRESS")
	dataBase := os.Getenv("MYSQL_DATABASE")
	if dataBase == "" {
		dataBase = "golang_demo"
	}
	source = fmt.Sprintf(source, user, pwd, addr, dataBase)
	fmt.Println("start init mysql with ", source)

	db, err := gorm.Open(mysql.Open(source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		}})
	if err != nil {
		fmt.Println("DB Open error,err=", err.Error())
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("DB Init error,err=", err.Error())
		return err
	}

	// 用于设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(100)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(200)
	// 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	dbInstance = db

	fmt.Println("finish init mysql with ", source)

	rows, err002 := dbInstance.Table("(select * from record where starttime<='2022-09-01 05:00:00' and endtime >='2022-09-01 05:00:00')t1 LEFT JOIN users on t1.userid=users.ID").Select("t1.userid,users.wechartid").Rows()
	if err002 != nil {
		fmt.Println("Query ", err.Error())

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
				fmt.Println("0=", colName, reflect.TypeOf(v), reflect.TypeOf(v).Name())
			} else {
				m[colName] = rawValue
				fmt.Println("1=", colName, reflect.TypeOf(rawValue), reflect.TypeOf(rawValue).Name())
			}

		}
		ret = append(ret, m) //将单行所有列的键值对附加在总的返回值上（以行为单位）
	}
	return nil
}

// Get ...
func Get() *gorm.DB {
	return dbInstance
}
