package models

type Professor struct {
	Name string	`db:"professor_name"`
}

type ProfessorStats struct {
	Name 				 string	 `db:"professor_name"`
	AvgRating          	 float32 `db:"AVG(rating)"`
	AvgDifficulty      	 float32 `db:"AVG(difficulty)"`
	AvgInterest        	 float32 `db:"AVG(interest)"`
	TotalHelpfulCount    uint64  `db:"SUM(helpfulCount)"`
	TotalNotHelpfulCount uint64	 `db:"SUM(notHelpfulCount)"`
}