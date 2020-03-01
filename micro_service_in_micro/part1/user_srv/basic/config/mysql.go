package config

// MysqlConfig mysql配置
type MysqlConfig interface {
	GetURL() string
	GetEnabled() bool
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
}

// defaultMysqlConfig 默认mysql配置
type defaultMysqlConfig struct {
	URL					string	`json:"url"`
	Enabled				bool	`json:"enabled"`
	MaxIdleConnection	int		`json:"maxIdleConnection"`
	MaxOpenConnection	int		`json:"maxOpenConnection"`
}

// GetURL mysql连接地址
func (config defaultMysqlConfig) GetURL() string {
	return config.URL
}

// GetEnabled mysql使能
func (config defaultMysqlConfig) GetEnabled() bool {
	return config.Enabled
}

// GetMaxIdleConnection 最大空闲连接数
func (config defaultMysqlConfig) GetMaxIdleConnection() int {
	return config.MaxIdleConnection
}

// GetMaxOpenConnection 最大开启连接数
func (config defaultMysqlConfig) GetMaxOpenConnection() int {
	return config.MaxOpenConnection
}