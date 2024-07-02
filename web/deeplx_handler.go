package web

import (
	"deeplx-local/domain"
	"deeplx-local/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeepLXHandler struct {
	service   service.TranslateService
	routePath string
}

func NewDeepLXHandler(service service.TranslateService, customRoute string) *DeepLXHandler {
	if customRoute == "" {
		customRoute = "/translate"
	}
	if customRoute[0] != '/' {
		customRoute = "/" + customRoute
	}
	return &DeepLXHandler{service: service, routePath: customRoute}
}
func (d *DeepLXHandler) Translate(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight request
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	var request domain.TranslateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	translatedText := d.service.GetTranslateData(request)
	c.JSON(http.StatusOK, translatedText)
}

func (d *DeepLXHandler) RegisterRoutes(engine *gin.Engine) {
	engine.OPTIONS(d.routePath, d.Translate)
	engine.POST(d.routePath, d.Translate)
}
