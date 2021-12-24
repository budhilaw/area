package common

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var Env env

type env struct {
	DbUser  string
	DbPass  string
	DbName  string
	DbHost  string
	DbPort  int
	AppPort int
}

func LoadEnv() {
	godotenv.Load()

	Env.DbUser = os.Getenv("DB_USER")
	Env.DbPass = os.Getenv("DB_PASS")
	Env.DbName = os.Getenv("DB_NAME")
	Env.DbHost = os.Getenv("DB_HOST")
	Env.DbPort, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	Env.AppPort, _ = strconv.Atoi(os.Getenv("APP_PORT"))
}
