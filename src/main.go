package main

import (
	"fmt"
	"os"

	"app/config"
	"app/db"
	"app/db/migrations"
	"app/db/seeds"
	"app/monitor"
	"app/routers"
	"app/utils/logs"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init Logger
	if config.Env() != config.DEV {
		os.MkdirAll("./log", os.ModePerm)
		logs.LogToFile("./log/app.log")
		gin.DefaultWriter = logs.DefaultWriter
		gin.DefaultErrorWriter = logs.DefaultWriter
		gin.SetMode(gin.ReleaseMode)
	}

	// Handle Command line params
	if handleCommandArgs() {
		return
	}
	fmt.Println("[HINT] Command 'help' for more options")

	// Start Monitor
	{
		addr := "0.0.0.0:38380"
		if v := os.Getenv("MONITOR_ADDR"); v != "" {
			addr = v
		}
		go func() {
			if err := monitor.Serve(addr); err != nil {
				fmt.Printf("\n[MONITOR] Start Failed: %v\n", err.Error())
				os.Exit(0)
			}
		}()
		fmt.Printf("\n[MONITOR] On http://%v/\n\n", addr)
	}

	// Do db migrate atomatically if env is set
	if config.Settings().Environment.AutoMigrate || os.Getenv("AUTO_MIGRATE") != "" {
		migrations.DbMigration()
	}

	// Do db seed atomatically if env is set
	if config.Settings().Environment.AutoSeed || os.Getenv("AUTO_SEED") != "" {
		seeds.DoSeed("")
	}

	logger().Printf("Starting with ENV='%v' ...", config.Env())

	// Init Db
	db.Instance()

	router := gin.Default()
	routers.Init(router)

	addr := "0.0.0.0:25250"
	if v := os.Getenv("SERVER_ADDR"); v != "" {
		addr = v
	}
	fmt.Printf("\n[SERVER] On http://%v/\n\n", addr)
	router.Run(addr)
}
