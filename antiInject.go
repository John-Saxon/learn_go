package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	host := "localhost"
	user := "oushu"
	dbName := "test"
	port := "4432"
	sslMode := "disable"
	password := "123"
	var connectionString = fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s password=%s", host, user, dbName, port, sslMode, password) // sslmode=require
	log.Println("InitDB():Initializing postgres database: " + connectionString)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func main() {
	conn := InitDB()
	res, err := conn.Exec(`SELECT * FROM action WHERE @col = @value`, sql.Named("col", "name"), sql.Named("value", "select"))
	if err != nil {
		fmt.Println(err)
		return
	}
	bs, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
