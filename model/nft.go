package model

import "gorm.io/gorm"

type NftTransfer struct {
	gorm.Model
	TokenID     string `gorm:"index"`
	FromAddress string
	ToAddress   string
	TxHash      string `gorm:"uniqueIndex"` // 全局去重
	BlockNumber uint64
}