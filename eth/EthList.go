package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"service/config"
)

// 1.先注册账号
// 2.注册链接: https://www.infura.io/zh
// 3.通知部署合约的同学给出abi和合约地址
// 4.将 ABI 通过 工具 编译成 go 代码
// 5. abigen --abi=path//YourContract.abi.json --pkg=yourpackage --out=path/yourpackage/YourContract.go
// 6.编写业务代码
//

func NewContract() {
	_, err := ethclient.DialContext(context.Background(), config.EthApiKeyUrl)
	if err != nil {
		log.Fatal(err)
	}
	//// 合约地址
	common.HexToAddress(config.EthContractUrl)
	//// 创建合约实例

	////调用合约函数

}
