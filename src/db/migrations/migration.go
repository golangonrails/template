/*
  Database table struct migrations definition
*/
package migrations

import (
	"app/db"
	"app/utils/logs"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var logger = logs.Logger("[MIGRATIONS] ")

var migrations []*gormigrate.Migration

func AddMigration(id string, migrate func(tx *gorm.DB) error, rollback func(tx *gorm.DB) error) {
	migrations = append(migrations, &gormigrate.Migration{ID: id, Migrate: migrate, Rollback: rollback})
}

func goMigrate(db *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(db, gormigrate.DefaultOptions, migrations)
}

// DbMigration is connect to db, do create db and migrate tables, then disconnect
func DbMigration() {
	db.Create()
	dbConnect := db.NewDbConnection(true)
	defer dbConnect.Close()
	m := goMigrate(dbConnect)
	if err := m.Migrate(); err != nil {
		logger().Fatalf("Could not migrate: %v", err)
	}
	logger().Printf("Migration did run successfully")
}

// DbRollback is connect to db, rollback migrate, then disconnect
func DbRollback(id string) {
	dbConnect := db.NewDbConnection(true)
	defer dbConnect.Close()
	m := goMigrate(dbConnect)
	var err error
	if id == "" {
		err = m.RollbackLast()
	} else {
		err = m.RollbackTo(id)
	}
	if err != nil {
		logger().Fatalf("Could not rollback: %v", err)
	}
	logger().Printf("Rollback did run successfully")
}

// DbDrop is drop db
func DbDrop() {
	if os.Getenv("NO_DATABASE_PROTECT") == "" {
		fmt.Println("Set env `NO_DATABASE_PROTECT` to confirm drop database")
		return
	}
	db.Drop()
}
