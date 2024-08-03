package main

import (
	"github.com/go-sql-driver/mysql"
	"github.com/sebomancien/api/logger"
	"github.com/sebomancien/api/storage"
	store "github.com/sebomancien/api/storage/mysql"
)

func main() {
	store := NewStore()

	server := NewAPIServer(":8080", &store)

	err := server.Start()
	if err != nil {
		logger.LogError(err)
	}
}

func NewStore() storage.Store {
	config := mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	store := store.NewStore(config)

	store.Init()

	return store
}
