package serializer

import "ADDD_DOUYIN/model"

type CommentListResponse struct {
	Response
	CommentList []*Comment `json:"comment_list,omitempty"`
}

type CommentResponse struct {
	Response
	Comment `json:"comment"`
}

func PackComment(comment *model.Comment, userId uint) *Comment {
	if comment == nil {
		return nil
	}

	return &Comment{
		Id:         comment.ID,
		User:       *PackUser(&comment.User, userId, true, false),
		Content:    comment.Content,
		CreateDate: comment.CreatedAt.Format("01-02"),
	}
}

func PackComments(cs []*model.Comment, userId uint) []*Comment {
	comments := make([]*Comment, 0)
	for _, c := range cs {
		if comment := PackComment(c, userId); comment != nil {
			comments = append(comments, comment)
		}
	}
	return comments
}
