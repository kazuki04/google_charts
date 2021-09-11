package main

import (
	"fmt"
	"google_charts/app/controllers"
	"google_charts/config"
)

func main() {
	fmt.Println(config.Config.Port)
	controllers.StartWebServer()
}
