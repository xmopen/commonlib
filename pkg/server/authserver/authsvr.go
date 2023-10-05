package authserver

import (
	"context"

	"github.com/xmopen/commonlib/pkg/database/xmuser"
)

// AuthSvrName auth svr name
const AuthSvrName = "AuthSvr"

// AuthSvrMethodName auth svr method name type
type AuthSvrMethodName string

const (
	// AuthSvrMethodTypeOfGetUserInfoByXMAccount get user info by account method name
	AuthSvrMethodTypeOfGetUserInfoByXMAccount AuthSvrMethodName = "GetUserInfoByAccount"
	// AuthSvrMethodTypeOfGetUserInfoByXMToken get user info by token method name
	AuthSvrMethodTypeOfGetUserInfoByXMToken AuthSvrMethodName = "GetUserInfoByToken"
)

// AuthSvrRequest auth svr req
type AuthSvrRequest struct {
	XMAccount string `json:"xm_account"`
	XMToken   string `json:"xm_token"`
}

// AuthSvrResponse auth svr rsp
type AuthSvrResponse struct {
	XMUserInfo *xmuser.XMUser `json:"xm_user_info"`
}

// IAuthServer auth svr
type IAuthServer interface {
	// GetUserInfoByAccount get user info from account
	GetUserInfoByAccount(ctx context.Context, request *AuthSvrRequest, response *AuthSvrResponse) error
	// GetUserInfoByToken get user info from token
	GetUserInfoByToken(ctx context.Context, request *AuthSvrRequest, response *AuthSvrResponse) error
}
