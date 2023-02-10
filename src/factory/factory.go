package factory

import (
	"todoapp/database"
	"todoapp/src/repository"

	"gorm.io/gorm"
)

type Factory struct {
	Db               *gorm.DB
	ListRepository repository.List
	SublistRepository repository.Sublist
}

func NewFactory() *Factory {
	f := &Factory{}
	f.SetupDb()
	f.SetupRepository()

	return f
}

func (f *Factory) SetupDb() {
	db, err := database.Connection("todo")
	if err != nil {
		panic("Failed setup db, connection is undefined")
	}
	f.Db = db
}

func (f *Factory) SetupRepository() {
	if f.Db == nil {
	    panic("Failed setup repository, db is undefined")
	}

	f.ListRepository = repository.NewList(f.Db)
	f.SublistRepository = repository.NewSublist(f.Db)
}