package words

import (
	"rusEGE/interfaces"
	"rusEGE/repositories"
)

func GetWordErrors(
	tr *repositories.GormTaskRepository,
	wr *repositories.GormWordRepository,
) ([]*interfaces.StatTask, error){
	tasks, err := tr.All()
	if err != nil{
		return nil, err
	}

	tasksStat := make([]*interfaces.StatTask, len(tasks))

	for i, task := range(tasks){
		words, err := wr.GetTaskWordsWithError(task.Id)

		if err != nil{
			return nil, err
		}

		tasksStat[i] = &interfaces.StatTask{
			Task: task.Number,
			Words: words,
		}
	}

	return tasksStat, nil
}