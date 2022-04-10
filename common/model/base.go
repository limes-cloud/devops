package model

import (
	"gorm.io/gorm"
)

type CreateModel struct {
	ID        int64 `gorm:"primary_key" json:"id"`
	CreatedAt int64 `json:"created_at,omitempty"`
}

type BaseModel struct {
	ID        int64 `gorm:"primary_key" json:"id"`
	CreatedAt int64 `json:"created_at,omitempty"`
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

type DeleteModel struct {
	ID        int64          `gorm:"primary_key" json:"id"`
	CreatedAt int64          `json:"created_at,omitempty"`
	UpdatedAt int64          `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
