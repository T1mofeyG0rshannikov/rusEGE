package schemas

type GetWordsRequest struct {
	Task    uint    `query:"task"`
	RuleIds *[]uint `query:"rule_ids"`
}

type CreateTaskRequest struct {
	Number      uint   `json:"number"`
	Description string `json:"description"`
}

type EditTaskRequest struct {
	Description string `json:"description"`
}

type CreateWordRequest struct {
	TaskNumber uint   `json:"task"`
	Word       string `json:"word"`
	Rule       string `json:"rule"`
	Exception  bool   `json:"exception"`
}

type BulkCreateWordRequest struct {
	TaskNumber uint   `json:"task"`
	Content    string `json:"content"`
}

type EditWordRequest struct {
	Id        uint    `json:"id"`
	Word      *string `json:"word"`
	Rule      *string `json:"rule"`
	Exception *bool   `json:"exception"`
}

type DeleteWordsRequest struct {
	Word string `query:"word"`
}

type CreateWordErrorRequest struct {
	Word string `json:"word"`
}

type DeleteUserErrorRequest struct {
	Word uint `json:"word_id"`
}

type EditRuleRequest struct {
	Id      uint      `json:"id"`
	NewRule *string   `json:"rule"`
	Options *[]string `json:"options"`
}
