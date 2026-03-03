package global

import (
	"os"
	"strconv"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	Log *logrus.Logger
)

var Config = struct {
	// 系统
	ListenHost  string
	ListenPort  int32
	ApiRootPath string
	// 数据库
	DbHost string
	DbPort int32
	DbUser string
	DbPwd  string
	DbName string
	// 日志
	IsSaveLog bool
	LogPath   string
	LogLevel  logrus.Level
	// 基础验证
	VerifyCode string
	// 统计天数
	ChartDay int64
}{
	ListenHost: "0.0.0.0",
	ListenPort: 8002,
	DbHost:     getEnv("DB_HOST", "127.0.0.1"),
	DbPort:     getEnvInt32("DB_PORT", "3306"),
	DbUser:     getEnv("DB_USER", "root"),
	DbPwd:      getEnv("DB_PWD", "123456"),
	DbName:     getEnv("DB_NAME", "wow_hong"),
	IsSaveLog:  true,
	LogPath:    "./logs/log.txt",
	LogLevel:   logrus.DebugLevel,
	VerifyCode: "testcode",
	ChartDay:   20,
}

func init() {
	Log = logrus.New()
	Log.Level = Config.LogLevel
	if Config.IsSaveLog {
		pathMap := lfshook.PathMap{
			logrus.InfoLevel:  Config.LogPath,
			logrus.ErrorLevel: Config.LogPath,
			logrus.WarnLevel:  Config.LogPath,
			logrus.DebugLevel: Config.LogPath,
		}
		Log.Hooks.Add(lfshook.NewHook(
			pathMap,
			&logrus.JSONFormatter{},
		))
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt32(key, defaultValue string) int32 {
	if value := os.Getenv(key); value != "" {
		return 3306
	}
	atoi, err := strconv.Atoi(getEnv("DB_PORT", "3306"))
	if err != nil {
		return 3306
	}
	return int32(atoi)
}
