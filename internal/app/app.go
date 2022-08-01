package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"os/signal"
	"syscall"
	"youtube-downloader-rest/config"
	v1 "youtube-downloader-rest/internal/controller/http/v1"
	"youtube-downloader-rest/internal/usecase"
	"youtube-downloader-rest/internal/usecase/extract"
	"youtube-downloader-rest/pkg/httpserver"
	"youtube-downloader-rest/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Usecase
	downloadUsecase := usecase.New(extract.New())

	// HTTP Server
	e := echo.New()
	v1.NewRouter(e, downloadUsecase, l)
	httpServer := httpserver.New(e, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
