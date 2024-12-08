package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		panic("DB_DSN environment variable not set")
	}

	log.Printf("Connecting to database with DSN: %s", dsn)

	// Register MySQL driver
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		panic(fmt.Sprintf("Failed to register driver: %v", err))
	}

	// Register default database
	err = orm.RegisterDataBase("default", "mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Failed to register default database: %v", err))
	}

	// Set database timezone
	err = orm.SetDataBaseTZ("default", time.UTC)
	if err != nil {
		panic(fmt.Sprintf("Failed to set database timezone: %v", err))
	}

	log.Println("Database connection successfully initialized")
}
