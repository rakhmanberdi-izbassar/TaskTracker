package repository

import (
	"time"

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

func (t *TaskRepository) Delete(id int) error {
	_, ok := t.tasksMap[id]
	if !ok {
		return errors.New("task not found")
	}
	delete(t.tasksMap, id)
	return nil
}

func (t *TaskRepository) Save(task model.Task) (int, error) {
	task.ID = len(t.tasksMap) + 1

	t.tasksMap[task.ID] = task
	return task.ID, nil
}

func (t *TaskRepository) Update(id int, input model.UpdateTaskInput) (model.Task, error) {
	task, ok := t.tasksMap[id]
	if !ok {
		return model.Task{}, errors.New("task not found")
	}
	if input.Title != nil {
		task.Title = *input.Title
	}
	if input.Description != nil {
		task.Description = *input.Description
	}
	if input.Priority != nil {
		task.Priority = *input.Priority
	}
	if input.UserID != nil {
		task.UserID = *input.UserID
	}
	if input.Status != nil {
		task.Status = *input.Status
	}

	task.UpdatedAt = time.Now()

	t.tasksMap[id] = task

	return task, nil
}
