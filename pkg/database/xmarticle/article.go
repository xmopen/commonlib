package xmarticle

// Article 文章.
type Article struct {
	ID        int    `json:"id" gorm:"column:id"`
	TypeID    int    `json:"type_id" gorm:"column:type_id"` // TypeID 分类ID.
	Publish   int    `json:"publish" gorm:"column:publish"` // Publish 是否发布.
	ReadCount int    `json:"read_count" gorm:"column:read_count"`
	Title     string `json:"title" gorm:"column:title"`
	Time      string `json:"time" gorm:"column:time"`
	Author    string `json:"author" gorm:"column:author"`
	Content   string `json:"content" gorm:"column:content"`
	SubHead   string `json:"sub_head" gorm:"column:subhead"`
	Img       string `json:"img" gorm:"column:img"`
	Type      string `json:"type" gorm:"column:type"`
}

// New return a article instance
func New() *Article {
	return &Article{}
}
