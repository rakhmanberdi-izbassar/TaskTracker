package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rakhmanberdi-izbassar/TaskTracker/internal/model"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(dbObject *sqlx.DB) *TaskRepository {
	task := &TaskRepository{
		db: dbObject,
	}
	return task
}

func (t *TaskRepository) GetAll() ([]model.Task, error) {
	tasksSlice := make([]model.Task, 0)

	query := `SELECT id, title, description, status, user_id, due_date, priority from tasks order by created_at`

	err := t.db.Select(&tasksSlice, query)

	if err != nil {
		return nil, err
	}

	return tasksSlice, nil
}

func (t *TaskRepository) GetById(id int) (model.Task, error) {

	return model.Task{}, nil
}

func (t *TaskRepository) Delete(id int) error {
	return nil
}

func (t *TaskRepository) Create(task model.Task) (int, error) {
	return task.ID, nil
}

func (t *TaskRepository) Update(id int, input model.UpdateTaskInput) (model.Task, error) {

	return model.Task{}, nil
}
