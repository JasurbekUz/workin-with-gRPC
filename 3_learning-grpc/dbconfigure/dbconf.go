package dbconfigure

import (
	"fmt"
	"log"
	"database/sql"
	_"github.com/lib/pq"
)

const (
	
	Host = "localhost"
	DUser = "jasurbek"
	Password = "1001"
	DbName = "example"
	Port = "5432"
	SslMode = "disable"
	TimeZone = "Asia/Tashkent"
)


var DB_CONFIG = fmt.Sprintf(

		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		Host, DUser, Password, DbName, Port, SslMode, TimeZone,
	)

func DataBase () *sql.DB {

	db, err := sql.Open("postgres", DB_CONFIG)

	if err != nil {
		log.Fatalf("db connection error: %v", err)
	}

	return db
}	