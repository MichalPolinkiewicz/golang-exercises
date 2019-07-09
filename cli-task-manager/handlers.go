package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	userInput string
)

const (
	errOrnament           = " --------- "
	errDuringScanning     = errOrnament + "Error during scanning userInput" + errOrnament
	errDuringTaskCreation = errOrnament + "Error during creating task" + errOrnament
	errDuringTaskUpdate   = errOrnament + "Error during updating task" + errOrnament
	errInvalidInput       = errOrnament + "Error - invalid input" + errOrnament
	errTaskDontExist      = errOrnament + "Error - task with given id don't exists" + errOrnament
	errNoTasksWithStatus  = errOrnament + "No tasks with given status" + errOrnament

	prefix               = " >> "
	suffix               = " <<"
	inputTaskName        = prefix + "Input task name" + suffix
	inputTaskDescription = prefix + "Input task description" + suffix
	inputTaskStatus      = prefix + "Input task status" + suffix
	taskCreated          = prefix + "Task created" + suffix
	taskUpdated          = prefix + "Task updated" + suffix
	inputTaskId          = prefix + "Input task id" + suffix
	editTaskDescription  = prefix + "Edit task description = 1" + suffix
	editTaskStatus       = prefix + "Edit task status = 2" + suffix
	chooseStatus         = prefix + "Choose status" + suffix
	toDoStatus           = prefix + "1 = to do" + suffix
	inProgressStatus     = prefix + "2 = in progress" + suffix
	doneStatus           = prefix + "3 = done" + suffix
)

func readUserInput() error {
	scanner.Scan()
	userInput = scanner.Text()
	if scanner.Err() != nil {
		log.Println(errDuringScanning, userInput)
		return scanner.Err()
	}
	return nil
}

func addNewTask(c *cli.Context) {
	var newTask task
	fmt.Println(inputTaskName)
	if readUserInput() != nil {
		return
	}
	newTask.name = userInput

	fmt.Println(inputTaskDescription)
	if readUserInput() != nil {
		return
	}
	newTask.description = userInput
	newTask.status = TO_DO

	if save(&newTask) != nil {
		fmt.Println(errDuringTaskCreation)
	} else {
		fmt.Println(taskCreated)
	}
}

func getAllTasks(c *cli.Context) {
	for _, t := range getAll() {
		fmt.Println(t.String())
	}
}

func getTasksByStatus(c *cli.Context) {
	fmt.Println(chooseStatus)
	fmt.Println(toDoStatus)
	fmt.Println(inProgressStatus)
	fmt.Println(doneStatus)
	if readUserInput() != nil {
		return
	}

	s := getConvertedTaskStatus(&userInput)
	if s == INVALID {
		return
	}
	tasks := getAllByStatus(&s)
	if len(tasks) > 0 {
		for _, t := range tasks {
			fmt.Println(t.String())
		}
	} else {
		fmt.Println(errNoTasksWithStatus)
	}
}

func updateTask(c *cli.Context) {
	fmt.Println(inputTaskId)
	if readUserInput() != nil {
		return
	}

	task := getById(&userInput)
	if task.isEmpty() {
		fmt.Println(errTaskDontExist)
		return
	}
	fmt.Println(editTaskDescription)
	fmt.Println(editTaskStatus)
	if readUserInput() != nil {
		return
	}
	if userInput != "1" && userInput != "2" {
		fmt.Println(errInvalidInput)
		return
	} else {
		switch userInput {
		case "1":
			fmt.Println(inputTaskDescription)
			if readUserInput() != nil {
				return
			}
			task.description = userInput

			if update(&task) != nil {
				fmt.Println(errDuringTaskUpdate)
			} else {
				fmt.Println(taskUpdated)
			}
			break

		case "2":
			fmt.Println(inputTaskStatus)
			if readUserInput() != nil {
				return
			}
			status := getConvertedTaskStatus(&userInput)
			if status != INVALID {
				task.status = userInput

				if update(&task) != nil {
					fmt.Println(errDuringTaskUpdate)
				} else {
					fmt.Println(taskUpdated)
				}
			}
			break
		}
	}
}

func getConvertedTaskStatus(status *string) string {
	switch *status {
	case "1":
		return TO_DO
	case "2":
		return IN_PROGRESS
	case "3":
		return DONE
	default:
		fmt.Println(errInvalidInput)
		return INVALID
	}
}
