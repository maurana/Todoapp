package model

import (
	"todoapp/src/abstraction"
	"todoapp/pkg/util/date"

	"gorm.io/gorm"
)

type SublistEntity struct {
	SublistId int `json:"sublist_id" form:"sublist_id" gorm:"primaryKey;autoIncrement;"`
	Title   string `json:"title" form:"title" validate:"required" gorm:"not null;type:varchar(100)"`
	Description string `json:"description" form:"description" validate:"required" gorm:"not null;type:varchar(1000)"`
}

type SublistFilter struct {
	Title   *string `query:"title" filter:"LIKE"`
	Description *string `query:"description" filter:"LIKE"`
}

type SublistEntityModel struct {
	SublistEntity
	ListId          int        `json:"list_id" form:"list_id"`
	abstraction.Entity
	Sublistfile []SublistfileEntityModel `json:"sublistfile" gorm:"foreignKey:SublistId"`
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type SublistFilterModel struct {
	//abstraction.Filter
	SublistFilter
}

func (SublistEntityModel) TableName() string {
	return "sublist"
}

func (m *SublistEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *SublistEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}