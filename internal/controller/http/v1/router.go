package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"youtube-downloader-rest/internal/usecase"
	"youtube-downloader-rest/pkg/logger"

	echoSwagger "github.com/swaggo/echo-swagger"
	// Swagger docs
	_ "youtube-downloader-rest/docs"
)

// @title Youtube Download REST
// @version 1.0
// @description This is a sample server.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func NewRouter(handler *echo.Echo, d usecase.Download, l logger.Interface) {
	// Options
	handler.Use(middleware.Logger())
	handler.Use(middleware.Recover())

	// Swagger
	handler.GET("/swagger/*", echoSwagger.WrapHandler)

	// index.html
	handler.File("/", "public/index.html")

	// Routes
	h := handler.Group("/v1")
	{
		newDownloadRoutes(h, d, l)
	}
}
