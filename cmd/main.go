package main

import (
	"qsoft/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	handlers.RegisterRoutes(r)
	r.Run()
}
