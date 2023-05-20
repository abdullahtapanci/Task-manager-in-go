package main

import (
	"fmt"
	"os"
)

// add, view, mark, delete, exit
// id, title, description, status
func main() {

	var command int
	idCounter := 0

	for command != 5 {
		command = showMenuAndGetCommand(command)

		switch command {
		case 1:
			idCounter += 1
			var id = idCounter
			var title string
			var description string
			status := "Pending"
			title = getString("Enter the title of the task :")
			description = getString("Enter the description of the task :")
			addTask(id, title, description, status)
		case 2:
			viewTask()
		case 3:
			var id int
			fmt.Print("Enter the ID of the task to be marked as completed:")
			fmt.Scanln(&id)
			markCompleted(id)
		case 4:
			var id int
			fmt.Print("Enter the ID of the task to be deleted:")
			fmt.Scanln(&id)
			deleteTask(id)
		case 5:
			os.Exit(5)
		}
	}

}
