package database

import (
	"database/sql"
	"fmt"
	"time"
	"todolist/internal/models"
)

func CreateTask(db *sql.DB, description string, dueBy time.Time) (sql.Result, error) {
	query := "INSERT INTO " + models.TaskTableName + " (description, due_by) VALUES (?, ?)"
	sqlResult, err := db.Exec(query, description, dueBy)

	return sqlResult, err
}

func DeleteTask(db *sql.DB, id int) error {
	query := "DELETE FROM " + models.TaskTableName + " WHERE id = ?"
	sqlResult, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no task found with id %d", id)
	}

	return nil
}

func CompleteTask(db *sql.DB, id int) error {
	query := "UPDATE " + models.TaskTableName + " SET is_completed = 1 WHERE id = ?"
	sqlResult, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no task found with id %d", id)
	}

	return nil
}
