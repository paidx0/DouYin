package global

import "DouYin/Enter/config"

var (
	CONFIG   config.Server // 配置信息
	ZAP      = InitZap()   // 操作Zap日志
	VIPER    = InitViper() // 操作viper
	DBEngine = InitDB()    // 操作mysql
	REDIS    = InitRedis() // 操作redis
)

var (
	Success int32 = 0 // 0-成功
	Error   int32 = 1 // 其他失败
)

var (
	LogDir = "./../log"
)
