package common

var Ok = ErrCode{Code: 0, Msg: "ok"}
var RedisError = ErrCode{Code: 10101, Msg: "redis error"}
var RequstBindError = ErrCode{Code: 20101, Msg: "Error occurred while binding the request body to the struct."}

/*
1系统级
2服务级
错误代码说明20101
2		01			01
错误级别	服务模块代码	具体错误代码
*/
type ErrCode struct {
	Code int
	Msg  string
}
