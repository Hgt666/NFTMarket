package dal

import (
	"easy-swap/config"
	"easy-swap/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(mysql.Open(config.MysqlDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("mysql connect failed: ", err)
	}
	// 自动建表
	err = db.AutoMigrate(&model.NftTransfer{}, &model.MarketOrder{})
	if err != nil {
		log.Fatal("migrate failed: ", err)
	}
	DB = db
}