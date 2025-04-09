package main

import (
	"log"
	"ysd_pos_server/initialization/database"
	"ysd_pos_server/initialization/env"
)

func main() {
	initialization()
}

func initialization() {
	var err error
	if err = env.InitializeEnv(); err != nil {
		log.Fatalf("Error loading .env file: %v", err.Error())
	}

	if err = database.InitializeDatabase(); err != nil {
		log.Fatalf("Error connecting to database: %v", err.Error())
	}
}
