package jsonrpc

import (
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strconv"
)

type Request struct {
	JsonRpc string          `json:"jsonrpc"`
	Id      int             `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type Response struct {
	JsonRpc string          `json:"jsonrpc"`
	Id      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
}

type LogQuery struct {
	BlockHash *common.Hash   `json:"blockHash,omitempty"`
	FromBlock *string        `json:"fromBlock,omitempty"`
	ToBlock   *string        `json:"toBlock,omitempty"`
	Address   *interface{}   `json:"address,omitempty"`
	Topics    *[]interface{} `json:"topics,omitempty"`
}

type Log struct {
	Address     common.Address `json:"address"`
	Topics      []common.Hash  `json:"topics"`
	Data        hexutil.Bytes  `json:"data"`
	BlockNumber HexUint        `json:"blockNumber"`
	TxHash      common.Hash    `json:"transactionHash"`
	TxIndex     HexUint        `json:"transactionIndex"`
	BlockHash   common.Hash    `json:"blockHash"`
	LogIndex    HexUint        `json:"logIndex"`
	Removed     bool           `json:"removed"`
}

type LogString struct {
	Address     string   `json:"address"`
	Topics      []string `json:"topics"`
	Data        string   `json:"data"`
	BlockNumber string   `json:"blockNumber"`
	TxHash      string   `json:"transactionHash"`
	TxIndex     string   `json:"transactionIndex"`
	BlockHash   string   `json:"blockHash"`
	LogIndex    string   `json:"logIndex"`
	Removed     bool     `json:"removed"`
}

func (l *Log) ConvertToLogString() LogString {
	topics := make([]string, 0, len(l.Topics))
	for _, topic := range l.Topics {
		topics = append(topics, topic.Hex())
	}
	return LogString{
		Address:     l.Address.Hex(),
		Topics:      topics,
		Data:        hex.EncodeToString(l.Data),
		BlockNumber: strconv.FormatUint(uint64(l.BlockNumber), 10),
		TxHash:      l.TxHash.Hex(),
		TxIndex:     strconv.FormatUint(uint64(l.TxIndex), 10),
		BlockHash:   l.BlockHash.Hex(),
		LogIndex:    strconv.FormatUint(uint64(l.LogIndex), 10),
		Removed:     l.Removed,
	}
}

type HexUint uint

func (h *HexUint) UnmarshalJSON(data []byte) error {
	str := string(data[1 : len(data)-1])

	val, err := strconv.ParseUint(str, 0, 64)
	if err != nil {
		return err
	}

	*h = HexUint(val)
	return nil
}
