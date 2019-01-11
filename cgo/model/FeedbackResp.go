package model

import "web-demo/cgo/entity"

type FeedbackResp struct {
	entity.Feedback
	Pictures []entity.Picture
}
