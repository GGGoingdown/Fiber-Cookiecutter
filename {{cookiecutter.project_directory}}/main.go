package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/GGGoingdown/Fiber-Cookiecutter/api"
	"github.com/GGGoingdown/Fiber-Cookiecutter/config"
	"github.com/GGGoingdown/Fiber-Cookiecutter/docs"
	"github.com/GGGoingdown/Fiber-Cookiecutter/utils"
)

func swaggerDocs() {
	docs.SwaggerInfo.Title = "{{ cookiecutter.project_title }}"
	docs.SwaggerInfo.Description = "{{ cookiecutter.project_description }}"
	docs.SwaggerInfo.Version = "{{ cookiecutter.project_version }}"
}

func main() {
	cfg, err := config.NewConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	logger := utils.NewLogHandler(cfg.LogPath, cfg.ToZapLogLevel())

	swaggerDocs()

	server := api.NewServer(cfg, logger)

	go func() {
		err := server.Run()
		if err != nil {
			log.Fatal("cannot start server:", err.Error())
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	log.Println("Gracefully shutting down...")
	_ = server.Shutdown()

	log.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	log.Println("server was successful shutdown.")

}
