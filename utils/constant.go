package utils

import (
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/domains"
)

var DbConfig = domains.DatabaseConnection{
	HOST:     commons.GoDotEnv("DB_HOST"),
	USER:     commons.GoDotEnv("DB_USER"),
	PASSWORD: commons.GoDotEnv("DB_PASSWORD"),
	DATABASE: commons.GoDotEnv("DB_DATABASE"),
	PORT:     commons.GoDotEnv("DB_PORT"),
}

var PORT = commons.GoDotEnv("PORT")
var GO_DEBUG = commons.GoDotEnv("DEBUG")
var MIGRATE_SCRIPT = commons.GoDotEnv("MIGRATE_SCRIPT")
var SEED_SCRIPT = commons.GoDotEnv("SEED_SCRIPT")
var SECRET_KEY=commons.GoDotEnv("SECRET_KEY")