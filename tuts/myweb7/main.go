package main

import (
	"amelia/gomvc"
	// "amelia/myweb/controllers"
)

func main() {
	gomvc.SetConfig("port", "9778")
	gomvc.RouteFolder("/", "public")
	// gomvc.Route("/proses/", &controllers.ProsesController{})
	gomvc.Run()
}
