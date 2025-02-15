package config

import (
	"github.com/omniful/go_commons/db/sql/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var Config = postgres.DBConfig{
	Host:     "localhost",
	Port:     "5432",
	Username: "postgres",
	Password: "namita123",
	Dbname:   "wms_db",
}
