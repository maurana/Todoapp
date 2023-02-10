package model

import (
	"todoapp/src/abstraction"
	"todoapp/pkg/util/date"

	"gorm.io/gorm"
)

type SublistfileEntity struct {
	SublistfileId int `json:"sublistfile_id" gorm:"primaryKey;autoIncrement;"`
	Filename   string `json:"filename" validate:"required" gorm:"not null;type:varchar(255)"`
}

type SublistfileEntityModel struct {
	SublistfileEntity
	SublistId          int        `json:"sublist_id"`
	abstraction.Entity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (SublistfileEntityModel) TableName() string {
	return "sublistfile"
}

func (m *SublistfileEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *SublistfileEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}