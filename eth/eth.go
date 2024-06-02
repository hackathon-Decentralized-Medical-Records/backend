package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"service/eth/contracts"
	"service/eth/contracts/role"
	"strings"
	"time"
)

// 1.先注册账号
// 2.注册链接: https://www.infura.io/zh
// 3.通知部署合约的同学给出abi和合约地址
// 4.将 ABI 通过 工具 编译成 go 代码
// 5. abigen --abi=path//YourContract.abi.json --pkg=yourpackage --out=path/yourpackage/YourContract.go
// 6.编写业务代码
//

type RoleType uint8

func AddContract(c *gin.Context) {
	client, err := ethclient.DialContext(context.Background(), "wss://sepolia.infura.io/ws/v3/7de12dd4579a405f85cf228a0c9beebf")
	if err != nil {
		log.Fatal(err)
	}
	//// 合约地址
	contractAddress := common.HexToAddress("0xEe524AF17809c97F1C803b69BB0e20d414588aB1")

	privateKey, err := crypto.HexToECDSA("35dd82fefef019ac146d2969a6e571e60cee7bc66dc360b1d1e498ef1dedd669")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	contractAbi, err := abi.JSON(strings.NewReader(string(role.ContractsABI)))

	txData, err := contractAbi.Pack("addNewUserToSystem", uint8(1))
	tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), auth.GasLimit, auth.GasPrice, txData)

	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())

}

func NewContract(c *gin.Context) {
	client, err := ethclient.DialContext(context.Background(), "wss://sepolia.infura.io/ws/v3/7de12dd4579a405f85cf228a0c9beebf")
	if err != nil {
		log.Fatal(err)
	}
	//// 合约地址
	address := common.HexToAddress("0xEe524AF17809c97F1C803b69BB0e20d414588aB1")
	//// 创建合约实例
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			address,
		},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("监听中........", time.Now())

	// 监听
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}

func GetPatient(c *gin.Context) {
	client, err := ethclient.DialContext(context.Background(), "wss://sepolia.infura.io/ws/v3/7de12dd4579a405f85cf228a0c9beebf")
	if err != nil {
		log.Fatal(err)
	}
	//// 合约地址
	contractAddress := common.HexToAddress("0xEe524AF17809c97F1C803b69BB0e20d414588aB1")
	instance, err := contracts.NewContracts(contractAddress, client)

	privateKey, err := crypto.HexToECDSA("5dc84b9b715e80fee3d22d34079336c66d20d5bbf288844b06cc20d725d25dc3")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println(instance)
	if err != nil {
		log.Fatal(err)
	}

	callOpts := &bind.CallOpts{
		Pending: false, // 或者根据你的需求设置其他选项
		// From:    common.HexToAddress("0xYourAddress"), // 可选
		// Context: context.Background(), // 可选
	}
	// 调用合约方法
	// 患者合约地址
	contract, err := instance.GetUserToContract(callOpts, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract:", contract)
	newContracts, err := role.NewContracts(contract, client)
	if err != nil {
		log.Fatal(err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	fmt.Println(opts)
	opts.GasLimit = uint64(3000000)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	address := common.HexToAddress("0xEBc29CA8Dd95DF5790746e172F79734D7D524898")
	material, err := newContracts.SetApprovalForAddingMaterial(opts, address)

	fmt.Println("material:", material.Data())

	//// 时间戳 转成bin。int
	timestamp := time.Now().Unix()
	timestampBigInt := big.NewInt(timestamp)
	opts.Value = big.NewInt(100000000000000) // 1 ETH，替换为你需要的值

	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatalf("Failed to retrieve account balance: %v", err)
	}
	fmt.Printf("Account balance: %s wei\n", balance.String())
	instance.RequestReservation(opts, address, timestampBigInt)

	// 设置调用选项
	infoOpt := &bind.CallOpts{
		From:    fromAddress,
		Context: context.Background(),
	}

	// 医生账户地址
	info, err := instance.GetReservationInfo(infoOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("info:", info.Hex())
}
