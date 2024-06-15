package config

type Logrus struct {
	OutputType   int    `mapstructure:"output-type"`
	Formatter    string `mapstructure:"formatter"`
	Level        string `mapstructure:"level"`
	ReportCaller bool   `mapstructure:"report-caller"`
}
