package main

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/pkg/database"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/pkg/server"
	"log"
)

func main() {

	// Setting up environment variables, reading from config
	cfg, err := config.LoadConfig("./pkg/config/config")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	// Connecting to database.
	_ = database.Connect(&cfg.DBConfig)

	server.StartServer(&cfg.ServerConfig)
}

func toDos() {
	// + project template
	// + create models (product, user, category)
	// + gin and server added.
	// add handler, creating router group, add logger
	// add jwt and middleware
	// add bulk csv
	// add basic services
	// add basket services
	// add swagger
	// add uuid
	// add advanced readme (brief explanation about project structure will be seemed complex)
	// add tests
}
