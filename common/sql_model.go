package common

import "time"

type SQLModel struct {
	Id        int        `json:"id" gorm:"column:id"` // tag name json, json encode
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}
