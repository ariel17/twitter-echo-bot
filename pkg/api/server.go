package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ariel17/twitter-echo-bot/pkg/configs"
)

const statusPath = "/status"

// StartServer creates a new instance of HTTP server with indicated handlers
// configured and begins serving content.
func StartServer() {
	http.HandleFunc(statusPath, statusHandler)
	s := http.Server{Addr: fmt.Sprintf(":%d", configs.APIPort)}

	log.Printf("Listening new connections on :%d", configs.APIPort)
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}