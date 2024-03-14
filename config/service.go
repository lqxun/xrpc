package config

type Service struct {
	Port int
}

// NewService 此处可以通过viper优化
func NewService() *Service {
	return &Service{Port: 7777}
}
