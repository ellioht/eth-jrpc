package tests

import (
	"fmt"
	"github.com/ellioht/eth-jrpc/internal/jsonrpc"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_EthBlockNumber(t *testing.T) {
	client := jsonrpc.NewClient("https://rpc.sepolia.org")

	blockNum, err := client.EthBlockNumber()
	if err != nil {
		t.Error(err)
	}

	t.Log("blockNum:", blockNum)

	assert.Equal(t, blockNum > 0, true)
}

func Test_EthGetLogs(t *testing.T) {
	client := jsonrpc.NewClient("https://rpc.sepolia.org")

	blockNum, err := client.EthBlockNumber()
	if err != nil {
		t.Error(err)
	}

	eventSignature := "VerifyBatches(uint64,bytes32,address)"
	eventSignatureHash := crypto.Keccak256Hash([]byte(eventSignature)).Hex()

	fromBlock := fmt.Sprintf("0x%X", blockNum-2000)
	toBlock := "latest"
	address := interface{}("0x32d33D5137a7cFFb54c5Bf8371172bcEc5f310ff")
	topics := []interface{}{
		eventSignatureHash,
	}
	_ = topics

	query := jsonrpc.LogQuery{
		FromBlock: &fromBlock,
		ToBlock:   &toBlock,
		Address:   &address,
		Topics:    nil,
	}

	logs, err := client.EthGetLogs(query)
	if err != nil {
		t.Error(err)
	}

	var logsString []jsonrpc.LogString
	for _, log := range logs {
		logsString = append(logsString, log.ConvertToLogString())
	}

	t.Log("logs:", logsString)

	assert.Equal(t, len(logs) > 0, true)
}
