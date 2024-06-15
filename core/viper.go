package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"service/config"
	"service/global"
)

func Viper(path ...string) *viper.Viper {
	var configFile string

	if len(path) == 0 {
		pflag.StringVar(&configFile, "c", "", "choose config file.")
		pflag.Parse()

		if configFile == "" {
			if configEnv := os.Getenv(config.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					configFile = config.ConfigDebugFile
				case gin.ReleaseMode:
					configFile = config.ConfigReleaseFile
				case gin.TestMode:
					configFile = config.ConfigTestFile
				default:
					configFile = config.ConfigDefaultFile
				}
			} else {
				configFile = configEnv
			}
		}
	} else {
		configFile = path[0]
	}

	vInstance := viper.New()
	vInstance.SetConfigFile(configFile)
	vInstance.SetConfigType("yaml")
	vInstance.AddConfigPath("./")
	err := vInstance.ReadInConfig()
	if err != nil {
		panic(err)
	}
	vInstance.WatchConfig()
	vInstance.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err = vInstance.Unmarshal(&global.GVA_SETTING); err != nil {
			fmt.Println(err)
		}
	})

	if err = vInstance.Unmarshal(&global.GVA_SETTING); err != nil {
		fmt.Println(err)
	}

	return vInstance
}
