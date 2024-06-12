package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/middlewares"
	"github.com/saddam-satria/posq-be/utils"
)

func main() {
	if err := commons.DatabaseConnect(utils.DbConfig); err != nil {
		panic("Database Failed To Connect" + err.Error())
	}

	args := os.Args

	if len(args) > 1 && args[1] == utils.MIGRATE_SCRIPT {
		migrate(commons.DatabaseConnection)
		return
	}

	if len(args) > 1 && args[1] == utils.SEED_SCRIPT {
		Seed(commons.DatabaseConnection)
		return
	}

	commons.DatabaseConnection.Debug()
	server := fiber.New(fiber.Config{
		ErrorHandler: commons.ErrorHandler,
	})

	server.Use(recover.New())
	server.Use(middlewares.BaseMiddleware)
	utils.GetRoute(server)

	if err := server.Listen(":" + utils.PORT); err != nil {
		panic("Failed to run server" + err.Error())
	}

	fmt.Println("server running on port" + utils.PORT)
}
