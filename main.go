package main

import (
	"jwt_token/database"
	"jwt_token/router"
)


func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run()

}