package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"service/config"
)

var (
	GVA_DATABASE *gorm.DB
	GVA_VIPER    *viper.Viper
	GVA_SETTING  config.Setting
	GVA_LOGRUS   *logrus.Logger
	GVA_REDIS    redis.UniversalClient
)
