package youtube

import (
	"github.com/kkdai/youtube/v2"
	"io"
)

type Youtube struct {
	c youtube.Client
}

func New() *Youtube {
	return &Youtube{
		c: youtube.Client{},
	}
}

func (y *Youtube) GetVideo(url string) (*youtube.Video, error) {
	//videoID := "BaW_jenozKc"
	//https://www.youtube.com/watch?v=BaW_jenozKc
	//https://www.youtube.com/watch?v=SzvCuqwD_4A
	return y.c.GetVideo(url)
}

func (y *Youtube) GetVideoFormats(video *youtube.Video, withAudioChannels bool) youtube.FormatList {
	formats := video.Formats
	if withAudioChannels {
		return formats.WithAudioChannels()
	}

	return formats
}

type Stream struct {
	MimeType string
	Title    string
	Size     int
	Reader   io.ReadCloser
}

func (y *Youtube) GetStreamByItag(url string, itag int) (*Stream, error) {
	v, err := y.GetVideo(url)
	if err != nil {
		return nil, err
	}

	format := v.Formats.FindByItag(itag)

	r, size, err := y.c.GetStream(v, format)
	if err != nil {
		return nil, err
	}

	s := &Stream{
		MimeType: format.MimeType,
		Title:    v.Title,
		Size:     int(size),
		Reader:   r,
	}

	return s, nil
}

func (y *Youtube) GetInfo(url string) (*youtube.Video, youtube.FormatList, error) {
	v, err := y.GetVideo(url)
	if err != nil {
		return nil, nil, err
	}

	formats := y.GetVideoFormats(v, true)

	return v, formats, nil
}
