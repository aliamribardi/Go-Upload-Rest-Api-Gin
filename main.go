package main

import (
	"Go-Upload-Rest-Api-Gin/Database"
	"Go-Upload-Rest-Api-Gin/Database/Migrations"
	"Go-Upload-Rest-Api-Gin/Routes"
)

func main() {
	Database.ConnectionDatabase()
	Migrations.Migration()
	Routes.Route()
	Routes.RouteApi()
}