package main

import (
	"Go-Upload-Rest-Api-Gin/Database"
	"Go-Upload-Rest-Api-Gin/Database/Migrations"
	"Go-Upload-Rest-Api-Gin/Routes"
)

func main() {
	Database.ConnectionDatabase()
	Migrations.Migration()

	// Replace if want to use api or web
	// Routes.Route()		// web
	Routes.RouteApi()		// api
	// =================
}