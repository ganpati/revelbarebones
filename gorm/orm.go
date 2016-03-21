// This file provides the init functions to setup ORM support for this application

package gorm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"os"
	_ "reflect"
)

var (
	dbHandle    gorm.DB
	initialized = false
)

func GetDbHandle() gorm.DB {
	if initialized == false {
		initializeDbHandle()
	}
	return dbHandle
}

func initializeDbHandle() {

	//db, err := gorm.Open("postgres", "user=gumbo dbname=gorm sslmode=disable")
	// db, err := gorm.Open("foundation", "dbname=gorm") // FoundationDB.
	db, err := gorm.Open("mysql", "gkalele@/test?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("sqlite3", "/tmp/gorm.db")

	// You can also use an existing database connection handle
	// dbSql, _ := sql.Open("postgres", "user=gorm dbname=gorm sslmode=disable")
	// db, _ := gorm.Open("postgres", dbSql)

	if err != nil {
		fmt.Println("Error opening database connection", err)
		os.Exit(-1)
	}

	// Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
	db.DB()

	// Then you could invoke `*sql.DB`'s functions with it
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	dbHandle = db
	initialized = true

	// Disable table name's pluralization
	db.SingularTable(true)

	fmt.Println("Initialized DB")
}
