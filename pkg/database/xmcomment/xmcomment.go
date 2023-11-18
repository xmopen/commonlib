package xmcomment

import "time"

// XMComment comment structural
type XMComment struct {
	ID         int       `json:"id" gorm:"column:id"`
	ArticleID  int       `json:"article_id" gorm:"column:article_id"` // ArticleID 评论所属文章ID
	Status     int       `json:"status" gorm:"column:status"`         // Status 0：不被用户所看到 1：允许用户能看到
	Account    string    `json:"account" gorm:"column:account"`       // Account 评论所属用户
	Content    string    `json:"content" gorm:"column:content"`       // Content 评论内容
	City       string    `json:"city" gorm:"column:city"`             // City 城市
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}

// NewComment return a xm comment instance
func NewComment(articleID, status int, account, content, city string) *XMComment {
	return &XMComment{
		ArticleID: articleID,
		Status:    status,
		Account:   account,
		Content:   content,
		City:      city,
	}
}
