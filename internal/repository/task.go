package repository

import (
	"github.com/rakhmanberdi-izbassar/TaskTracker/internal/model"

	"errors"
)

type TaskRepository struct {
	tasksMap map[int]model.Task
}

func NewTaskRepository() *TaskRepository {
	task := &TaskRepository{
		tasksMap: make(map[int]model.Task),
	}

	task.tasksMap[1] = model.Task{ID: 1, Title: "Task 1", Description: "Task 1 Description", Status: "asd", Priority: 1, UserID: 1}
	task.tasksMap[2] = model.Task{ID: 2, Title: "Task 2", Description: "Task 2 Description", Status: "asd", Priority: 1, UserID: 1}
	task.tasksMap[3] = model.Task{ID: 3, Title: "Task 3", Description: "Task 3 Description", Status: "asd", Priority: 1, UserID: 1}

	return task
}

func (t *TaskRepository) GetAll() ([]model.Task, error) {
	tasksSlice := make([]model.Task, 0)

	for _, task := range t.tasksMap {
		tasksSlice = append(tasksSlice, task)
	}

	return tasksSlice, nil
}

func (t *TaskRepository) GetById(id int) (model.Task, error) {
	task, ok := t.tasksMap[id]
	if !ok {
		return model.Task{}, errors.New("task not found")
	}
	return task, nil
}
