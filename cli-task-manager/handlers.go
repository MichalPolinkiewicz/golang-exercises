package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"strconv"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	userInput string
)

func addNewTask(c *cli.Context) {
	var newTask task
	fmt.Println("Please input task name")
	scanner.Scan()
	newTask.name = scanner.Text()
	if scanner.Err() != nil {
		log.Println("error during creating userInput", userInput)
	}
	fmt.Println("Please input task description")
	scanner.Scan()
	newTask.description = scanner.Text()
	newTask.status = TO_DO
	if scanner.Err() != nil {
		log.Println("error during creating userInput", userInput)
	}

	fmt.Printf("Task '%s' created \n", newTask.name)
	if err := save(&newTask); err != nil {
		log.Println(err)
	}
}

func getAllTasks(c *cli.Context) {
	for _, t := range getAll() {
		fmt.Printf("TASK_ID: %s, DESCRIPTION: %s, STATUS: %s \n", t.name, t.description, t.status)
	}
}

func getTasksByStatus(c *cli.Context) {
	fmt.Printf("Choose valid status: \n1 = %s \n2 = %s \n3 = %s \n", TO_DO, IN_PROGRESS, DONE)
	scanner.Scan()
	userInput = scanner.Text()

	isValid := func(status *string) bool {
		statusAsInt, err := strconv.Atoi(*status)
		if err != nil {
			log.Print("error during converting status to int")
		}
		if statusAsInt >= 0 && statusAsInt <= 3 {
			return true
		}
		return false
	}

	if isValid(&userInput) {

	} else {
		fmt.Println("Invalid status")
		getTasksByStatus(c)
	}
}

func editTask(c *cli.Context) {
	fmt.Println("Put task name")
	scanner.Scan()
	userInput = scanner.Text()

	if taskExists() {
		isValid := func() bool {
			fmt.Printf("Edit task description = 1\nEdit task status = 2\n")
			scanner.Scan()
			userInput = scanner.Text()

			if userInput != "1" && userInput != "2" {
				fmt.Println("Wrong input")
				return false
			}

			fmt.Println("put new text")
			scanner.Scan()
			userInput = scanner.Text()

			return true
		}

		for {
			if isValid() == false {
				isValid()
			} else {
				break
			}
		}

		fmt.Print("end")
	}
}

func taskExists() bool {
	return true
}
