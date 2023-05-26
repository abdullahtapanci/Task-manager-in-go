package main

import (
	"context"
	"fmt"
	"log"
	"os"
)

var idCounter int = 0

func main() {

	var command int

	client := connectMongo()
	defer func() {
		err := client.Disconnect(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}()
	taskDatabase := client.Database("TaskDatabase")
	taskCollection := taskDatabase.Collection("TaskCollection")
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	for {
		command = showMenuAndGetCommand()
		switch command {
		case 1:
			idCounter += 1
			addTask(idCounter, taskCollection)
		case 2:
			viewTasks(taskCollection)
		case 3:
			markAsCompleted(taskCollection)
		case 4:
			deleteTask(taskCollection)
		case 5:
			filterByStatus(taskCollection)
		case 6:
			fmt.Println("Thank you for using Task Manager. See you soon:)")
			os.Exit(5)
		}
	}
}
