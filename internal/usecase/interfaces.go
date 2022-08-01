package usecase

import (
	"context"
	"youtube-downloader-rest/internal/entity"
)

type (
	Download interface {
		Download(context.Context, entity.Video) (*entity.Stream, error)
		GetInfo(context.Context, entity.Video) (*entity.Info, error)
	}

	Extractor interface {
		GetStream(entity.Video) (*entity.Stream, error)
		GetInfo(entity.Video) (*entity.Info, error)
	}
)
