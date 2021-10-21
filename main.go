package main

import (
	"log"

	"github.com/ariel17/twitter-echo-bot/pkg/clients"
)

func main() {
	// api.StartServer()
	twitter := clients.NewTwitterClient()
	tweet, response, err := twitter.Statuses.Update("hola mundo :]", nil)
	if err != nil {
		panic(err)
	}
	log.Printf("%+v\n", tweet)
	log.Printf("%+v\n", response)
}