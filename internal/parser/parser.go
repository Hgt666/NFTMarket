package parser

import (
	"strings"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// 内置ABI，也可以读本地文件
const NftABI = `[{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"tokenId","type":"uint256"}],"name":"Transfer","type":"event"}]`

var nftAbi abi.ABI

func InitParser() error {
	parsedAbi, err := abi.JSON(strings.NewReader(NftABI))
	if err != nil {
		return err
	}
	nftAbi = parsedAbi
	return nil
}

// 解析NFT Transfer事件
type TransferEvent struct {
	From    common.Address
	To      common.Address
	TokenId string
}

func ParseTransferLog(log *types.Log) (*TransferEvent, error) {
	var event TransferEvent
	// 解析data
	err := nftAbi.UnpackIntoInterface(&event, "Transfer", log.Data)
	if err != nil {
		return nil, err
	}
	// 解析topic  indexed参数
	event.From = common.BytesToAddress(log.Topics[1].Bytes())
	event.To = common.BytesToAddress(log.Topics[2].Bytes())

	return &event, nil
}