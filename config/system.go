package config

type System struct {
	Env      string `mapstructure:"env"`
	Port     int    `mapstructure:"port"`
	LogPath  string `mapstructure:"log-path"`
	DBType   string `mapstructure:"db-type"`
	UseRedis bool   `mapstructure:"use-redis"`
}
