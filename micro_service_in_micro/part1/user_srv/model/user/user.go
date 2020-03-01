package user

import (
	"fmt"
	"sync"

	proto "github.com/pylrichard/micro_study/micro_service_in_micro/part1/user_srv/proto/user"
)

var (
	srv		*Service
	mutex	sync.RWMutex
)

// Service 服务
type Service struct {
}

// UserService 用户服务类
type UserService interface {
	// QueryUserByName 根据用户名获取用户信息
	QueryUserByName(userName string) (ret *proto.User, err error)
}

// GetService 获取服务类
func GetService() (Service, error) {
	if srv == nil {
		return nil, fmt.Errorf("[GetService] Service is not inited")
	}
	return srv, nil
}

// Init 初始化用户服务层
func Init() {
	mutex.Lock()
	defer mutex.Unlock()

	if srv != nil {
		return
	}
	srv = &Service{}
}