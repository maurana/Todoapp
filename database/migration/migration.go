package migration

import (
	"fmt"
	"todoapp/database"
	"todoapp/src/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Migration interface {
	AutoMigrate()
	SetDb(*gorm.DB)
}

type migration struct {
	Db            *gorm.DB
	DbModels      *[]interface{}
	IsAutoMigrate bool
}

func Init() {
	mgConfigurations := map[string]Migration{
		"TODO": &migration{
			DbModels: &[]interface{}{
				&model.ListEntityModel{},
				&model.SublistEntityModel{},
				&model.ListfileEntityModel{},
				&model.SublistfileEntityModel{},
			},
			IsAutoMigrate: true,
		},
	}

	for k, v := range mgConfigurations {
		dbConnection, err := database.Connection(k)
		if err != nil {
			logrus.Error(fmt.Sprintf("Failed to run migration, database not found %s", k))
		} else {
			v.SetDb(dbConnection)
			v.AutoMigrate()
			logrus.Info(fmt.Sprintf("Successfully run migration for database %s", k))
		}
	}

}

func (m *migration) AutoMigrate() {
	if m.IsAutoMigrate {
		m.Db.AutoMigrate(*m.DbModels...)
	}
}

func (m *migration) SetDb(db *gorm.DB) {
	m.Db = db
}