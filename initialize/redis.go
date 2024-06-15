package initialize

import (
	"context"
	"github.com/redis/go-redis/v9"
	"service/global"
)

func Redis() {
	redisCfg := global.GVA_SETTING.Redis
	var client redis.UniversalClient
	// 使用集群模式
	if redisCfg.UseCluster {
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    redisCfg.ClusterAddrs,
			Password: redisCfg.Password,
		})
	} else {
		// 使用单例模式
		client = redis.NewClient(&redis.Options{
			Addr:     redisCfg.Addr,
			Password: redisCfg.Password,
			DB:       redisCfg.DB,
		})
	}
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GVA_LOGRUS.Error("redis connect ping failed, err:" + err.Error())
		panic(err)
	} else {
		global.GVA_LOGRUS.Info("redis connect ping response:", pong)
		global.GVA_REDIS = client
	}
}
