package main

import (
	"service/router"
)

func main() {
	engin := router.Init()
	engin.Run(":8081")
}
