package main

import (
	"easy-swap/api"
	"easy-swap/dal"
	"easy-swap/internal/parser"
	"easy-swap/internal/scan"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 1.初始化数据库
	dal.InitDB()

	// 2.初始化ABI解析
	err := parser.InitParser()
	if err != nil {
		log.Fatal("parser init err: ", err)
	}

	// 3.启动扫链服务
	scanner, err := scan.NewScanner()
	if err != nil {
		log.Fatal("scanner init err: ", err)
	}
	go scanner.Start()

	// 4.启动Gin接口
	r := gin.Default()
	// 路由
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/nft/transfer/list", api.GetNftTransferList)
	}

	log.Println("server run :8080")
	_ = r.Run(":8080")
}