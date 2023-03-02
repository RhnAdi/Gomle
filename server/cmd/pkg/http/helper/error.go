package helper

import "errors"

var (
	ErrYouAreNotOwner error = errors.New("you aren't own this post")
)
