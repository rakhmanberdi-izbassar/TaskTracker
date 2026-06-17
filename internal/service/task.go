package service

import (
	"github.com/rakhmanberdi-izbassar/TaskTracker/internal/model"
	"github.com/rakhmanberdi-izbassar/TaskTracker/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(taskRepo *repository.TaskRepository) *TaskService {
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
