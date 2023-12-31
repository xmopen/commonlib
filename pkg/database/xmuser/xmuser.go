package xmuser

import "time"

// XMUser xm user.
type XMUser struct {
	UserAccount  string    `json:"user_account" gorm:"column:user_account"`
	UserID       string    `json:"user_id" gorm:"column:user_id"`
	UserName     string    `json:"user_name" gorm:"column:user_name"`
	UserPassword string    `json:"user_password" gorm:"column:user_password"`
	UserIcon     string    `json:"user_icon" gorm:"column:user_icon"`
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`
	LastLogin    time.Time `json:"last_login" gorm:"column:last_login"`
}

// New a xm user instance.
func New() *XMUser {
	return &XMUser{}
}
