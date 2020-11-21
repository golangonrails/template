/*
  Database modules
*/
package db

import (
	"app/config"
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type dbLogger struct {
}

func (dbLogger) Print(v ...interface{}) {
	if len(v) > 5 && v[0] == "sql" {
		logger().Printf("(%v) %% %v, effect: %v row(s), spent: %v", v[3], v[4], v[5], v[2])
	} else if len(v) > 2 { // skip tag and file line msg
		logger().Println(v[2:]...)
	} else {
		logger().Println(v...)
	}
}

// NewDbConnection is new db connection, connect to specific db if bindDb is true
func NewDbConnection(bindDB bool) *gorm.DB {
	database := config.Settings().Database
	dbName, dbURL := "", ""
	if bindDB {
		dbName = database.DbName
	}
	switch database.DbType {
	case "postgres":
		if dbName == "" {
			dbName = "postgres"
		}
		dbURL = fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=%v",
			database.UserName, database.Password, database.Server, database.Port, dbName, "disable")
	case "mysql":
		dbURL = fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
			database.UserName, database.Password, database.Server, database.Port, dbName)
	default:
		logger().Fatalf("Invaid dbType: %v\n", database.DbType)
	}
	if config.PRO != config.Env() {
		logger().Printf("Connecting: (%v://%v)", database.DbType, dbURL)
	}
	gormDB, err := gorm.Open(database.DbType, dbURL)
	if err != nil {
		logger().Fatalf("Connect Failed: %s\n", err)
	}
	logger().Println("Connected")
	gormDB.LogMode(true).SetLogger(dbLogger{})
	return gormDB
}

var sGormDB *gorm.DB
var once sync.Once

// Instance is singleton DB connection
func Instance() *gorm.DB {
	once.Do(func() {
		sGormDB = NewDbConnection(true)
	})
	return sGormDB
}

// Create DB
func Create() {
	db := NewDbConnection(false)
	defer db.Close()

	dbConf := config.Settings().Database
	db.Exec(fmt.Sprintf("CREATE DATABASE %v;", dbConf.DbName))
}

// Drop DB
func Drop() {
	db := NewDbConnection(false)
	defer db.Close()

	dbConf := config.Settings().Database
	db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %v;", dbConf.DbName))
}
