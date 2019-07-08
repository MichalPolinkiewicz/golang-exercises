package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

const (
	TO_DO       = "to do"
	IN_PROGRESS = "in progress"
	DONE        = "done"
)

var (
	app   *cli.App
	//tasks = []task{
	//	{name: "name 1", description: "task number 1", status: TO_DO},
	//	{name: "name 2", description: "second task", status: DONE},
	//}
)

type task struct {
	name        string
	description string
	status      string
}

func info() {
	app.Name = "CLI-Task-Manager"
	app.Author = "MichalPolinkiewicz"
	app.Description = "simple task manager produced for exercises"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "all",
			Aliases: []string{"at"},
			Usage:   "Get all tasks",
			Action:  getAllTasks,
		},
		{
			Name:    "new",
			Aliases: []string{"nt"},
			Usage:   "Creates a new task with status to do",
			Action:  addNewTask,
		},
		{
			Name:    "byStatus",
			Aliases: []string{"bs"},
			Usage:   "Get all tasks with given status",
			Action:  getTasksByStatus,
		},
		{
			Name:    "editTask",
			Aliases: []string{"et"},
			Usage:   "Edit task",
			Action: editTask,
		},
	}
}

func main() {
	app = cli.NewApp()
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
