package logs

import (
	"os"
	"path/filepath"

	"github.com/jiebiubiu/base/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const DefaultLogPath = "~/data/jarvis/log" // 默认输出日志文件路径

var logger *zap.SugaredLogger

func GetLogger() *zap.SugaredLogger {
	return logger
}

var DefaultLogConfigs = config.Log{
	LogLevel:          "debug",
	LogFormat:         "",
	LogPath:           "./logs",
	LogFileName:       "logs.log",
	LogFileMaxSize:    5,
	LogFileMaxBackups: 10,
	LogMaxAge:         10,
	LogCompress:       false,
	LogStdout:         true,
}

// InitLogger 初始化 log
func InitLogger(logConfs ...config.Log) error {
	var conf config.Log = DefaultLogConfigs
	if len(logConfs) > 0 {
		conf = logConfs[0]
	}

	logLevel := map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
	}
	writeSyncer, err := getLogWriter(conf) // 日志文件配置 文件位置和切割
	if err != nil {
		return err
	}
	encoder := getEncoder(conf)          // 获取日志输出编码
	level, ok := logLevel[conf.LogLevel] // 日志打印级别
	if !ok {
		level = logLevel["info"]
	}

	core := zapcore.NewCore(encoder, writeSyncer, level)

	_logger := zap.New(core, zap.AddCaller()) // zap.Addcaller() 输出日志打印文件和行数如： logger/logger_test.go:33
	// 1. zap.ReplaceGlobals 函数将当前初始化的 logger 替换到全局的 logger,
	// 2. 使用 logger 的时候 直接通过 zap.S().Debugf("xxx") or zap.L().Debug("xxx")
	// 3. 使用 zap.S() 和 zap.L() 提供全局锁，保证一个全局的安全访问logger的方式
	zap.ReplaceGlobals(_logger)
	//zap.L().Debug("")
	//zap.S().Debugf("")

	logger = zap.S()

	return nil
}

// getEncoder 编码器(如何写入日志)
func getEncoder(conf config.Log) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // log 时间格式 例如: 2021-09-11t20:05:54.852+0800
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 输出level序列化为全大写字符串，如 INFO DEBUG ERROR
	// encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	if conf.LogFormat == "json" {
		return zapcore.NewJSONEncoder(encoderConfig) // 以json格式写入
	}
	return zapcore.NewConsoleEncoder(encoderConfig) // 以logfmt格式写入
}

// getLogWriter 获取日志输出方式  日志文件 控制台
func getLogWriter(conf config.Log) (zapcore.WriteSyncer, error) {
	// 判断日志路径是否存在，如果不存在就创建
	if exist := IsExist(conf.LogPath); !exist {
		if conf.LogPath == "" {
			conf.LogPath = DefaultLogPath
		}
		if err := os.MkdirAll(conf.LogPath, os.ModePerm); err != nil {
			conf.LogPath = DefaultLogPath
			if err := os.MkdirAll(conf.LogPath, os.ModePerm); err != nil {
				return nil, err
			}
		}
	}

	// 日志文件 与 日志切割 配置
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(conf.LogPath, conf.LogFileName), // 日志文件路径
		MaxSize:    conf.LogFileMaxSize,                           // 单个日志文件最大多少 mb
		MaxBackups: conf.LogFileMaxBackups,                        // 日志备份数量
		MaxAge:     conf.LogMaxAge,                                // 日志最长保留时间
		Compress:   conf.LogCompress,                              // 是否压缩日志
	}
	if conf.LogStdout {
		// 日志同时输出到控制台和日志文件中
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout)), nil
	} else {
		// 日志只输出到日志文件
		return zapcore.AddSync(lumberJackLogger), nil
	}
}

// IsExist 判断文件或者目录是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
