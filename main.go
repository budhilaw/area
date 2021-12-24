package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"question-2/area/repository"
	"question-2/common"
	"question-2/model"
)

func init() {
	common.LoadEnv()
}

func main() {
	conn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		common.Env.DbUser, common.Env.DbPass, common.Env.DbHost,
		common.Env.DbPort, common.Env.DbName)

	dbConn, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.AutoMigrate(&model.Area{})
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		sqlDB, err := dbConn.DB()
		if err != nil {
			log.Panic(err)
		}
		sqlDB.Close()
	}()

	mysqlRepo := repository.NewMysqlAreaRepository(dbConn)
	contentType := "persegi panjang"
	err = mysqlRepo.InsertArea(20, 10, contentType)
	if err != nil {
		log.Println(err.Error())
		err = errors.New(common.ERROR_DATABASE)
	}
}
