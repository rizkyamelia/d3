package main

import (
	"amelia/gomvc"
	"amelia/myweb/controllers"
)

func main() {
	gomvc.RouteFolder("/", "public")
	gomvc.Route("/proses/", &controllers.ProsesController{})
	gomvc.Run()
}
