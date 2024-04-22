package xmreport

import "time"

// XMReportVisit 访问。
type XMReportVisit struct {
	ID        int       `json:"id" gorm:"column:id"`
	Path      string    `json:"path" gorm:"column:path"`
	IP        string    `json:"ip" gorm:"column:ip"`
	VisitTime time.Time `json:"visit_time" gorm:"column:visit_time"`
}
