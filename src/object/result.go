package object

import (
	"golang.org/x/xerrors"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func DataResult(data interface{}) Result {
	result := Result{Data: data, Code: 0, Msg: "成功"}
	return result
}

func EmptyResult() Result {
	result := Result{Data: nil, Code: 0, Msg: "成功"}
	return result
}

func FailedResult(exception Exception) Result {
	result := Result{Data: nil, Code: exception.Code, Msg: exception.Error()}
	return result
}

func CheckException(err error) Result {
	if xerrors.As(err, &Exception{}) {
		return FailedResult(err.(Exception))
	} else {
		return FailedResult(UnknownError)
	}
}
