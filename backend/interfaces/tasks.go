package interfaces

type Word struct {
	Id        uint    `json:"id"`
	Word      string  `json:"word"`
	Rule      *string `json:"rule"`
	Exception bool    `json:"exception"`
	Options   []Option
}

type TaskRule struct {
	Id     uint   `json:"id"`
	Rule   string `json:"rule"`
	Errors *int64 `json:"errors"`
}

type Task struct {
	Number      uint        `json:"number" gorm:"unique"`
	Description string      `json:"description"`
	Rules       *[]TaskRule `json:"rules"`
}

type Option struct {
	Letter  string `json:"letter"`
	Correct bool   `json:"correct"`
}

var LETTEROPTIONS = map[rune][]Option{
	'О': {
		Option{Letter: "А", Correct: false},
		Option{Letter: "О", Correct: true},
	},
	'А': {
		Option{Letter: "А", Correct: true},
		Option{Letter: "О", Correct: false},
	},
	'И': {
		Option{Letter: "И", Correct: true},
		Option{Letter: "Е", Correct: false},
	},
	'Ы': {
		Option{Letter: "Ы", Correct: true},
		Option{Letter: "И", Correct: false},
	},
	'Е': {
		Option{Letter: "И", Correct: false},
		Option{Letter: "Е", Correct: true},
	},
	'Ь': {
		Option{Letter: "Ь", Correct: true},
		Option{Letter: "Ъ", Correct: false},
		Option{Letter: "-", Correct: false},
	},
	'Ъ': {
		Option{Letter: "Ь", Correct: false},
		Option{Letter: "Ъ", Correct: true},
		Option{Letter: "-", Correct: false},
	},
	'?': {
		Option{Letter: "Ь", Correct: false},
		Option{Letter: "Ъ", Correct: false},
		Option{Letter: "-", Correct: true},
	},
	'С': {
		Option{Letter: "З", Correct: false},
		Option{Letter: "С", Correct: true},
	},
	'З': {
		Option{Letter: "С", Correct: false},
		Option{Letter: "З", Correct: true},
	},
}
