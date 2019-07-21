package models

type Professor struct {
	Uniqname string `db:"professor_uniqname" json:"uniqname"`
	Name 	 string	`db:"professor_name" json:"name"`
}

type ProfessorStats struct {
	Uniqname 			 string	 `db:"professor_uniqname" json:"uniqname"`
	AvgRating          	 float32 `db:"AVG(rating)" json:"avgRating"`
	AvgDifficulty      	 float32 `db:"AVG(difficulty)" json:"avgDifficulty"`
	AvgInterest        	 float32 `db:"AVG(interest)" json:"avgInterest"`
	TotalHelpfulCount    uint64  `db:"SUM(helpfulCount)" json:"totalHelpfulCount"`
	TotalNotHelpfulCount uint64	 `db:"SUM(notHelpfulCount)" json:"totalNotHelpfulCount"`
}