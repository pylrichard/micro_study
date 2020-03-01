package basic

import (
	"github.com/pylrichard/micro_study/micro_service_in_micro/part1/user_srv/basic/config"
	"github.com/pylrichard/micro_study/micro_service_in_micro/part1/user_srv/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}