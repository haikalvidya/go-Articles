package main

import (
	"goarticle/app"
	"goarticle/app/database"
	"log"
	"net/http"
	"os"
	_ "github.com/lib/pq"
)

func main() {
	apps := app.New()
	apps.DB = &database.DB{}
	err := apps.DB.Open()
	check(err)

	defer apps.DB.Close()

	http.HandleFunc("/", apps.Router.ServeHTTP)

	log.Println("App running...")

	err = http.ListenAndServe(":9000", nil)
	check(err)
}

func check(err error){
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
