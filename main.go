package main

import (
	"golang-echo-rest-api/db"
	"golang-echo-rest-api/routes"
)

func main() {
	db.Init()
	
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1234"))
}
