package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
		ErrorHandler:  commons.ErrorHandler,
		StrictRouting: true,
	})
	file, err := os.OpenFile("./logs/access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		os.Create("./logs/access.log")
	}
	defer file.Close()

	server.Use(cors.New())
	server.Use(logger.New(middlewares.GetConfigFile(file)))
	server.Static("/assets", "./public")
	server.Use(recover.New())
	router := server.Group("/api/v1", middlewares.BaseMiddleware)
	utils.GetRoute(router)

	runtimeLogMessage := fmt.Sprintf("%s server running on port %s \n", time.Now().Format("2006-01-02 15:04:05"), utils.PORT)

	if utils.GO_DEBUG == "false" {
		file.WriteString(runtimeLogMessage)
	} else {
		fmt.Println(runtimeLogMessage)
	}

	if err := server.Listen(":" + utils.PORT); err != nil {
		message := "Failed to run server " + err.Error() + "\n"
		if utils.GO_DEBUG == "false" {
			file.WriteString(message)
		} else {
			fmt.Println(message)
		}
		panic("Failed to run server" + err.Error())
	}
}
