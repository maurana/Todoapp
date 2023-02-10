package abstraction

import (
	"time"
	"todoapp/pkg/util/date"

	"gorm.io/gorm"
)

type Entity struct {
	CreatedAt  time.Time  `json:"created_at"`
	CreatedBy  string     `json:"created_by"`
	ModifiedAt *time.Time `json:"modified_at"`
	ModifiedBy *string    `json:"modified_by"`

	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Filter struct {
	CreatedAt  *time.Time `query:"created_at"`
	CreatedBy  *string    `query:"created_by"`
	ModifiedAt *time.Time `query:"modified_at"`
	ModifiedBy *string    `query:"modified_by"`
}

func (m *Entity) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	return
}

func (m *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	return
}