//This project is demonstration of circular queue using redis
package main

import (
	"log"
	"redis-scheduler/config"
	"redis-scheduler/pkg/redisclient"
	"redis-scheduler/pkg/scheduler"
)

func main() {
	conf, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	redisclient.Init(conf)
	if err := scheduler.RunJobs(); err != nil {
		log.Fatalf("scheduler failed: %v", err)
	}
}