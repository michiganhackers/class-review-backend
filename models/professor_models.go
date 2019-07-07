package models

type Professor struct {
	Uniqname string `db:"professor_uniqname" json:"uniqname"`
	Name 	 string	`db:"professor_name" json:"name"`
}

type ProfessorStats struct {
	Uniqname 			 string	 `db:"professor_uniqname"`
	AvgRating          	 float32 `db:"AVG(rating)"`
	AvgDifficulty      	 float32 `db:"AVG(difficulty)"`
	AvgInterest        	 float32 `db:"AVG(interest)"`
	TotalHelpfulCount    uint64  `db:"SUM(helpfulCount)"`
	TotalNotHelpfulCount uint64	 `db:"SUM(notHelpfulCount)"`
}