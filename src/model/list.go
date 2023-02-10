package model

import (
	"todoapp/src/abstraction"
	"todoapp/pkg/util/date"

	"gorm.io/gorm"
)

type ListEntity struct {
	ListId int `json:"list_id" form:"list_id" gorm:"primaryKey;autoIncrement;"`
	Title   string `json:"title" form:"title" validate:"required" gorm:"not null;type:varchar(100)"`
	Description string `json:"description" form:"description" validate:"required" gorm:"not null;type:varchar(1000)"`
}

type ListFilter struct {
	Title   *string `query:"title" json:"title" filter:"LIKE"`
	Description *string `query:"description" json:"description" filter:"LIKE"`
}

type ListEntityModel struct {
	ListEntity
	abstraction.Entity
	Sublist []SublistEntityModel `json:"sublist" gorm:"foreignKey:ListId"`
	Listfile []ListfileEntityModel `json:"listfile" gorm:"foreignKey:ListId"`
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type ListFilterModel struct {
	//abstraction.Filter
	ListFilter
}

func (ListEntityModel) TableName() string {
	return "list"
}

func (m *ListEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *ListEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}