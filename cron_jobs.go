package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

//create list to process
func createInitialList(key string) {
	//set the initial list
	client := GetClient()

	//close connection at end
	defer client.Close()

	client.RPush(key, 1)
	client.RPush(key, 2)
	client.RPush(key, 3)
	client.RPush(key, 4)
	client.RPush(key, 5)
	client.RPush(key, 6)
	client.RPush(key, 7)
	client.RPush(key, 8)
	client.RPush(key, 9)
	client.RPush(key, 10)
}

//task to run on each interval
func task() {

	//get redis client
	client := GetClient()

	//close connection at end
	defer client.Close()

	//get current list
	s, err := client.LRange("items", 0, -1).Result()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("current list:", s)

	//get the three element from last
	strings, err := client.LRange("items", 0, 2).Result()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("current processing list:", strings)

	//remove three elements from front
	//todo find out a function to multiple element remove from the starting
	for i := 0; i < len(strings); i++ {
		client.LPop("items")
	}

	//push three elements to last
	//todo find function to push multiple element in the list
	for i := 0; i < len(strings); i++ {
		client.RPush("items", strings[i])
	}
}

//run job on certain interval
//below job is running every second
func runJobs() {

	//verify redis connection is available
	ping := Ping()

	if ping != nil {
		fmt.Println("unable to connect to redis")
	} else {
		//create the list
		createInitialList("items")

		//schedule a job
		s := gocron.NewScheduler()
		s.Every(1).Second().Do(task)
		<-s.Start()
	}
}
