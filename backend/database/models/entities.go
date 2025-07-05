package models

type Task struct {
	Id     uint `json:"id"`
	Number uint `json:"number" gorm:"unique"`
	Description string `json:"description"`
}

type Word struct {
	Id     uint   `json:"id"`
	TaskId uint   `json:"task_id"`
	Word   string `json:"word"`
	Rule   *string `json:"rule"`
}

type User struct {
	Id           uint   `json:"id"`
	Username     string `json:"username" gorm:"unique"`
	HashPassword string `json:"hash_password"`
}

type UserWord struct {
	Id      uint   `json:"id"`
	TaskId  uint   `json:"task_id"`
	Word    string `json:"word"`
	Rule    *string `json:"rule"`
	UserId  uint `json:"user_id"`
}
