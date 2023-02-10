package main

import (
	"testing"
	"todoapp/src/server"
	db "todoapp/database"
)

func testmain(t *testing.T) {
	db.InitMySql()
	db.InitPostgreSql()
	server.Init()
}