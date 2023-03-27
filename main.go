package main

import (
	"ginchat/router"
)

func main() {
	r := router.Router()
	r.Run(":9999")
}
