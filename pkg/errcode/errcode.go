// Package errcode  错误码.
package errcode

// CommonResult result structural.
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

	// ErrorXMUserTokenExpired 登录态失效
	ErrorXMUserTokenExpired = &CommonResult{Code: 10011, Msg: "登录态失效"}

	// ErrorSystemError 系统错误
	ErrorSystemError = &CommonResult{Code: 10012, Msg: "系统错误"}

	// ErrorGORPCError GORPC error
	ErrorGORPCError = &CommonResult{Code: 10013, Msg: "GORPC Request Error"}

	// ErrorParseJSONError 解析Json失败
	ErrorParseJSONError = &CommonResult{Code: 10014, Msg: "解析Json失败"}

	// ErrorUserLoginFail 用户登录失败.
	ErrorUserLoginFail = &CommonResult{Code: 20000, Msg: "用户登录失败,请重试！"}

	// ErrorUserRegisterFail 用户注册失败.
	ErrorUserRegisterFail = &CommonResult{Code: 20001, Msg: "用户注册失败,请重试！"}

	// ErrorUserAccountIllegal 用户账号或密码不合法.
	ErrorUserAccountIllegal = &CommonResult{Code: 20002, Msg: "用户账号或密码不合法,请重新尝试！"}

	// ErrorUserAccountHaveAlreadyRegistered 用户账号已经注册
	ErrorUserAccountHaveAlreadyRegistered = &CommonResult{Code: 20003, Msg: "用户已经注册，请稍后重试"}

	// ErrorGetIndexArticleList 获取文章错误.
	ErrorGetIndexArticleList = &CommonResult{Code: 30010, Msg: "获取文章列表出错"}

	// ErrorCreateArticleError 文章评论失败
	ErrorCreateArticleError = &CommonResult{Code: 30011, Msg: "评论失败，请稍后重试"}
)
