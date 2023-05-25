package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getString(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.TrimSpace(data)
	return data
}

var menu string = `
*************** WELCOME TO TASK MANAGER *****************
*** 1 -> Add a task                                   ***   
*** 2 -> View tasks                                   ***   
*** 3 -> Mark a task as completed                     ***   
*** 4 -> Delete a task                                ***  
*** 5 -> Filter tasks by status                       ***
*** 6 -> Exit                                         ***   
*********************************************************
`

func showMenuAndGetCommand() int {
	fmt.Println(menu)
	var command int
	fmt.Print("Enter your choice : ")
	fmt.Scanln(&command)
	return command
}
