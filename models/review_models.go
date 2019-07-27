package models

import (
    "time"
)

// TODO: add appropriate struct tags

type Review struct {
    Id                  uint64    `db:"id" json:"id"`
    Rating              uint8     `db:"rating" json:"rating"`
    Difficulty          uint8     `db:"difficulty" json:"difficulty"`
    Interest            uint8     `db:"interest" json:"interest"`
    CourseId            uint64    `db:"courseId" json:"courseId"`
    Date                time.Time `db:"review_date" json:"date"`
    Anonymous           bool      `db:"is_anonymous" json:"anonymous"`
    Text                *string   `db:"review_text" json:"text"`
    ProfessorUniqname   *string   `db:"professor_uniqname" json:"professorUniqname"`
    HelpfulCount        uint64    `db:"helpfulCount" json:"helpfulCount"`
    NotHelpfulCount     uint64    `db:"notHelpfulCount" json:"notHelpfulCount"`
    Semester            *string   `db:"semester" json:"semester"`
    UserEmail           *string   `db:"userEmail" json:"userEmail"`
}
