package config

// Service is a set of service-related configurations
type Service struct {
	Port string `yaml:"port" env:"PORT" env-description:"Service port"`
	Host string `yaml:"host" env:"HOST" env-description:"Service host"`
}

// MainConfig is a whole set of app settings
type MainConfig struct {
	Service Service `yaml:"service"`
}
