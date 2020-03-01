package config

// EtcdConfig etcd配置
type EtcdConfig interface {
	GetEnabled() bool
	GetPort() int
	GetHost() string
}

// defaultEtcdConfig 默认etcd配置
type defaultEtcdConfig struct {
	Enabled bool	`json:"enabled"`
	Host	int		`json:"host"`
	Port	string	`json:"port"`
}

// GetEnabled etcd使能
func (config defaultEtcdConfig) GetEnabled() bool {
	return config.Enabled
}

// GetHost etcd主机地址
func (config defaultEtcdConfig) GetHost() string {
	return config.Host
}

// GetPort etcd端口
func (config defaultEtcdConfig) GetPort() int {
	return config.Port
}