package util

import (
	"time"
)

type TimeInterface interface {
	Now() time.Time
}

type RealTime struct {
}

func NewRealTime() TimeInterface {
	return &RealTime{}
}

func (r *RealTime) Now() time.Time {
	return time.Now().UTC()
}
