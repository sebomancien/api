package main

import "github.com/sebomancien/api/logger"

func main() {
	server := NewAPIServer(":8080")

	err := server.Start()
	if err != nil {
		logger.LogError(err)
	}
}
