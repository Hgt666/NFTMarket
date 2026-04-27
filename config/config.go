package config

import "github.com/ethereum/go-ethereum/common"

// 链配置
var (
	RpcUrl      = "https://bsc-testnet-rpc.bnbchain.org"
	NftContract = common.HexToAddress("0x你的NFT合约")
	MarketContract = common.HexToAddress("0x你的交易市场合约")
)

// 数据库配置
const (
	MysqlDSN = "root:密码@tcp(127.0.0.1:3306)/easyswap?charset=utf8mb4&parseTime=True&loc=Local"
)

// 扫链配置
const (
	ScanInterval = 3 // 秒级轮询
	StartBlock   = 0 // 自定义起始区块
)