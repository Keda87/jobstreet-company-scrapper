package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type DBHelper struct {
	DB *sql.DB
}

func CreateDB() *DBHelper {
	file, err := os.Create("companies.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	sqliteDB, err := sql.Open("sqlite3", "./companies.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	createDDL := `
	CREATE TABLE IF NOT EXISTS companies (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"company_name" TEXT,
		"company_size" TEXT,
		"industry" TEXT,
		"address" TEXT,
		"map_address" TEXT,
		"map_latitude" TEXT,
		"map_longitude" TEXT
	)
	`
	statement, err := sqliteDB.Prepare(createDDL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	defer statement.Close()

	return &DBHelper{
		DB: sqliteDB,
	}
}

func (db *DBHelper) Save(data *Company) {
	sql := `
		INSERT INTO companies (
			company_name, 
			company_size, 
			industry, 
			address, 
			map_address,
			map_latitude,
			map_longitude
		) VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	statement, err := db.DB.Prepare(sql)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()

	statement.Exec(
		data.CompanyName,
		data.CompanySize,
		data.Industry,
		data.Address,
		data.MapAddress,
		data.MapLatitude,
		data.MapLongitude,
	)
}
