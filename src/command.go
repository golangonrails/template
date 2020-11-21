package main

import (
	"fmt"
	"os"

	"app/db/migrations"
	"app/db/seeds"
)

func help() {
	fmt.Println(`
Options:
  db:migrate                            # Migrate Database. Create db (if not exists) and sync tables
  db:seed [seedName]                    # Add Data in SeedName / All Seeds To Database
  db:drop                               # Drop Database
  db:rollback [migrationId]             # Rollback 1 migration or Rollback to migrationId
  help                                  # Show this

Environment Variables:
  APP_DIR: [$PWD]                                           # set program working dir
  CONFIG_FILE: [config.toml]                                # local config file
  CONFIG_CENTER_FILE: [config.center.toml]                  # config center configuration file
  SERVER_ADDR: [0.0.0.0:25250]                              # serv on this addr
  MONITOR_ADDR: [0.0.0.0:38380]                             # monitor on this addr

  AUTO_MIGRATE:                         # set for auto migrate during execute
  AUTO_SEED:                            # set for auto seed during execute
`)
}

func at(arr []string, index int) string {
	if len(arr) > index {
		return os.Args[index]
	}
	return ""
}

func handleCommandArgs() bool {
	if action := at(os.Args, 1); action != "" {
		if f := (map[string]func(){
			"db:migrate": migrations.DbMigration,
			"db:seed": func() {
				seeds.DoSeed(at(os.Args, 2))
			},
			"db:drop": migrations.DbDrop,
			"db:rollback": func() {
				migrations.DbRollback(at(os.Args, 2))
			},
			"help": help,
		})[action]; f != nil {
			f()
		} else {
			fmt.Printf("Command `%v` not found\n", action)
		}
		return true
	}
	return false
}
