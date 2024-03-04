package handler

import (
	"net/http"
	"time"

	"code.stakefish.test/service/ip_validator/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

type Handler struct {
	queries service.QueriesService
}

func NewHandler(services service.QueriesService) *Handler {
	return &Handler{queries: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("", h.version)
	router.GET("/health", h.health)

	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	m.Use(router)

	v1Api := router.Group("/v1")
	{
		v1Api.GET("/history", h.history)
	}

	tools := v1Api.Group("/tools")
	tools.POST("validate", h.validate)
	tools.POST("lookup", h.lookUp)

	return router
}

func (h *Handler) version(c *gin.Context) {
	// todo for now just hard-coding
	type CurrentVersion struct {
		Version    string `json:"version"`
		Date       int64  `json:"date"`
		Kubernetes bool   `json:"kubernetes"`
	}

	c.JSON(http.StatusOK, &CurrentVersion{Version: "0.1.0", Date: time.Now().UnixMilli(), Kubernetes: false})
}

func (h *Handler) health(c *gin.Context) {
	c.JSON(http.StatusOK, `{}`)
}
