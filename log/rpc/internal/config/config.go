package config

import (
	"io"
	"strings"
	"time"

	rotatelog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/zeromicro/go-zero/zrpc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	zrpc.RpcServerConf
	LogConfig LogConfig `json:"logConfig"`
}

type LogConfig struct {
	MaxAge       int `json:"maxAge,default=604800"`
	RotationTime int `json:"rotationTime,default=86400"`
}

func InitLog(config LogConfig) *zap.SugaredLogger {
	encoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.InfoLevel
	})
	errLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.ErrorLevel
	})
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(getWriter("./logs/info.log", config.MaxAge, config.RotationTime)), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(getWriter("./logs/error.log", config.MaxAge, config.RotationTime)), errLevel),
	)
	log := zap.New(core, zap.AddCaller())
	return log.Sugar()
}

//
//  getWriter
//  @Description: 分割日志文件
//  @param fileName 文件名
//  @param maxAge 文件最多存在多少 秒
//  @param rotationTime 文件分割时间段 秒
//  @return io.Writer
//
func getWriter(fileName string, maxAge, rotationTime int) io.Writer {
	hook, err := rotatelog.New(
		strings.Replace(fileName, ".log", "", -1)+"-%Y%m%d%H.log",
		rotatelog.WithLinkName(fileName),
		rotatelog.WithMaxAge(time.Second*time.Duration(maxAge)),
		rotatelog.WithRotationTime(time.Second*time.Duration(rotationTime)),
	)
	if err != nil {
		panic(err)
	}
	return hook
}
