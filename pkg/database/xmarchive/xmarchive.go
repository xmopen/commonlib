package xmarchive

// XMBlogsArchive xm blogs archive
type XMBlogsArchive struct {
	ID      int    `json:"id" gorm:"column:id"`
	Name    string `json:"name" gorm:"column:name"`
	BKStyle string `json:"bk_style" gorm:"column:bk_style"`
	Weight  int    `json:"weight" gorm:"column:weight"`
	Status  int8   `json:"status" gorm:"column:status"`
}

// New return an archive instance
func New() *XMBlogsArchive {
	return &XMBlogsArchive{}
}
