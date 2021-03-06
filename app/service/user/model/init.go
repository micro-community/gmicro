package model

import (
	"fmt"
	"os"

	"github.com/liangjfblue/gmicro/app/service/user/configs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func Init(mysqlConf *configs.MysqlConfig) {
	var (
		err  error
		addr string
	)

	addr = os.Getenv("CONFIGOR_MYSQL_ADDR")
	if addr == "" {
		addr = mysqlConf.Addr
	}
	str := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", mysqlConf.User, mysqlConf.Password, addr, mysqlConf.Db)
	DB, err = gorm.Open("mysql", str)
	if err != nil {
		panic(err)
	}

	DB.LogMode(true)
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(mysqlConf.MaxIdleConns)
	DB.DB().SetMaxOpenConns(mysqlConf.MaxOpenConns)

	DB.AutoMigrate(&TBUser{}, &TBMoney{})

	return
}

func CheckPageSize(offset, limit int32) (int32, int32) {
	if offset < 0 {
		offset = 0
	}
	if limit > 20 {
		limit = 20
	}
	return offset, limit
}
