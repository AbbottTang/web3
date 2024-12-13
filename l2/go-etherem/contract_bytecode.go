package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//有时您需要读取已部署的智能合约的字节码。 由于所有智能合约字节码都存在于区块链中，因此我们可以轻松获取它。
	//
	//首先设置客户端和要读取的字节码的智能合约地址。
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/d36c22d4559d427e821074165c1f07fa")
	if err != nil {
		log.Fatal(err)
	}

	//现在你需要调用客户端的 codeAt 方法。 codeAt 方法接受智能合约地址和可选的块编号，并以字节格式返回字节码。
	contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(bytecode)) // 60806...10029
}
