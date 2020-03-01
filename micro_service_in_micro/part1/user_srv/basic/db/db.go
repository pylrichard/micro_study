package db

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/pylrichard/micro_study/micro_service_in_micro/part1/user_srv/basic/config"
	"github.com/micro/go-micro/util/log"
)

var (
	inited	bool
	mysqlDB *sql.DB
	mutex	sync.RWMutex
)

// Init 初始化数据库
func Init() {
	mutex.Lock()
	defer mutex.Unlock()

	var err error
	if inited {
		err = fmt.Errorf("[Init] db already inited")
		log.Logf(err.Error())
		return
	}
	//判断配置是否使用mysql
	if config.GetMysqlConfig().GetEnabled() {
		initMysql()
	}
	inited = true
}

// GetDB 获取DB信息
func GetDB() *sql.DB {
	return mysqlDB
}