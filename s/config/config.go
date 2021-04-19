package config

type config struct{}

type Config interface {
	SetConfigName(cName string)
	SetConfigType(cType string)
	AddConfigPath(cPath string)
}
