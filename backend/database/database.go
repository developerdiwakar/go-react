package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Db struct {
	DB *gorm.DB
}

// Connect and load the database
func Load() *Db {
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASS := os.Getenv("MYSQL_PASS")
	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_PORT := os.Getenv("MYSQL_PORT")
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")
	MYSQL_CHARSET := os.Getenv("MYSQL_CHARSET")
	MYSQL_PARSETIME := os.Getenv("MYSQL_PARSETIME")

	dbString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=Local",
		MYSQL_USER,
		MYSQL_PASS,
		MYSQL_HOST,
		MYSQL_PORT,
		MYSQL_DATABASE,
		MYSQL_CHARSET,
		MYSQL_PARSETIME,
	)

	// Open the database connection
	dbConn, err := gorm.Open(mysql.Open(dbString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "golist_",
		},
	})
	if err != nil {
		log.Fatalln("Failed to connect database", err)
	}
	log.Println("Successfully connected to database")
	return &Db{DB: dbConn}
}

// Close the database connection
func (d *Db) Close() {
	mysqlDB, err := d.DB.DB()
	if err != nil {
		log.Fatalln("Database closing error", err)
	}
	mysqlDB.Close()
}
