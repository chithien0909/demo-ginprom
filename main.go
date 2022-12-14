package main

import (
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	r := gin.Default()

	registry := prometheus.NewRegistry() // creates new prometheus metric registry

	p := ginprom.New(
		ginprom.Registry(registry),
		ginprom.Engine(r),
		ginprom.Subsystem("ginprom"),
		ginprom.Namespace("my_prometheus"),
		ginprom.Path("/api/metrics"),
		ginprom.Ignore("/api/no/no/no", "/hello/:id"),
		ginprom.Token("supersecrettoken"),
	)
	p.AddCustomGauge("custom", "Some help text to provide", []string{"label"})
	r.Use(p.Instrument())

	r.GET("/hello/:id", func(c *gin.Context) {})
	r.GET("/world/:id", func(c *gin.Context) {})
	r.Run("127.0.0.1:8080")
}
