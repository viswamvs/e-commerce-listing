package main

import (
	"e-commerce-listing/routes"
	"e-commerce-listing/utils/db"
)

func main() {
	
	db.InitializeDB()

	r := routes.GetRouter()
	r.Run(":8080")
}
