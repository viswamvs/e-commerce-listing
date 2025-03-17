package main

import (
	"e-commerce-listing/routes"
	"e-commerce-listing/utils/cache"
	"e-commerce-listing/utils/db"
)

func main() {

	db.InitializeDB()

	redisClient := cache.InitializeCache()

	r := routes.GetRouter(redisClient)
	r.Run(":8080")
}
