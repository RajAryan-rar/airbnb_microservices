package main

import (
	"AuthInGo/app"
	"fmt"
)

func main() {
	fmt.Println("")
	cfg := app.Config {
		Addr: ":3001",
	}
	app := app.Application {
		Config: cfg,
	}
	
	app.Run()
}