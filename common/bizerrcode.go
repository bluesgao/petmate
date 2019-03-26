package common

var BizOk = BizErrCode{Code: 0, Msg: "ok"}
var BizErrBind = BizErrCode{Code: 10002, Msg: "Error occurred while binding the request body to the struct."}
var BizErrDB = BizErrCode{Code: 20001, Msg: "db error"}

type BizErrCode struct {
	Code int
	Msg  string
}
