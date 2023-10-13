package xmarchive

// XMBlogsArchive xm blogs archive
type XMBlogsArchive struct {
	ID         int    `json:"id" gorm:"column:id"`
	Name       string `json:"name" gorm:"column:name"`
	BKStyle    string `json:"bk_style" gorm:"column:bk_style"`
	Weight     int    `json:"weight" gorm:"column:weight"`
	NodeType   string `json:"node_type" gorm:"column:node_type"`
	NodeColor  string `json:"node_color" gorm:"column:node_color"`
	NodeIcon   string `json:"node_icon" gorm:"column:node_icon"`
	NodeHollow string `json:"node_hollow" gorm:"column:node_hollow"`
	Status     int8   `json:"status" gorm:"column:status"`
}

// New return an archive instance
func New() *XMBlogsArchive {
	return &XMBlogsArchive{}
}
