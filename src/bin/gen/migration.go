package gen

import (
	"fmt"
	"log"
)

func generateMigration(id string) string {
	return fmt.Sprintf(`package migrations

import "github.com/jinzhu/gorm"

func init() {
	AddMigration("%v", func(tx *gorm.DB) (err error) {
		// Do Something Migrate
		return
	}, func(tx *gorm.DB) (err error) {
		// Do Something Rollback
		return
	})
}
`, id)
}

// Migration Gen
func Migration(name string) {
	id := generateID()
	if path, err := generateFile(name, "./db/migrations", id, generateMigration(id)); err != nil {
		log.Printf("Create Migration File Failed: %v\n", path)
	} else {
		log.Printf("Create Migration File Success: %v\n", path)
	}
}
