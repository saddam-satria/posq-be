package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/saddam-satria/posq-be/commons"
)

func GetConfigFile(file *os.File) logger.Config {
	if commons.GoDotEnv("DEBUG") == "false" {
		return logger.Config{
			Output:     file,
			TimeFormat: "2006-01-02 15:04:05",
			TimeZone:   "Asia/Jakarta",
		}
	}

	return logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}

}
