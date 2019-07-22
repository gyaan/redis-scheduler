package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

func createInitialList() {
	//set the initial list
	client := GetClient()

	//close connection at end
	defer client.Close()

	client.RPush("items", 1)
	client.RPush("items", 2)
	client.RPush("items", 3)
	client.RPush("items", 4)
	client.RPush("items", 5)
	client.RPush("items", 6)
	client.RPush("items", 7)
	client.RPush("items", 8)
	client.RPush("items", 9)
	client.RPush("items", 10)
}

func task() {

	//get redis client
	client := GetClient()

	//close connection at end
	defer client.Close()

	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}  //this isn't work
	//client.RPush("items",arr)
	//client.Set("items", arr, 0)
	//prepare list

	//display list
	s, _ := client.LRange("items", 0, -1).Result()
	fmt.Println("current list:",s)

	//get the three element from last
	strings, _ := client.LRange("items", 0, 2).Result()
	fmt.Println("current processing list:", strings)

	//remove three elements from left side
	//find out a function to multiple element remove from the starting
	for i := 0; i < len(strings); i++ {
		client.LPop("items")
	}

	//push three elements to last
	//find function to push multiple element in the list
	for i := 0; i < len(strings); i++ {
		client.RPush("items", strings[i])
	}
}

func runJobs() {
	//create the list
	createInitialList()

	//schedule a job
	s := gocron.NewScheduler()
	s.Every(1).Second().Do(task)
	<-s.Start()
}