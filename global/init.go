package global

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
	"xorm.io/xorm"
)

// InitDB 初始化数据库
func InitDB() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", CONFIG.DB.DBUser, CONFIG.DB.DBPasswd, CONFIG.DB.Path, CONFIG.DB.DBName, CONFIG.DB.Path))
	if err != nil {
		ZAP.Warn("数据库初始化失败", zap.Error(err))
		return nil
	}
	return engine
}

// InitRedis 初始化Redis
func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     CONFIG.Redis.Addr,
		Password: CONFIG.Redis.Password, // no password set
		DB:       CONFIG.Redis.DB,       // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		ZAP.Warn("Redis初始化失败", zap.Error(err))
		return nil
	}
	return rdb
}

// InitViper 初始化viper
func InitViper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("./config/config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		ZAP.Warn("viper读取配置失败", zap.Error(err))
		return nil
	}

	// 监控配置文件重新读取
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("global file changed:", e.Name)
		if err := v.Unmarshal(&CONFIG); err != nil {
			ZAP.Warn("viper监控更新失败", zap.Error(err))
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&CONFIG); err != nil {
		ZAP.Warn("viper初始化失败", zap.Error(err))
		return nil
	}
	return v
}

// InitZap 初始化Zap
func InitZap() *zap.Logger {
	// 先要检查LogDir文件夹存不存在
	if exist := PathExists(LogDir); !exist {
		fmt.Printf("创建日志文件夹Path：%v\n", LogDir)
		os.Mkdir(LogDir, os.ModePerm)
	}

	logger := zap.New(getEncoderCore())
	// 附带文件名和行号返回日志
	logger = logger.WithOptions(zap.AddCaller())
	return logger
}

func getEncoderCore() (core zapcore.Core) {
	writer, err := GetWriteSyncer()
	if err != nil {
		fmt.Printf("日志写入失败：%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, zap.InfoLevel)
}

// GetWriteSyncer 日志分割
func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(LogDir, "%Y-%m-%d.log"),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
}

// 输出格式 json|console，默认 console
func getEncoder() zapcore.Encoder {
	if CONFIG.ZAP.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stackracekey",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	return config
}

// CustomTimeEncoder 日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05"))
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
