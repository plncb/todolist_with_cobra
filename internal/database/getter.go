package database

import (
	"database/sql"
	"todolist/internal/models"
)

func GetTask(db *sql.DB, id int) (*models.Task, error) {
	task := &models.Task{}
	query := "SELECT id, description, created_at, due_by, is_completed FROM " + models.TaskTableName + " WHERE id = ?"
	err := db.QueryRow(query, id).
		Scan(&task.ID, &task.Description, &task.CreatedAt, &task.DueBy, &task.IsCompleted)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func GetAllTasks(db *sql.DB) ([]models.Task, error) {
	query := "SELECT id, description, created_at, due_by, is_completed FROM " + models.TaskTableName
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.Description,
			&task.CreatedAt,
			&task.DueBy,
			&task.IsCompleted,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
