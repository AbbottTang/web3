// 加载智能合约
package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "go-etherem/story"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/d36c22d4559d427e821074165c1f07fa")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	//一旦使用 abigen 工具将智能合约的 ABI 编译为 Go 包，下一步就是调用“New”方法，其格式为“New”，所以在我们的例子中如果你 回想一下它将是_NewStore_。
	//此初始化方法接收智能合约的地址，并返回可以开始与之交互的合约实例。
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance
}
