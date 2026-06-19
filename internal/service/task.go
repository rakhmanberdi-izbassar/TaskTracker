package service

import (
	"errors"
	"time"

	"github.com/rakhmanberdi-izbassar/TaskTracker/internal/model"
	"github.com/rakhmanberdi-izbassar/TaskTracker/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepoInterface
}

func NewTaskService(taskRepo repository.TaskRepoInterface) *TaskService {
	service := &TaskService{
		repo: taskRepo,
	}

	return service
}

func (ts *TaskService) GetAll() ([]model.Task, error) {
	tasks, err := ts.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (ts *TaskService) GetById(id int) (model.Task, error) {
	return ts.repo.GetById(id)
}

func (ts *TaskService) Delete(id int) error {
	return ts.repo.Delete(id)
}

func (ts *TaskService) Create(input model.CreateTaskInput) (int, error) {
	if input.Title == "" {
		return 0, errors.New("title is required")
	}
	if input.Description == "" {
		return 0, errors.New("description is required")
	}
	task := model.Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		UserID:      input.UserID,
		Priority:    input.Priority,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return ts.repo.Create(task)
}

func (ts *TaskService) Update(id int, input model.UpdateTaskInput) (model.Task, error) {
	_, err := ts.repo.GetById(id)
	if err != nil {
		return model.Task{}, errors.New("task not found")
	}

	if input.Title != nil && *input.Title == "" {
		return model.Task{}, errors.New("title is required")
	}
	if input.Description != nil && *input.Description == "" {
		return model.Task{}, errors.New("description is required")
	}

	return ts.repo.Update(id, input)

}
