package main

import (
	"AuthInGo/app"
	"fmt"
)

func main() {
	fmt.Println("")
	cfg := app.NewConfig(":3001")
	app := app.NewApplication(*cfg)
	
	app.Run()
}