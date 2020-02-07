package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"

	"./bot"
	"./db"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func main() {
	botToken := os.Getenv("TOKEN")
	errChan := make(chan error)
	go bot.StartBot(botToken, errChan)
	// var user = db.User{101, "Hello", "Moto"}
	// err := db.InsertUserIntoDB(user)
	// if err != nil {
	// 	panic(err)
	// }
	index := db.Index{101, "Hellwwwwo", "Moto", "Hell", time.Now()}
	dbSQL, err := sql.Open("sqlite3", "./db.sqlite3")
	defer dbSQL.Close()
	if err != nil {
		panic(err)
	}
	err = db.AddIndex(dbSQL, index)
	if err != nil {
		panic(err)
	}
	go logger(errChan)
}

func logger(errChan <-chan error) {
	for {
		err := <-errChan
		fmt.Println(err.Error())
	}
}
