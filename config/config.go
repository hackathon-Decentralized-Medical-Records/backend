package config

type Setting struct {
	JWT      JWT      `mapstructure:"jwt"`
	Logrus   Logrus   `mapstructure:"logrus"`
	Database Database `mapstructure:"database"`
	Cors     CORS     `mapstructure:"cors"`
	System   System   `mapstructure:"system"`
	Ethereum Ethereum `mapstructure:"ethereum"`
	Redis    Redis    `mapstructure:"redis"`
}

const (
	ConfigEnv         = "GVA_CONFIG"
	ConfigDefaultFile = "config.yaml"
	ConfigTestFile    = "config.test.yaml"
	ConfigDebugFile   = "config.debug.yaml"
	ConfigReleaseFile = "config.release.yaml"
)
