package schemas


type CreateTaskRequest struct {
	Number uint `json:"number"`
	Description string `json:"description"`
}

type EditTaskRequest struct {
	Description string `json:"description"`
}

type CreateWordRequest struct {
	TaskNumber uint `json:"task_number"`
	Word string `json:"word"`
	Rule *string `json:"rule"`
}
