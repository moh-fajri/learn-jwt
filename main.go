package main

import (
	"fmt"
	"log"
	"os"

	"github.com/moh-fajri/learn-jwt/repo"
	"github.com/moh-fajri/learn-jwt/repo/mysql"
	"github.com/moh-fajri/learn-jwt/route"
	"github.com/moh-fajri/learn-jwt/util"

	"github.com/joho/godotenv"
)

func main() {
	// load ev
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// connect mysql
	err = mysql.InitDatabase()
	if err != nil {
		log.Fatal("Error Connect Database")
	}
	// migrate database
	err = mysql.DB.AutoMigrate(
		&repo.User{},
		&repo.Product{},
	)
	if err != nil {
		log.Fatal("Migrate Failed")
	}
	// connect route api
	e := route.Init()
	data, err := util.Json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		panic(fmt.Sprint(err))
	}
	fmt.Println(string(data))

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
