package repository

import "github.com/rakhmanberdi-izbassar/TaskTracker/internal/model"

type TaskRepoInterface interface {
	GetAll() ([]model.Task, error)
	GetById(id int) (model.Task, error)
	Delete(id int) error
	Create(task model.Task) (int, error)
	Update(id int, input model.UpdateTaskInput) (model.Task, error)
}
