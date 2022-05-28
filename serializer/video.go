package serializer

import "ADDD_DOUYIN/model"

//序列化的FeedResponse
type FeedResponse struct {
	Response
	VideoList []*Video `json:"video_list,omitempty"`
	NextTime  int64    `json:"next_time,omitempty"`
}

func PackVideo(v *model.Video) *Video {
	if v == nil {
		return nil
	}

	return &Video{
		Id:            v.ID,
		Author:        *PackUser(&v.Author, false),
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}
}

func PackVideos(vs []*model.Video) []*Video {
	videos := make([]*Video, 0)
	for _, v := range vs {
		if v2 := PackVideo(v); v2 != nil {
			videos = append(videos, v2)
		}
	}
	return videos
}
