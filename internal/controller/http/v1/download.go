package v1

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"youtube-downloader-rest/internal/entity"
	"youtube-downloader-rest/internal/usecase"
	"youtube-downloader-rest/pkg/logger"
)

type downloadRoutes struct {
	downloadUsecase usecase.Download
	logger          logger.Interface
}

func newDownloadRoutes(handler *echo.Group, d usecase.Download, l logger.Interface) {
	r := &downloadRoutes{
		d,
		l,
	}

	h := handler.Group("/api")
	{
		h.GET("/download", r.doDownload)
		h.POST("/info", r.doGetInfo)
	}
}

type DownloadRequest struct {
	URL  string `query:"url"`
	Itag int    `query:"itag"`
}

// @Summary     Download video
// @Description Download video
// @Param 		url query string true "URL of youtube video"
// @Param 		itag query string true "Itag of youtube video"
// @ID          download
// @Accept      json
// @Produce     json
// @Router      /v1/api/download [get]
func (r *downloadRoutes) doDownload(c echo.Context) error {
	var payload DownloadRequest
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &payload); err != nil {
		r.logger.Error(err, "http - v1 - download")

		return echo.NewHTTPError(http.StatusBadRequest, "Please provide valid data")
	}

	vs, err := r.downloadUsecase.Download(c.Request().Context(), entity.Video{
		URL:    payload.URL,
		ItagNo: payload.Itag,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Response().Header().Set(
		echo.HeaderContentLength, strconv.Itoa(vs.Size),
	)

	c.Response().Header().Set(
		echo.HeaderContentDisposition,
		fmt.Sprintf("attachment; filename=%s.mp4", vs.Title),
	)

	return c.Stream(http.StatusOK, "video/mp4", vs.Reader)
}

type GetInfoRequest struct {
	URL string `json:"url"`
}

// @Summary     Get video info
// @Description Get information of video
// @Param 		url query string true "URL of youtube video"
// @ID          get-info
// @Accept      json
// @Produce     json
// @Router      /v1/api/info [post]
func (r *downloadRoutes) doGetInfo(c echo.Context) error {
	var payload *GetInfoRequest

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return err
	}

	info, err := r.downloadUsecase.GetInfo(c.Request().Context(), entity.Video{
		URL: payload.URL,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, info)
}
