package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func SetupGin(requestLatency *RequestLatency) *gin.Engine {

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		start := time.Now()
		requestLatency.Observe(c.Request.Method, c.Request.URL.Path, time.Since(start).Seconds())
		logrus.Infof("%s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
		logrus.Infof("Status code: %d", c.Writer.Status())

	})

	return router
}
