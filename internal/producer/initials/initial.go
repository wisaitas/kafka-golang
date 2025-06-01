package initials

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/wisaitas/kafka-golang/internal/producer"
	"github.com/wisaitas/share-pkg/config"
)

func init() {
	if err := config.ReadConfig(producer.Config); err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

}

func InitializeApp() {
	server := fiber.New()

	server.Use(logger.New())

	util := newUtil()
	service := newService(util)
	handler := newHandler(service)

	newRoute(server, handler)

	run(server)
}

func run(router *fiber.App) {
	go func() {
		if err := router.Listen(fmt.Sprintf(":%d", producer.Config.Server.Port)); err != nil {
			log.Fatalf("error starting server: %v\n", err)
		}
	}()

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-gracefulShutdown

	close()
}

func close() {

	log.Println("gracefully shutdown")
}
