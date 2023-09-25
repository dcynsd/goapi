package models

import (
	"github.com/golang-module/carbon/v2"
)

type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id"`
}

type CommonTimestampsField struct {
	CreatedAt carbon.DateTime `gorm:"column:created_at;" json:"created_at,omitempty"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at;" json:"updated_at,omitempty"`
}
