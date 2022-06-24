package myerr

import "autotrigger/log/rpc/types/log"

var (
	Success = log.MyErrModel{
		Code:    0,
		Message: "success",
	}
)
