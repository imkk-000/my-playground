package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Context struct {
	router         *gin.Engine
	specificRouter *gin.RouterGroup
}

func api(c *gin.Context) {
	log.Println("Do api")
}

func main() {
	ctx := &Context{}
	r := gin.New()
	ctx.router = r

	r.Use(func(c *gin.Context) {
		log.Println("Apply this middleware for all api")
	})
	r.GET("/test", api)

	rg := r.Group("/")
	ctx.specificRouter = rg
	rg.Use(func(c *gin.Context) {
		log.Println("Apply this middleware for specific api")
	})
	rg.GET("/sptest", api)

	r.Run()
}
