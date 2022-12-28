package main

import (
	"dc/router"
)

func main() {
	engin := router.Init()
	engin.Run(":8084")
}
