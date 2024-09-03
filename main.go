package main

import (
	"todolist/cmd"
	"todolist/internal/database"
)

func main() {
	database.Init("./internal/database/task.sqlite3")
	cmd.Execute()
}
