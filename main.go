package main

import "shop/router"

func main() {
	r := router.InitRouter()
	r.Run(":9999")
}
