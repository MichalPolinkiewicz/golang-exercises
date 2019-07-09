package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

const (
	TO_DO       = "to do"
	IN_PROGRESS = "in progress"
	DONE        = "done"
	INVALID     = "invalid"
)

type task struct {
	id          int
	name        string
	description string
	status      string
}

func (t *task) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, Description = %s, status = %s", t.id, t.name, t.description, t.status)
}

func (t *task) isEmpty() bool {
	if t.id != 0 {
		return false
	}
	return true
}

func createApp() *cli.App {
	app := cli.NewApp()
	info(app)
	commands(app)

	return app
}

func info(app *cli.App) {
	app.Name = "CLI-Task-Manager"
	app.Author = "MichalPolinkiewicz"
	app.Description = "simple task manager produced for exercises"
}

func commands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "Get all tasks",
			Action:  getAllTasks,
		},
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "Creates a new task with default status to do",
			Action:  addNewTask,
		},
		{
			Name:    "byStatus",
			Aliases: []string{"s"},
			Usage:   "Get all tasks with given status",
			Action:  getTasksByStatus,
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "Update task description or status",
			Action:  updateTask,
		},
	}
}

func main() {
	app := createApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
