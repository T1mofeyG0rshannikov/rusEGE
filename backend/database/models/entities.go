package models

type Task struct {
	Id          uint   `json:"id"`
	Number      uint   `json:"number" gorm:"unique"`
	Description string `json:"description"`
}

type Rule struct {
	Id      uint         `json:"id"`
	Rule    string       `json:"rule"`
	Options []RuleOption `json:"options"`
}

type Word struct {
	Id          uint         `json:"id"`
	TaskId      uint         `json:"task_id"`
	Task        Task         `gorm:"foreignKey:TaskId"`
	Word        string       `json:"word" gorm:"unique"`
	RuleId      uint         `json:"rule_id"`
	Rule        Rule         `gorm:"foreignKey:RuleId"`
	Exception   bool         `json:"exception" gorm:"default:false"`
	Description *string      `json:"description"`
	Options     []WordOption `json:"options"`
}

type WordOption struct {
	Id     uint   `json:"id"`
	WordId uint   `json:"word_id"`
	Word   Word   `gorm:"foreignKey:WordId"`
	Letter string `json:"letter"`
}

type User struct {
	Id           uint   `json:"id"`
	Username     string `json:"username" gorm:"unique"`
	HashPassword string `json:"hash_password"`
}

type UserWord struct {
	Id          uint             `json:"id"`
	TaskId      uint             `json:"task_id"`
	Task        Task             `gorm:"foreignKey:TaskId"`
	Word        string           `json:"word" gorm:"unique"`
	RuleId      uint             `json:"rule_id"`
	Rule        Rule             `gorm:"foreignKey:RuleId"`
	UserId      uint             `json:"user_id"`
	User        User             `gorm:"foreignKey:UserId"`
	Exception   bool             `json:"exception" gorm:"default:false"`
	Description *string          `json:"description"`
	Options     []UserWordOption `json:"options" gorm:"foreignKey:WordId"`
}

type UserWordOption struct {
	Id     uint     `json:"id"`
	WordId uint     `json:"word_id"`
	Word   UserWord `gorm:"foreignKey:WordId"`
	Letter string   `json:"letter"`
}

type Error struct {
	Id     uint `json:"id"`
	WordId uint `json:"word_id"`
	Word   Word `gorm:"foreignKey:WordId"`
}

type UserError struct {
	Id     uint     `json:"id"`
	WordId uint     `json:"word_id"`
	Word   UserWord `gorm:"foreignKey:WordId"`
}

type RuleOption struct {
	Id     uint   `json:"id"`
	RuleId uint   `json:"rule_id"`
	Rule   Rule   `gorm:"foreignkey:RuleId"`
	Letter string `json:"letter"`
}
