package database

import (
	"database/sql"
	"log"

	configurations "alshashiguchi/quiz_gem/core"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var Db *sql.DB

//InitDB - initializes the database
func InitDB(config *configurations.Configurations) {

	db, err := sql.Open(config.DataBase.Drive, config.DataBase.URL)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	Db = db
}

//Migrate - runs migrations
func Migrate(config *configurations.Configurations) {

	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql.WithInstance(Db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		config.DataBase.PathMigrations,
		config.DataBase.Drive,
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

}
