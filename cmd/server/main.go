package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/loissascha/go-http-server/server"
	"github.com/loissascha/go-logger/logger"
	"github.com/loissascha/go-templ-template/internal/handlers/homehandler"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT not defined. Make sure there is a .env file or a environment variable set!")
	}

	s, err := server.NewServer(
		server.EnableTranslations(),
		server.EnableAutoDetectLanguage(),
		server.SetDefaultLanguage("en"),
		server.AddTranslationFile("en", "src/en.json"),
		server.AddTranslationFile("de", "src/de.json"),
		server.AddTranslationFile("es", "src/es.json"),
		server.AddTranslationFile("fr", "src/fr.json"),
	)
	if err != nil {
		panic(err)
	}

	// handler
	homeH := homehandler.New(s)

	// register routes
	homeH.RegisterHandlers(s)

	fs := http.FileServer(http.Dir("./static"))
	s.GetMux().Handle("/static/", http.StripPrefix("/static/", fs))

	logger.Info(nil, "Server starting at port: {port}", port)
	err = s.Serve(fmt.Sprintf(":%v", port))
	if err != nil {
		logger.Error(err, "Server failed to start...")
	}
}
