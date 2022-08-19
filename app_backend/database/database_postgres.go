package database

import (
	"github.com/magic-of-gnu/echo-foodchain/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection struct {
}

type PGConnectionConfig struct {
	host     string
	port     string
	db_name  string
	username string
	password string
	sslmode  string
}

func getPGConfigFromEnv() *PGConnectionConfig {

	cfg := PGConnectionConfig{
		host:     utils.MustGet("host"),
		port:     utils.MustGet("port"),
		db_name:  utils.MustGet("db_name"),
		username: utils.MustGet("username"),
		password: utils.MustGet("password"),
		sslmode:  utils.Get("sslmode", "disable"),
	}

	return &cfg

}

func CreatePGDBConnection(cfg PGConnectionConfig) {

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db

}
