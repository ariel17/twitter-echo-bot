package api

import (
	"fmt"
	"net/http"

	"github.com/ariel17/twitter-echo-bot/pkg/configs"
)

const statusPath = "/status"

// NewServer creates a new instance of HTTP server with indicated handlers
// configured.
func NewServer() *http.Server {
	http.HandleFunc(statusPath, statusHandler)
	return &http.Server{Addr: fmt.Sprintf(":%d", configs.APIPort)}
}