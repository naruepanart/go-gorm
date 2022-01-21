package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ConDB *gorm.DB

func ConnectDatabaseSQL() {
	dsn := "host=localhost user=postgres password=example dbname=edutech2 port=5432 TimeZone=Asia/Shanghai"

	var err error
	if ConDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		panic(err)
	}
}
