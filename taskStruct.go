package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type task struct {
	Id          int    `bson:"Id"`
	Title       string `bson:"Title"`
	Description string `bson:"Description"`
	Status      string `bson:"Status"`
}

func addTask(task_id int, taskCollection *mongo.Collection) {
	title := getString("Enter a title four the task : ")
	description := getString("Enter a description for your task : ")
	newTask := task{
		Id:          task_id,
		Title:       title,
		Description: description,
		Status:      "Pending",
	}
	_, err := taskCollection.InsertOne(context.Background(), newTask)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("New task has been added successfully!")
	}
}

func viewTasks(taskCollection *mongo.Collection) {
	filter := bson.M{}
	cursor, err := taskCollection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var values task
		err = cursor.Decode(&values)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("*******************************************")
		fmt.Println("Task id : ", values.Id)
		fmt.Println("Task title : ", values.Title)
		fmt.Println("Task description : ", values.Description)
		fmt.Println("Task status : ", values.Status)

	}
}

func markAsCompleted(taskCollection *mongo.Collection) {
	fmt.Print("Enter the ID of the task that you want to mark as completed : ")
	var task_id int
	fmt.Scanln(&task_id)
	filter := bson.M{"Id": task_id}
	update := bson.M{
		"$set": bson.M{
			"Status": "Completed",
		},
	}

	_, err := taskCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Task has been marked as completed successfully!")
}

func deleteTask(taskCollection *mongo.Collection) {
	fmt.Print("Enter the ID of the task that you want to delete : ")
	var task_id int
	fmt.Scanln(&task_id)
	filter := bson.M{}

	_, err := taskCollection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Task hes been deleted successfully!")

}

func filterByStatus(taskCollection *mongo.Collection) {
	fmt.Println(`
1 -> Filter by status pending	
2 -> Filter by status completed
	`)
	var choice int
	fmt.Print("Enter your choice : ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		filter := bson.M{"Status": "Pending"}
		cursor, err := taskCollection.Find(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close(context.Background())
		for cursor.Next(context.Background()) {
			var values task
			err = cursor.Decode(&values)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("*******************************************")
			fmt.Println("Task id : ", values.Id)
			fmt.Println("Task title : ", values.Title)
			fmt.Println("Task description : ", values.Description)
			fmt.Println("Task status : ", values.Status)
		}
	case 2:
		filter := bson.M{"Status": "Completed"}
		cursor, err := taskCollection.Find(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close(context.Background())
		for cursor.Next(context.Background()) {
			var values task
			err = cursor.Decode(&values)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("*******************************************")
			fmt.Println("Task id : ", values.Id)
			fmt.Println("Task title : ", values.Title)
			fmt.Println("Task description : ", values.Description)
			fmt.Println("Task status : ", values.Status)
		}
	}
}
