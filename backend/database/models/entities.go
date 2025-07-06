package models

type Task struct {
	Id          uint   `json:"id"`
	Number      uint   `json:"number" gorm:"unique"`
	Description string `json:"description"`
}

type Rule struct {
	Id   uint   `json:"id"`
	Rule string `json:"rule"`
}

type Word struct {
	Id        uint   `json:"id"`
	TaskId    uint   `json:"task_id"`
	Word      string `json:"word"`
	RuleId    uint   `json:"rule_id"`
	Rule      Rule   `gorm:"foreignKey:RuleId"`
	Exception bool   `json:"exception" gorm:"default:false"`
}

type User struct {
	Id           uint   `json:"id"`
	Username     string `json:"username" gorm:"unique"`
	HashPassword string `json:"hash_password"`
}

type UserWord struct {
	Id     uint   `json:"id"`
	TaskId uint   `json:"task_id"`
	Word   string `json:"word"`
	RuleId uint   `json:"rule_id"`
	Rule   Rule   `gorm:"foreignKey:RuleId"`
	UserId uint   `json:"user_id"`
	Exception bool   `json:"exception" gorm:"default:false"`
}
