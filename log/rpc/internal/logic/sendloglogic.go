package logic

import (
	"context"

	"autotrigger/log/rpc/internal/svc"
	"autotrigger/log/rpc/types/log"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogLogic {
	return &SendLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  发送日志
func (l *SendLogLogic) SendLog(in *log.MyErrModel) (*log.MyErrModel, error) {
	l.svcCtx.Logger.DealLogInfo(in)
	return &log.MyErrModel{}, nil
}
