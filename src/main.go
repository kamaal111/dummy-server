package main

import (
	"fmt"
	"os"

	"github.com/kamaal111/dummy-server/src/router"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}

	router.HandleRequests(fmt.Sprintf(":%s", PORT))
}
