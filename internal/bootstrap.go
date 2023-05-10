package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.con/tgarcia/seed-golang-server/internal/config"
)

type Bootstrap struct {
	Router   *gin.Engine
	Logger   *logrus.Logger
	Registry *prometheus.Registry
}

func NewBootstrap() (*Bootstrap, error) {

	logger := config.SetupLogging()

	requestLatency := config.NewRequestLatency()

	registry := config.SetupPrometheus(requestLatency)

	router := config.SetupGin(requestLatency)

	router.GET("/health", func(c *gin.Context) {
		logger.Info("[health]")
		c.JSON(200, gin.H{
			"message": "Green",
		})
	})

	router.GET("/metrics", gin.WrapH(promhttp.HandlerFor(registry, promhttp.HandlerOpts{})))

	return &Bootstrap{
		Router:   router,
		Logger:   logger,
		Registry: registry,
	}, nil
}

// Start inicia el servidor HTTP.
func (b *Bootstrap) Start() error {
	return b.Router.Run()
}
