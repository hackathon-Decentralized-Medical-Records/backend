## 项目结构

```shell
├── api
│   └── v1
├── config
├── core
├── docs
├── eth
│   ├── abi
│   ├── contracts
├── global
├── initialize
├── logslogrus.log
├── main.go
├── middleware
├── model
│   ├── request
│   └── response
├── router
├── service
└── utils
```

| 文件夹       | 说明                    | 描述                           |
| ------------ | ----------------------- |------------------------------|
| `api`        | api层                   | api层                         |
| `--v1`       | v1版本接口              | v1版本接口                       |
| `config`     | 配置包                  | config.yaml的配置结构体            |
| `core`       | 核心文件                | 核心组件(logrus, viper, server)的初始化 |
| `docs`       | swagger文档目录         | swagger文档目录                  |
| `global`     | 全局对象                | 全局对象                         |
| `initialize` | 初始化 | router,redis,gorm初始化         |
| `middleware` | 中间件层 | `gin` 中间件代码                |
| `model`      | 模型层                  | 数据表的模型                       |
| `--request`  | 入参结构体              | 前端->后端的数据结构体                 |
| `--response` | 出参结构体              | 后端->前端的数据结构体                 |
| `router`     | 路由层                  | 路由层                          |
| `service`    | service层               | 业务逻辑                         |
| `utils`      | 工具包                  | 工具函数封装                       |


## 项目规划

- WebSocket构建聊天系统
- Redis构建缓存
- 引入AI机器人
