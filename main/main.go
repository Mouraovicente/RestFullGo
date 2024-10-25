package main

import (
	"RestFullGo/databate"
	"RestFullGo/routes"
)

func main() {
	databate.ConnectDatabase()
	routes.HandleRequest()

}
