package main

import (
	"github.com/ariel17/twitter-echo-bot/pkg/api"
	"github.com/ariel17/twitter-echo-bot/pkg/jobs"
)

func main() {
	jobs.NewScheduler().Start()
	api.StartServer()
}
