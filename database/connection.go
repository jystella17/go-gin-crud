package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func Connection()(*sql.DB,error) {
	db,err := sql.Open("mysql","root:theo93@go_gin_crud")
	if err != nil{
		log.Fatalf("DB Connection Error : %s",err)
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	defer db.Close()

	ping := db.Ping()
	if err != nil{
		log.Fatalf("Ping Error")
		panic(ping.Error())
	}
	fmt.Println(db)
	return db,nil
}