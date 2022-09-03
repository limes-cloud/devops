package main

import (
	"configure/router"
)

func main() {
	engin := router.Init()
	engin.Run(":8081")
}
