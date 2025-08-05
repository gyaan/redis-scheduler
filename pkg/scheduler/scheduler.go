package scheduler

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/pkg/errors"
	"redis-scheduler/pkg/redisclient"
)

//createInitialList create list to process
func createInitialList(key string) error {
	//set the initial list
	client := redisclient.GetClient()

	//close connection at end
	defer client.Close()

	for i := 1; i <= 10; i++ {
		err := client.RPush(key, i).Err()
		if err != nil {
			return errors.Wrapf(err, "failed to push %d to %s", i, key)
		}
	}
	return nil
}

//task task to run on each interval
func task() error {

	//chunk size

	chunkSize := 3
	circularListName := "circular_list_for_update"
	currentListName := "current_list_for_update"

	//get redis client
	client := redisclient.GetClient()

	//close connection at end
	defer client.Close()

	//get current list
	s, err := client.LRange(circularListName, 0, -1).Result()

	if err != nil {
		return errors.Wrap(err, "failed to get current list")
	}

	fmt.Println("current list:", s)

	//get the three element from last
	//current_list_for_update
	strings, err := client.LRange(circularListName, 0, int64(chunkSize-1)).Result()
	if err != nil {
		return errors.Wrap(err, "failed to get chunk from circular list")
	}

	//remove old current list
	if err := client.Del(currentListName).Err(); err != nil {
		return errors.Wrap(err, "failed to delete old current list")
	}

	//create current_list_for_update for update
	//this list can be shared with multiple worker or job
	for i := 0; i < chunkSize; i++ {
		if err := client.RPush(currentListName, strings[i]).Err(); err != nil {
			return errors.Wrapf(err, "failed to push %s to %s", strings[i], currentListName)
		}
	}

	result, err := client.LRange(currentListName, 0, int64(chunkSize-1)).Result()
	if err != nil {
		return errors.Wrap(err, "failed to get current processing list")
	}
	fmt.Println("current processing list:", result)

	//remove three elements from front
	//todo find out a function to multiple element remove from the starting
	for i := 0; i < len(strings); i++ {
		if err := client.LPop(circularListName).Err(); err != nil {
			return errors.Wrap(err, "failed to pop from circular list")
		}
	}

	//push three elements to last
	//todo find function to push multiple element in the list
	for i := 0; i < len(strings); i++ {
		if err := client.RPush(circularListName, strings[i]).Err(); err != nil {
			return errors.Wrapf(err, "failed to push %s to %s", strings[i], circularListName)
		}
	}
	return nil
}

//Run job on certain interval
//below job is running every second
func RunJobs() error {

	//verify redis connection is available
	if err := redisclient.Ping(); err != nil {
		return errors.Wrap(err, "unable to connect to redis")
	}

	//create the list
	if err := createInitialList("circular_list_for_update"); err != nil {
		return errors.Wrap(err, "failed to create initial list")
	}

	//schedule a job
	s := gocron.NewScheduler()
	s.Every(1).Second().Do(func() {
		if err := task(); err != nil {
			fmt.Printf("task failed: %v\n", err)
		}
	})
	<-s.Start()
	return nil
}
