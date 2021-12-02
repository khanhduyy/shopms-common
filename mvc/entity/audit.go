package entity

import "time"

type Audit struct {
	CreatedBy       uint      `gorm: "column:CreatedBy"`
	CreatedDate     time.Time `gorm: "column:CreatedDate;autoCreateTime;<-:create"`
	LastUpdatedBy   uint      `gorm: "column:LastUpdatedBy"`
	LastUpdatedDate time.Time `gorm: "column:LastUpdatedDate;autoUpdateTime"`
}

func NewAudit(base *Base, userId uint) *Audit {
	if base != nil {
		return &Audit{
			LastUpdatedBy: userId,
		}
	}
	return &Audit{
		CreatedBy:     userId,
		LastUpdatedBy: userId,
	}
}
