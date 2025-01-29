package main

import (
	"golang-echo-rest-api/routes"
)

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1234"))
}
