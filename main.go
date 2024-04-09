package main

import (
	"github.com/danieljvx/talent-pitch-api/app"

	"github.com/danieljvx/talent-pitch-api/config"
)

func main() {
	port := config.Config("APP_PORT")
	if len(port) == 0 {
		port = "3000"
	}
	appFiber := app.App()
	err := appFiber.Listen(":" + port)
	if err != nil {
		return
	}

}
