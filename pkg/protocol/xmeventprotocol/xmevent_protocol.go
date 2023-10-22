package xmeventprotocol

// XMEventKey  事件广播Key类型
type XMEventKey string

// 事件广播Key
const (
	// XMEventKeyOfArticleUpdate 文章更新事件
	XMEventKeyOfArticleUpdate XMEventKey = "xm_event_article_update"
	// XMEventKeyOfArticleCommentUpdate 文章评论更新事件
	XMEventKeyOfArticleCommentUpdate XMEventKey = "xm_event_article_comment_update"
	// XMEventKeyOfArchiveUpdate 文章分类更新
	XMEventKeyOfArchiveUpdate XMEventKey = "xm_event_archive_update"
)
