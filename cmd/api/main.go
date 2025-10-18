package main

import (
	"fmt"

	"github.com/loissascha/go-http-server/server"
	"github.com/loissascha/go-logger/logger"
	"github.com/loissascha/go-templ-template/internal/handlers/homehandler"
)

func main() {
	port := "42069"
	s := server.NewServer()

	// handler
	homeH := homehandler.New()

	// register routes
	homeH.RegisterHandlers(s)

	logger.Info(nil, "Server starting at port: {port}", port)
	err := s.Serve(fmt.Sprintf(":%v", port))
	if err != nil {
		logger.Error(err, "Server failed to start...")
	}
}
