package gen

import (
	"app/db/seeds"
	"fmt"
	"log"
)

func generateSeed(id, name string) string {
	return fmt.Sprintf(`package seeds

import "github.com/jinzhu/gorm"

func init() {
	AddSeed("%v", "%v", func(tx *gorm.DB) (err error) {
		// Do Something Seed
		return
	})
}
`, id, name)
}

// Seed Gen
func Seed(name string) {
	if seeds.HasSeed(name) {
		log.Printf("Seed `%v` already exists\n", name)
		return
	}
	id := generateID()
	if path, err := generateFile(name, "./db/seeds", id, generateSeed(id, name)); err != nil {
		log.Printf("Create Seed File Failed: %v\n", path)
	} else {
		log.Printf("Create Seed File Success: %v\n", path)
	}
}
