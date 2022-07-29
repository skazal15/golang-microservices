package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnDatabase() *gorm.DB {
	cfg, _ := ConfigYaml()
	dsn := cfg.User + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("can't connect to database")
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(cfg.MaxIdleUser)
	sqlDB.SetMaxOpenConns(cfg.MaxUserConnection)
	return db
}
