package config

type Redis struct {
	Addr         string   `mapstructure:"addr"`         // 服务器地址:端口
	Password     string   `mapstructure:"password"`     // 密码
	DB           int      `mapstructure:"db"`           // 单实例模式下redis的哪个数据库
	UseCluster   bool     `mapstructure:"useCluster" `  // 是否使用集群模式
	ClusterAddrs []string `mapstructure:"clusterAddrs"` // 集群模式下的节点地址列表
}
