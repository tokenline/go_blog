package config

type AppConfig struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Default string
}
