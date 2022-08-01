package entity

import "io"

type Video struct {
	URL    string
	ItagNo int
}

type Stream struct {
	Size   int
	Title  string
	Reader io.ReadCloser
}

type Info struct {
	ItagNo      string   `json:"itag"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	PublishDate string   `json:"publish_date,string"`
	Thumbnails  string   `json:"thumbnails"`
	FormatList  []Format `json:"format_list"`
}

type Format struct {
	ItagNo           int    `json:"itag"`
	URL              string `json:"url"`
	MimeType         string `json:"mimeType"`
	Quality          string `json:"quality"`
	Cipher           string `json:"signatureCipher"`
	Bitrate          int    `json:"bitrate"`
	FPS              int    `json:"fps"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	LastModified     string `json:"lastModified"`
	ContentLength    int64  `json:"contentLength,string"`
	QualityLabel     string `json:"qualityLabel"`
	ProjectionType   string `json:"projectionType"`
	AverageBitrate   int    `json:"averageBitrate"`
	AudioQuality     string `json:"audioQuality"`
	ApproxDurationMs string `json:"approxDurationMs"`
	AudioSampleRate  string `json:"audioSampleRate"`
	AudioChannels    int    `json:"audioChannels"`
}
