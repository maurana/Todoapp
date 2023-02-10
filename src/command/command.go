package command

import (
	"os"
	"log"
	"todoapp/src/server"
	"todoapp/database/migration"
	db "todoapp/database"
	gen "todoapp/pkg/constant"

	cli "github.com/urfave/cli/v2"
)

func Init() {
	app := cli.NewApp()
	app.Name = gen.APP_NAME
	app.Description = gen.APP_DESC
	app.Commands = []*cli.Command{
		{
			Name:        "start-postgres",
			Description: "Start Echo server with connected PostgreSQL",
			Action: func(c *cli.Context) error {
				db.InitPostgreSql()
				server.Init()
				return nil
			},
		},
		{
			Name:        "start-mysql",
			Description: "Start Echo server with connected MySQL",
			Action: func(c *cli.Context) error {
				db.InitMySql()
				server.Init()
				return nil
			},
		},
		{
			Name:        "migrate-postgres",
			Description: "Migration to PostgreSQL",
			Action: func(c *cli.Context) error {
				db.InitPostgreSql()
				migration.Init()
				return nil
			},
		},
		{
			Name:        "migrate-mysql",
			Description: "Migration to MySQL",
			Action: func(c *cli.Context) error {
				db.InitMySql()
				migration.Init()
				return nil
			},
		},
		{
			Name:        "launch-postgres",
			Description: "Start Echo server and Migration to PostgreSQL",
			Action: func(c *cli.Context) error {
				db.InitPostgreSql()
				migration.Init()
				server.Init()
				return nil
			},
		},
		{
			Name:        "launch-mysql",
			Description: "Start Echo server and Migration to MySQL",
			Action: func(c *cli.Context) error {
				db.InitMySql()
				migration.Init()
				server.Init()
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}