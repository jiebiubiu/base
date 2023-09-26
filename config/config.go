package config

const FileServiceEndpoint = "http://127.0.0.1:6660"

type IConfig interface {
	LoadConfig() error
	GetConfig() *Config
}
