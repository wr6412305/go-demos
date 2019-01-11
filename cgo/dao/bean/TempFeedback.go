package bean

import "web-demo/cgo/entity"

type TempFeedback struct {
	entity.Feedback
	entity.Picture
}
