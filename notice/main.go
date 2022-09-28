package main

import (
	"notice/router"
)

func main() {
	engin := router.Init()
	engin.Run(":8082")
}
