package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pylrichard/micro_study/micro_service_in_micro/part1/user_srv/basic/config"
	"github.com/micro/go-micro/util/log"
)

func initMysql() {
	var err error
	mysqlDB, err = sql.Open("mysql", config.GetMysqlConfig().GetURL())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	//设置最大开启连接数
	mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())
	//设置最大空闲连接数
	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())
	//测试连接是否成功
	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
		panic(err)
	}
}