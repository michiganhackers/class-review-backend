package models

import (
	"time"
)

// TODO: add appropriate struct tags

type Review struct {
	id              uint64
	rating          uint8
	difficulty      uint8
	interest        uint8
	classId         uint64
	date            time.Time
	anonymous       bool
	text            string
	professorName   string
	helpfulCount    uint64
	notHelpfulCount uint64
}
