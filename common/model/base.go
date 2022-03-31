package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type Page struct {
	Page  int `json:"page"`
	Count int `json:"count"`
}

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
	ID        int64      `gorm:"primary_key" json:"id"`
	CreatedAt int64      `json:"created_at,omitempty"`
	UpdatedAt int64      `json:"updated_at,omitempty"`
	DeletedAt *DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type DeletedAt sql.NullInt64

// Scan implements the Scanner interface.
func (n *DeletedAt) Scan(value interface{}) error {
	return (*sql.NullInt64)(n).Scan(value)
}

// Value implements the driver Valuer interface.
func (n DeletedAt) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int64, nil
}

func (n DeletedAt) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Int64)
	}
	return json.Marshal(nil)
}

func (n *DeletedAt) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		n.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &n.Int64)
	if err == nil {
		n.Valid = true
	}
	return err
}
