package db

import (
	"Go-RestfulAPI/config"
	"Go-RestfulAPI/pkg/logging"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
)

/**
* @Author: super
* @Date: 2020-08-26 18:16
* @Description: 操作数据库
**/

var _db *gorm.DB

func InitDB() {
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	var err error
	mysqlURL := config.GetMysqlUrl()

	_db, err = gorm.Open("mysql", mysqlURL)
	if err != nil {
		logging.GetLogger().Error("conn mysql error", zap.String("err", err.Error()))
		panic("连接数据库失败, error=" + err.Error())
	}

	//设置数据库连接池参数
	_db.DB().SetMaxOpenConns(100) //设置数据库连接池最大连接数
	_db.DB().SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

//通过此方法获得数据库连接
func GetDB() *gorm.DB {
	return _db
}
