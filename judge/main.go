package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	initRouter(router)

	panic(router.Run(":8080"))
}
