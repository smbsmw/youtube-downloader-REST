package usecase

import (
	"context"
	"youtube-downloader-rest/internal/entity"
)

type DownloadUseCase struct {
	extractor Extractor
}

func New(e Extractor) *DownloadUseCase {
	return &DownloadUseCase{
		extractor: e,
	}
}

func (d *DownloadUseCase) Download(ctx context.Context, v entity.Video) (*entity.Stream, error) {
	s, err := d.extractor.GetStream(v)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (d *DownloadUseCase) GetInfo(ctx context.Context, v entity.Video) (*entity.Info, error) {
	return d.extractor.GetInfo(v)
}
