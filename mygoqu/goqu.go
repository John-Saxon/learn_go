package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
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

//go:generate go run github.com/doug-martin/goqu/v9
func main() {
	// t1 := goqu.Dialect("postgres").From(goqu.S("public").Table("asda")).Prepared(true)
	// stmt, args, err := goqu.Prepared(true).Where(goqu.Ex{
	// 	"a\";": "1;Drop table test;",
	// }).ToSQL()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(stmt)
	// fmt.Println(args)

	t1 := goqu.Dialect("postgres").From(goqu.S("public").Table("action")).Prepared(true)
	stmt, args, err := t1.Where(goqu.C("name").Gt(goqu.V("all'\";DELETE table action where actionid = -5;"))).ToSQL()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(stmt)
	fmt.Println(args)

	conn := InitDB()
	res, err := conn.Exec(stmt, args...)
	if err != nil {
		fmt.Println(err)
		return
	}
	bs, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	stmt, args, err = t1.Insert().Rows(
		goqu.Record{
			"actionid": 0,
			"name":     "test';\"DELETE table action where actionid = -5;",
		},
	).ToSQL()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(stmt)
	fmt.Println(args)

	res, err = conn.Exec(stmt, args...)
	if err != nil {
		fmt.Println(err)
		return
	}
	bs, err = json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	conn.Close()
}
