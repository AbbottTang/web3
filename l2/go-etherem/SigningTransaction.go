package main

import (
	"log"
	"math/big"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/miguelmota/go-ethereum-hdwallet"
)

func main() {
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, true)
	if err != nil {
		log.Fatal(err)
	}

	nonce := uint64(0)
	value := big.NewInt(1000000000000000000)
	///toAddress 被设置为 "0x0"，这是一个无效的以太坊地址。您需要提供一个有效的以太坊地址作为接收方。
	toAddress := common.HexToAddress("0x0")
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(21000000000)
	var data []byte
	///data 字段是空的，这通常用于包含智能合约的调用数据或其他交易相关的数据。如果您只是进行简单的以太币转账，这可以是空的。
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	//在调用 wallet.SignTx 时，您传递了一个 nil 作为链ID。虽然这在某些情况下可能是可以接受的（例如，当您不在特定链上签名时），
	//但通常最好指定一个链ID来确保签名的有效性
	signedTx, err := wallet.SignTx(account, tx, nil)
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(signedTx)
}
