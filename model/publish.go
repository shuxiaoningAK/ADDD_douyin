package model

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}
