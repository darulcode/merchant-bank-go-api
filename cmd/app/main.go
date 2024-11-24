package main

import (
	"mncTest/internal/app/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
