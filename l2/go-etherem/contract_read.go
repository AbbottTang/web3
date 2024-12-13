package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "go-etherem/story"
)

// 现在我们将使用新合约实例提供的方法来阅读智能合约。
// 如果你还记得我们在部署过程中设置的合约中有一个名为 version 的全局变量。
// 因为它是公开的，这意味着它们将成为我们自动创建的 getter 函数。
// 常量和 view 函数也接受 bind.CallOpts 作为第一个参数。
// 了解可用的具体选项要看相应类的文档 一般情况下我们可以用 nil。
func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/d36c22d4559d427e821074165c1f07fa")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version) // "1.0"
}
