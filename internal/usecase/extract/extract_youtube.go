package extract

import (
	"youtube-downloader-rest/internal/entity"
	"youtube-downloader-rest/pkg/youtube"
)

type Extract struct {
	client *youtube.Youtube
}

func New() *Extract {
	return &Extract{
		client: youtube.New(),
	}
}

func (e Extract) GetStream(video entity.Video) (*entity.Stream, error) {
	s, err := e.client.GetStreamByItag(video.URL, video.ItagNo)
	if err != nil {
		return nil, err
	}

	return &entity.Stream{
		Size:   s.Size,
		Title:  s.Title,
		Reader: s.Reader,
	}, nil
}

func (e Extract) GetInfo(video entity.Video) (*entity.Info, error) {
	v, fl, err := e.client.GetInfo(video.URL)
	if err != nil {
		return nil, err
	}

	formats := make([]entity.Format, 0)

	for _, f := range fl {
		if f.QualityLabel != "" {
			formats = append(formats, entity.Format{
				ItagNo:   f.ItagNo,
				URL:      f.URL,
				FPS:      f.FPS,
				Quality:  f.QualityLabel,
				MimeType: f.MimeType,
			})
		}
	}

	return &entity.Info{
		Title:       v.Title,
		Description: v.Description,
		Author:      v.Author,
		Thumbnails:  v.Thumbnails[0].URL,
		FormatList:  formats,
	}, nil
}
