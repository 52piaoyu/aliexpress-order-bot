package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"./bot"
	"./db"
)

func main() {
	botToken := os.Getenv("TOKEN")
	errChan := make(chan error)
	go bot.StartBot(botToken, errChan)
	// var user = db.User{101, "Hello", "Moto"}
	// err := db.InsertUserIntoDB(user)
	// if err != nil {
	// 	panic(err)
	// }
	var index = db.Index{101, "Hellwwwwo", "Moto", "Hell", time.Now()}
	dbSql, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}
	err = db.AddIndex(dbSql, index)
	if err != nil {
		panic(err)
	}
	logger(errChan)
}

func logger(errChan <-chan error) {
	for {
		err := <-errChan
		fmt.Println(err.Error())
	}
}
