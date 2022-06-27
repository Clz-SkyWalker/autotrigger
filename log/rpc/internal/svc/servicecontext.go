package svc

import (
	"autotrigger/log/rpc/internal/config"
	"autotrigger/log/rpc/types/proto"
	"go.uber.org/zap"
)

type ServiceContext struct {
	Config config.Config
	Logger *LoggerStruct
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Logger: NewLoggerStruct(c.LogConfig),
	}
}

func NewLoggerStruct(logConfig config.LogConfig) *LoggerStruct {
	return &LoggerStruct{Logger: config.InitLog(logConfig)}
}

type LoggerStruct struct {
	Logger *zap.SugaredLogger
}

//
//  DealLogInfo
//  @Description: 日志处理
//  @receiver l
//  @param in
//
func (l *LoggerStruct) DealLogInfo(in *proto.MyErrModel) {
	switch in.Level {
	case proto.MyErrModel_DebugLevel:
		l.Logger.Debugf("code:%d,msg:%s,params:%s", in.Code, in.Message, in.Params)
	case proto.MyErrModel_InfoLevel:
		l.Logger.Infof("code:%d,msg:%s,params:%s", in.Code, in.Message, in.Params)
	case proto.MyErrModel_ErrorLevel:
		l.Logger.Errorf("code:%d,msg:%s,params:%s,file:%s,method:%s", in.Code, in.Message, in.Params, in.Stack, in.MethodName)
	case proto.MyErrModel_PanicLevel:
	}
}
