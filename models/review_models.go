package models

import (
	"time"
)

// TODO: add appropriate struct tags

type Review struct {
	id              uint32
	rating          uint8
	difficulty      uint8
	interest        uint8
	classId         uint32
	date            time.Time
	anonymous       bool
	text            string
	professorName   string
	helpfulCount    uint32
	notHelpfulCount uint32
}
