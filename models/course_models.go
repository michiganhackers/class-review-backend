package models

import (
	"time"
)

// TODO: add appropriate struct tags

type Course struct {
	id         uint64
	department string
	number     uint8
	subsection uint8
	title      string
	date       time.Time
	reviews    []Review
}
