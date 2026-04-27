package api

import (
	"easy-swap/dal"
	"easy-swap/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取NFT流转记录
func GetNftTransferList(c *gin.Context) {
	var list []model.NftTransfer
	address := c.Query("address")

	tx := dal.DB
	if address != "" {
		tx = tx.Where("from_address = ? OR to_address = ?", address, address)
	}

	err := tx.Order("id desc").Limit(20).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "query fail"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": list,
	})
}