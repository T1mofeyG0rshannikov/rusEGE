package database

type Task struct {
	Id     uint `json:"id"`
	Number uint `json:"number"`
}

type Word struct {
	Id     uint   `json:"id"`
	TaskId uint   `json:"task_id"`
	Rule   string `json:"rule"`
}

type User struct {
	Id           uint   `json:"id"`
	Username     string `json:"username" gorm:"unique"`
	HashPassword string `json:"hash_password"`
}

type UserWord struct {
	Id     uint `json:"id"`
	UserId uint `json:"user_id"`
	WordId uint `json:"word_id"`
}
