package main

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"net/http"
)

func main() {
	router := gin.Default()

	metrics := ginmetrics.GetMonitor()
	metrics.SetMetricPath("/metrics")
	metrics.Use(router)

	router.GET("/api/v1/gin", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from Go!")
	})

	router.Run(":8000")
}
