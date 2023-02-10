package model

import (
	"todoapp/src/abstraction"
	"todoapp/pkg/util/date"

	"gorm.io/gorm"
)

type ListfileEntity struct {
	ListfileId int `json:"listfile_id" gorm:"primaryKey;autoIncrement;"`
	Filename   string `json:"filename" validate:"required" gorm:"not null;type:varchar(255)"`
}

type ListfileEntityModel struct {
	ListfileEntity
	ListId          int        `json:"list_id"`
	abstraction.Entity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (ListfileEntityModel) TableName() string {
	return "listfile"
}

func (m *ListfileEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *ListfileEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}