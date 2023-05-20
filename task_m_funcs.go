package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var tasks = make(map[int]map[string]string)

var menu string = `
*************** WELCOME TO TASK MANAGER *****************
*** 1 -> Add a task                                   ***   
*** 2 -> View tasks                                   ***   
*** 3 -> Mark a task as completed                     ***   
*** 4 -> Delete a task                                ***   
*** 5 -> Exit                                         ***   
*********************************************************
`

func getString(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.TrimSpace(data)
	return data
}

func showMenuAndGetCommand(command int) int {
	fmt.Println(menu)
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&command)
	return command
}

func addTask(id int, title string, description string, status string) {
	tasks[id] = map[string]string{
		"title":       title,
		"description": description,
		"status":      status,
	}
	fmt.Println("Task hes been added successfully!")
}

func viewTask() {
	if len(tasks) > 0 {
		fmt.Println("Tasks:")
		fmt.Println("")
		fmt.Println("**********************************************")
		for key, value := range tasks {
			fmt.Println("Task ID: ", key)
			fmt.Println("Title: ", value["title"])
			fmt.Println("Description: ", value["description"])
			fmt.Println("Status: ", value["status"])
			fmt.Println("**********************************************")
		}
	} else {
		fmt.Println("No tasks found!")
	}
	var answer string
	fmt.Println("Do you want to make any change on tasks?(y/n)")
	fmt.Scanln(&answer)
	if answer == "y" {
		var options string = `
************* OPTİONS ********************
*** 1-> Change title of a task
*** 2-> Change dexcription of a task
******************************************
		`
		fmt.Println(options)
		fmt.Print("Choose one : ")
		var selectedOption int
		fmt.Scanln(&selectedOption)
		if selectedOption == 1 {
			var selectedId int
			var newTitle string
			fmt.Print("Enter the id of task that you want to make a change : ")
			fmt.Scanln(&selectedId)
			newTitle = getString("Enter a new title :")
			tasks[selectedId]["title"] = newTitle
			fmt.Println("Change has been done successfully!")
		} else if selectedOption == 2 {
			var selectedId int
			var newDescription string
			fmt.Print("Enter the id of task that you want to make a change : ")
			fmt.Scanln(&selectedId)
			newDescription = getString("Enter a new description : ")
			tasks[selectedId]["description"] = newDescription
			fmt.Println("Change has been done successfully!")
		}

	} else if answer == "n" {
		fmt.Println("Okay, let's continue.")
	}

}

func markCompleted(id int) {
	_, isPresent := tasks[id]
	if isPresent {
		tasks[id]["status"] = "completed"
		fmt.Println("Task has been marked successfully!")
	} else {
		fmt.Println("No task found ")
	}
}

func deleteTask(id int) {
	_, isPresent := tasks[id]
	if isPresent {
		delete(tasks, id)
		fmt.Println("Task hea been deleted successfully!")
	} else {
		fmt.Println("No task found ")
	}
}
