package models

import (
	"time"
)

const TaskTableName = "tasks"

type Task struct {
	ID          int
	Description string
	CreatedAt   time.Time
	DueBy       time.Time
	IsCompleted bool
}

func (t *Task) TableName() string {
	return TaskTableName
}
