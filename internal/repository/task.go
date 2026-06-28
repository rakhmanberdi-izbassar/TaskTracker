package repository

import (
	"errors"
	"fmt"
	"strings"
	"time"

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

	var tasks model.Task

	query := `SELECT * FROM tasks WHERE id = $1 LIMIT 1`

	err := t.db.Get(&tasks, query, id)
	if err != nil {
		return model.Task{}, err
	}

	return tasks, nil
}

func (t *TaskRepository) Delete(id int) error {

	query := `DELETE FROM tasks WHERE id = $1`

	_, err := t.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (t *TaskRepository) Create(task model.Task) (int, error) {

	query := `INSERT INTO tasks (title, description, status, user_id, priority, created_at, updated_at) VALUES (:title, :description, :status, :user_id, :priority, :created_at, :updated_at) RETURNING id`

	rows, err := t.db.NamedQuery(query, task)

	if err != nil {
		return 0, err
	}
	var id int
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

func (t *TaskRepository) Update(id int, input model.UpdateTaskInput) (int, error) {
	var setParts []string //["title=$1", "description=$2"]
	var args []any        //["Test","qweqwe"]

	argId := 1

	if input.Title != nil {
		setParts = append(setParts, fmt.Sprintf("title=$%d", argId))
		args = append(args, input.Title)
		argId++
	}
	if input.Description != nil {
		setParts = append(setParts, fmt.Sprintf("description=$%d", argId))
		args = append(args, input.Description)
		argId++
	}
	if input.Status != nil {
		setParts = append(setParts, fmt.Sprintf("status=$%d", argId))
		args = append(args, input.Status)
		argId++
	}
	if input.UserID != nil {
		setParts = append(setParts, fmt.Sprintf("user_id=$%d", argId))
		args = append(args, input.UserID)
		argId++
	}
	if input.Priority != nil {
		setParts = append(setParts, fmt.Sprintf("priority=$%d", argId))
		args = append(args, input.Priority)
		argId++
	}

	if len(setParts) == 0 {
		return 0, errors.New("no field to update")
	}

	setParts = append(setParts, fmt.Sprintf("updated_at=$%d", argId))
	args = append(args, time.Now())
	argId++

	query := fmt.Sprintf(`
		UPDATE tasks 
		SET %s 
		WHERE id = $%d
		`, strings.Join(setParts, ","), argId)
	args = append(args, id)

	_, err := t.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return id, nil
}
