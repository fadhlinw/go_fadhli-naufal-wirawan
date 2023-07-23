package main

import (
	"belajar-go-echo/routes"
)

func main() {
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
