package scan

import (
	"context"
	"easy-swap/config"
	"easy-swap/dal"
	"easy-swap/internal/parser"
	"easy-swap/model"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Scanner struct {
	client     *ethclient.Client
	lastBlock  uint64
	contracts  []common.Address
}

func NewScanner() (*Scanner, error) {
	client, err := ethclient.Dial(config.RpcUrl)
	if err != nil {
		return nil, err
	}

	// 监听目标合约
	contracts := []common.Address{
		config.NftContract,
		config.MarketContract,
	}

	// 初始区块
	latest, _ := client.BlockNumber(context.Background())
	return &Scanner{
		client:    client,
		lastBlock: latest,
		contracts: contracts,
	}, nil
}

// 轮询扫描
func (s *Scanner) Start() {
	log.Println("scan service start...")
	for {
		s.scanLatest()
	}
}

func (s *Scanner) scanLatest() {
	ctx := context.Background()
	latestBlock, err := s.client.BlockNumber(ctx)
	if err != nil {
		return
	}

	if s.lastBlock >= latestBlock {
		return
	}

	// 构造过滤器
	filter := ethereum.FilterQuery{
		Addresses: s.contracts,
		FromBlock: big.NewInt(int64(s.lastBlock + 1)),
		ToBlock:   big.NewInt(int64(latestBlock)),
	}

	logs, err := s.client.FilterLogs(ctx, filter)
	if err != nil {
		return
	}

	// 循环解析日志
	for _, vLog := range logs {
		s.handleLog(&vLog)
	}

	// 更新已扫描区块
	s.lastBlock = latestBlock
}

// 分发不同事件
func (s *Scanner) handleLog(log *types.Log) {
	eventSig := log.Topics[0].Hex()

	switch eventSig {
	// Transfer 事件签名
	case "0xddf252ad1be2c89b69c2b068fc378daa952fdfc71349660f863f565bd9dd65ea":
		transfer, err := parser.ParseTransferLog(log)
		if err != nil {
			return
		}
		// 入库
		dal.DB.Create(&model.NftTransfer{
			TokenID:     transfer.TokenId,
			FromAddress: transfer.From.Hex(),
			ToAddress:   transfer.To.Hex(),
			TxHash:      log.TxHash.Hex(),
			BlockNumber: log.BlockNumber,
		})
	}
}