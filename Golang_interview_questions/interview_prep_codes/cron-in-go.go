package mycronjob

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

// Task is the function to run as a cron job
func Task() {
	fmt.Println("Running scheduled task:", time.Now())
}

// StartCron starts the cron scheduler
func StartCron() {
	c := cron.New(cron.WithSeconds())
	// Schedule the task to run every minute
	_, err := c.AddFunc("*/1 * * * *", Task)
	if err != nil {
		fmt.Println("Error scheduling task:", err)
		return
	}
	c.Start()
	fmt.Println("Cron job started.")
	select {} // Block forever
}

//CALLING THIS CRON PACKAGE

// package main

// import (
// 	"mycronjob"
// )

// func main() {
// 	mycronjob.StartCron()
// }
