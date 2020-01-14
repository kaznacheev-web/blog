package config

// Server is a set of service-related configurations
type Server struct {
	Port string `yaml:"port" env:"PORT" env-description:"Server port"`
	Host string `yaml:"host" env:"HOST" env-description:"Server host"`
}

type Database struct {
	URI string `yaml:"uri" env:"MONGO_URI" env-description:"MongoDB connection URI"`
}

// MainConfig is a whole set of app settings
type MainConfig struct {
	Service  Server   `yaml:"service"`
	Database Database `yaml:"database"`
	RootDir  string
}
