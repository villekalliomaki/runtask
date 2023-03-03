package main

import (
	"runtask/internal/config"
	"runtask/internal/db"
	"runtask/internal/js"
	"runtask/internal/logs"
	"runtask/internal/server"
)

func main() {
	log := logs.NewLogger()
	cfg := config.New(log)
	_ = db.New(log, cfg)
	_ = js.New(log)
	server := server.New(log, cfg)

	server.Start()
}
