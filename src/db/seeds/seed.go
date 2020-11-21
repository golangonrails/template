/*
  Database initialize data define here, generate by codegen
*/
package seeds

import (
	"app/db"
	"log"

	"github.com/jinzhu/gorm"
)

type dbAction func(tx *gorm.DB) error

type seed struct {
	ID     string
	Name   string
	Action dbAction
}

func (s *seed) Do() (err error) {
	if err = db.Instance().Transaction(s.Action); err != nil {
		log.Printf("[ERROR] Do Seed `%v` (id:%v) Failed: %v\n", s.Name, s.ID, err)
	} else {
		log.Printf("[SEED] Do Seed `%v` (id:%v) Success\n", s.Name, s.ID)
	}
	return
}

var seedsMap = make(map[string]*seed)
var seedsList []*seed

func AddSeed(id, name string, action dbAction) {
	s := &seed{id, name, action}
	seedsList = append(seedsList, s)
	seedsMap[name] = s
}

func DoSeed(name string) {
	if name != "" {
		if s, ok := seedsMap[name]; ok {
			s.Do()
		} else {
			log.Printf("[ERROR] Seed `%v` Not found\n", name)
			return
		}
	} else {
		for _, s := range seedsList {
			if s.Do() != nil {
				return
			}
		}
	}
	log.Printf("[SEED] All Done\n")
}

func HasSeed(name string) bool {
	_, ok := seedsMap[name]
	return ok
}
