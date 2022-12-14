package main

import (
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	p := ginprom.New(
		ginprom.Engine(r),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)
	p.AddCustomGauge("custom", "Some help text to provide", []string{"label"})
	r.Use(p.Instrument())

	r.GET("/hello/:id", func(c *gin.Context) {})
	r.GET("/world/:id", func(c *gin.Context) {})
	r.Run("127.0.0.1:8080")
}
