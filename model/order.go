package model

import "gorm.io/gorm"

type MarketOrder struct {
	gorm.Model
	OrderId     string `gorm:"uniqueIndex"`
	TokenID     string
	Seller      string
	Price       string
	Status      uint8 // 0:挂单 1:已成交 2:取消
	TxHash      string
	BlockNumber uint64
}