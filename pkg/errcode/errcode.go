// Package errcode  错误码.
package errcode

// CommonResult result struct.
type CommonResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(data any) *CommonResult {
	return &CommonResult{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
}

var (
	/*******通用错误码********/

	// ErrorParam 请求参数错误.
	ErrorParam = &CommonResult{Code: 10010, Msg: "请求参数错误"}

	// ErrorGetIndexArticleList 获取文章错误.
	ErrorGetIndexArticleList = &CommonResult{Code: 10010, Msg: "获取文章列表出错"}

	// ErrorUserLoginFail 用户登录失败.
	ErrorUserLoginFail = &CommonResult{Code: 20000, Msg: "用户登录失败,请重试！"}

	// ErrorUserRegisterFail 用户注册失败.
	ErrorUserRegisterFail = &CommonResult{Code: 20001, Msg: "用户注册失败,请重试！"}
)
