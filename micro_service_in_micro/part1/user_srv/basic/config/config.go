package config

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
)

var (
	err error
)

var (
	//见application.yml
	defaultRootPath			= "app"
	defaultConfigFilePrefix = "application-"
	etcdConfig				defaultEtcdConfig
	mysqlConfig				defaultMysqlConfig
	profiles				defaultProfiles
	mutex					sync.RWMutex
	inited					bool
	separator				= string(filepath.Separator)
)

// Init 初始化配置，见part1.md的初始化配置的过程
func Init() {
	mutex.Lock()
	defer mutex.Unlock()

	if inited {
		log.Logf("[Init] already configed")
		return
	}

	/*
		加载yml配置
	*/
	//加载基础配置
	appPath, _ := filepath.Abs(".")
	// appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("." + separator, separator)))
	confPath := filepath.Join(appPath, "conf")
	os.Chdir(appPath)
	//加载application.yml
	err = config.Load(file.NewSource(file.WithPath(confPath + separator + "application.yml")))
	if err != nil {
		panic(err)
	}
	//获取配置include: etcd, db
	err = config.Get(defaultRootPath, "profiles").Scan(&profiles)
	if err != nil {
		panic(err)
	}
	log.Logf("[Init] load conf, path: %s, %+v", confPath + separator + "application.yml", profiles)
	/*
		加载其他配置文件
	*/
	//见profiles.go
	if len(profiles.GetInclude()) > 0) {
		//得到etcd和db
		include := strings.Split(profiles.GetInclude(), ",")
		sources := make([]source.Source, len(include))
		for i := 0; i < len(include); i++ {
			filePath := confPath + string(filepath.Separator) 
						+ defaultConfigFilePrefix + strings.TrimSpace(include[i]) + ".yml"
			log.Logf("[Init] load conf, path: %s\n", filePath)
			sources[i] = file.NewSource(file.WithPath(filePath))
		}
		err = config.Load(sources...)
		if err != nil {
			panic(err)
		}
	}
	//见application-etcd.yml和etcd.go
	config.Get(defaultRootPath, "etcd").Scan(&etcdConfig)
	config.Get(defaultRootPath, "mysql").Scan(&mysqlConfig)

	inited = true
}

// GetMysqlConfig 获取mysql配置
func GetMysqlConfig() (ret MysqlConfig) {
	return mysqlConfig
}

// GetEtcdConfig 获取etcd配置
func GetEtcdConfig() (ret EtcdConfig) {
	return etcdConfig
}