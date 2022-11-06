package cmd

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"log"
	"myclass_service/src/config"
	"os"
)

func setupM(config config.IConfig) *migrate.Migrate {
	db, err := config.GetDB().DB()
	if err != nil {
		log.Fatalln("err connect to db")
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	dir, _ := os.Getwd()
	path := "file://" + dir + "/src/database/migrations"

	m, err := migrate.NewWithDatabaseInstance(path, "mysql", driver)

	if err != nil {
		fmt.Println("err migrate up 1 ", err)
	}

	return m
}

func migrateUp(config config.IConfig) *cobra.Command {
	return &cobra.Command{

		Use: "migrate-up",
		Run: func(cmd *cobra.Command, args []string) {

			m := setupM(config)

			err := m.Up()

			if err != nil {
				fmt.Println("err migrate up 2", err)
			}
		},
	}
}

func migrateDown(config config.IConfig) *cobra.Command {
	return &cobra.Command{
		Use: "migrate-down",
		Run: func(cmd *cobra.Command, args []string) {

			m := setupM(config)

			err := m.Steps(-1)

			if err != nil {
				fmt.Println("err migrate up 2", err)
			}
		},
	}
}

func migrateRefresh(config config.IConfig) *cobra.Command {
	return &cobra.Command{
		Use: "migrate-refresh",
		Run: func(cmd *cobra.Command, args []string) {

			m := setupM(config)

			err := m.Down()
			err = m.Up()

			if err != nil {
				fmt.Println("err migrate up 2", err)
			}
		},
	}
}
