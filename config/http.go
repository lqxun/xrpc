package config

type HttpConfig struct {
	Port int
}

// NewHttpConfig 此处可以通过viper优化
func NewHttpConfig() *HttpConfig {
	return &HttpConfig{Port: 7000}
}
