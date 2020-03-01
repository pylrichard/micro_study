package config

// Profiles 属性配置文件
type Profiles interface {
	GetInclude() string
}

// defaultProfiles 默认配置
type defaultProfiles struct {
	Include string `json:"include"`
}

// GetInclude 获取配置文件名
// 配置文件名格式为"application-xxx.yml"
// 多个文件名以逗号隔开，省略前缀"application-"
func (profile defaultProfiles) GetInclude() string {
	return profile.Include
}