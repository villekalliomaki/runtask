package server

import (
	"fmt"
	"runtask/internal/config"

	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// The server holds the Fiber server and related configuration
type Server struct {
	server        *fiber.App
	logger        *zap.Logger
	port          int
	listenAddress string
}

// Create a new server, doesn't start it
func New(logger *zap.Logger, conf config.Config) Server {
	fiberServer := fiber.New()

	// The Zap for Fiber logging
	// https://github.com/gofiber/contrib/tree/main/fiberzap
	fiberServer.Use(fiberzap.New(fiberzap.Config{Logger: logger}))

	return Server{server: fiberServer, logger: logger, port: conf.Port, listenAddress: conf.Host}
}

// Start the server on the address and port defined before
// Block until server is stopped
func (server *Server) Start() {
	listen := fmt.Sprintf("%s:%v", server.listenAddress, server.port)

	server.logger.Info("Starting web server", zap.String("listen", listen))

	server.server.Listen(listen)
}
