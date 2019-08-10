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

	//chunk size

	chunkSize := 3
	circularListName := "circular_list_for_update"
	currentListName := "current_list_for_update"

	//get redis client
	client := GetClient()

	//close connection at end
	defer client.Close()

	//get current list
	s, err := client.LRange(circularListName, 0, -1).Result()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("current list:", s)

	//get the three element from last
	//current_list_for_update
	strings, err := client.LRange(circularListName, 0, int64(chunkSize-1)).Result()
	if err != nil {
		fmt.Println(err)
	}

	//remove old current list
	client.Del(currentListName)

	//create current_list_for_update for update
	//this list can be shared with multiple worker or job
	for i := 0; i < chunkSize; i++ {
		client.RPush(currentListName, strings[i])
	}

	result, err := client.LRange(currentListName,0,int64(chunkSize-1)).Result()
	fmt.Println("current processing list:", result)

	//remove three elements from front
	//todo find out a function to multiple element remove from the starting
	for i := 0; i < len(strings); i++ {
		client.LPop(circularListName)
	}

	//push three elements to last
	//todo find function to push multiple element in the list
	for i := 0; i < len(strings); i++ {
		client.RPush(circularListName, strings[i])
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
		createInitialList("circular_list_for_update")

		//schedule a job
		s := gocron.NewScheduler()
		s.Every(1).Second().Do(task)
		<-s.Start()
	}
}
