package Migrations

import (
	"fmt"
	"Go-Upload-Rest-Api-Gin/Models"
	"Go-Upload-Rest-Api-Gin/Database"
)

func Migration() {
	err := Database.DB.AutoMigrate(
		&Models.Upload{},
	)

	if err != nil {
		fmt.Println("Can't Running Migrations")
		return
	}

	fmt.Println("Migrations Success")
}