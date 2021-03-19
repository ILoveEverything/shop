package main

import (
	"shop/db"
	"shop/router"
)

func main() {
	r := router.InitRouter()
	r.Run(":9999")
	defer db.RedisDB.Close()
}
