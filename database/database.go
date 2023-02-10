package database

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	dbConnections map[string]*gorm.DB
	dbConfigurations map[string]Db
)

func InitPostgreSql() {
	dbConfigurations = map[string]Db{
		"TODO": &dbPostgreSQL{
			db: db{
				Host: os.Getenv("POSTGRESQL_HOST"),
				User: os.Getenv("POSTGRESQL_USER"),
				Pass: os.Getenv("POSTGRESQL_PASS"),
				Port: os.Getenv("POSTGRESQL_PORT"),
				Name: os.Getenv("POSTGRESQL_NAME"),
			},
			SslMode: os.Getenv("POSTGRESQL_SSLMODE"),
			Tz:      os.Getenv("POSTGRESQL_TZ"),
		},
	}

    Configure()
}

func InitMySql() {
	dbConfigurations = map[string]Db{
		"TODO": &dbMySQL{
			db: db{
				Host: os.Getenv("MYSQL_HOST"),
				User: os.Getenv("MYSQL_USER"),
				Pass: os.Getenv("MYSQL_PASS"),
				Port: os.Getenv("MYSQL_PORT"),
				Name: os.Getenv("MYSQL_NAME"),
			},
			Charset: os.Getenv("MYSQL_CHARSET"),
			ParseTime: os.Getenv("MYSQL_PARSETIME"),
			Loc: os.Getenv("MYSQL_LOC"),
		},
	}

	Configure()
}

func Configure() {
	dbConnections = make(map[string]*gorm.DB)
	for k, v := range dbConfigurations {
		db, err := v.Init()
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to database %s", k))
		}
		dbConnections[k] = db
		logrus.Info(fmt.Sprintf("Successfully connected to database %s", k))
	}
}

func Connection(name string) (*gorm.DB, error) {
	if dbConnections[strings.ToUpper(name)] == nil {
		return nil, errors.New("Connection is undefined")
	}
	return dbConnections[strings.ToUpper(name)], nil
}