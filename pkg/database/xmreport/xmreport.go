// Package xmreport Package xmreport
package xmreport

import "time"

// XMReport report
type XMReport struct {
	ID         int       `json:"id" gorm:"column:id"`
	IncrValue  int       `json:"incr_value" gorm:"column:incr_value"`
	ReportTime time.Time `json:"report_time" gorm:"column:report_time"`
	Desc       string    `json:"desc" gorm:"column:desc"`
	Value      string    `json:"value" gorm:"column:value"`
	ReportIP   string    `json:"report_ip" gorm:"column:report_ip"`
}

// New return a xmreport
func New() *XMReport {
	return &XMReport{}
}
