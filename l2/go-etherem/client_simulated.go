package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/node"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
)

func main() {
	//client, err := ethclient.Dial("https://sepolia.infura.io/v3/d36c22d4559d427e821074165c1f07fa")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//您可以使用模拟客户端来快速轻松地在本地测试您的交易，非常适合单元测试。
	//为了开始，我们需要一个带有初始 ETH 的账户。为此，首先生成一个账户私钥。
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	//接着从 accounts/abi/bind 包创建一个 NewKeyedTransactor，并为其传递私钥。
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))

	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 eth in wei
	///下一步是创建一个创世账户并为其分配初始余额。我们将使用 core 包的 GenesisAccount 类型。
	address := auth.From
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: balance,
		},
	}

	//现在我们将创世分配结构体和配置好的汽油上限传给 account/abi/bind/backends 包的 NewSimulatedBackend 方法，该方法将返回一个新的模拟以太坊客户端。
	blockGasLimit := uint64(4712388)
	backend := simulated.NewBackend(genesisAlloc, func(nodeConf *node.Config, ethConf *ethconfig.Config) {
		// 在这里可以添加额外的配置选项
		// 例如：调整gas限制、设置挖矿奖励等
		ethConf.Genesis = &core.Genesis{
			GasLimit: blockGasLimit,
		}
	})
	client := backend.Client()

	//您可以像往常一样使用此客户端。作为一个示例，我们将构造一个新的交易并进行广播。
	fromAddress := auth.From
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = backend.Client().SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex()) // tx sent: 0xec3ceb05642c61d33fa6c951b54080d1953ac8227be81e7b5e4e2cfed69eeb51
	//到现在为止，您可能想知道交易何时才会被开采。为了“开采”它，您还必须做一件额外的事情，在客户端调用 Commit 提交新开采的区块。
	backend.Commit()

	//现在您可以获取交易收据并看见其已被处理。
	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt == nil {
		log.Fatal("receipt is nil. Forgot to commit?")
	}

	fmt.Printf("status: %v\n", receipt.Status) // status: 1
}