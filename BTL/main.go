package main

import (
	"NFT1218Admin/global"
	"NFT1218Admin/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

const (
	ApiKey               = "FMNAVZAI6WZYIYE3M1WMP2UZXT9XRQZHTZ" // 替换为你的BscScan API密钥
	BscScanAPIBaseURL    = "https://api.bscscan.com/api"
	Grpc                 = "https://bsc-dataseed.bnbchain.org"
	TokenContractAddress = "0x167e87635C70134902f63780a77F609Dcb7149aC"
	LpContractAddress    = "0x612B0D34F78681b14a4F6fFA41de6B299402bf9A"
	PancakeRouterAddress = "0x10ED43C718714eb63d5aA57B78B54704E256024E"
	LpAddress            = "0x872eCfF6F0C01f5A73EaAFd69a34Aab583a00Ef8"
	//NFTContractAddress   = "0xEe3E55aB6dce0854810d6D3DaE54D5e349Cf3b9f"
	NFTAddress = "0xBe7906FB9346755Db2Fa5ACddF9808e18a5B6858"

	USDTContractAddress = "0x55d398326f99059fF775485246999027B3197955"
)

func init() {
	err := global.ConnectToEthereumNode(Grpc)
	if err != nil {
		log.Fatalf("连接节点失败 err:%v", err)
	}
	log.Println("连接节点成功")

	PancakeContractAddress := common.HexToAddress(LpContractAddress) // 替换为您的合约地址
	err = global.LoadLPContract(PancakeContractAddress)
	if err != nil {
		log.Fatalf("加载LP合约失败失败 err:%v", err)
	}
	log.Println("加载LP合约成功")

	BTLContractAddress := common.HexToAddress(TokenContractAddress) // 替换为您的合约地址
	err = global.LoadBTLContract(BTLContractAddress)
	if err != nil {
		log.Fatalf("加载BTL合约失败失败 err:%v", err)
	}
	log.Println("加载BTL合约成功")

	//NFTContractAddress1 := common.HexToAddress(NFTContractAddress) // 替换为您的合约地址
	//err = global.LoadNFTContract(NFTContractAddress1)
	//if err != nil {
	//	log.Fatalf("加载NFT合约失败失败 err:%v", err)
	//}
	//log.Println("加载NFT合约成功")

	PancakeRouterAddress1 := common.HexToAddress(PancakeRouterAddress) // 替换为您的合约地址
	err = global.LoadPancakeRouterContract(PancakeRouterAddress1)
	if err != nil {
		log.Fatalf("加载pancake路由合约失败失败 err:%v", err)
	}
	log.Println("加载pancake路由合约成功")

	USDTContractAddress1 := common.HexToAddress(USDTContractAddress) // 替换为您的合约地址
	err = global.LoadUSDTContract(USDTContractAddress1)
	if err != nil {
		log.Fatalf("加载USDTContract失败 err:%v", err)
	}
	log.Println("加载USDTContract合约成功")
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	// 定义一个路由，处理GET请求
	r.GET("/", func(c *gin.Context) {
		admin := util.GetBalanceBTL(TokenContractAddress, USDTContractAddress, LpContractAddress, LpAddress, NFTAddress)
		LpHolder := util.GetLPHolders(TokenContractAddress, BscScanAPIBaseURL, ApiKey)
		admin.LpCount = len(LpHolder)
		c.JSON(200, gin.H{
			"message": admin,
		})
	})

	// 启动Gin服务，默认监听在 :8080 端口
	r.Run(":9090")
}

//package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"strings"
//)
//
//const (
//	apiKey            = "FMNAVZAI6WZYIYE3M1WMP2UZXT9XRQZHTZ" // Replace with your BscScan API key
//	bscScanAPIBaseURL = "https://api-testnet.bscscan.com/api"
//	transferEventHash = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
//	AdminContract     = "0xxxxxx"
//)
//
//var (
//	pairContract = "0xD3802033a9df0C0C6A6E272e4B51C98dA8826d52"
//)
//
//type LogInfo struct {
//	From string
//	To   string
//}
//
//func main() {
//	var nonZeroTopics []LogInfo
//	getTransferLogs(pairContract, &nonZeroTopics)
//	var remainingTopics []LogInfo
//	fmt.Println("Non-Zero Topics:")
//	for _, info := range nonZeroTopics {
//		if info.From == strings.ToLower(pairContract) || info.To == strings.ToLower(pairContract) || info.From == strings.ToLower(AdminContract) || info.To == strings.ToLower(AdminContract) {
//			continue
//		}
//		remainingTopics = append(remainingTopics, info)
//	}
//	nonZeroTopics = remainingTopics
//	for _, info := range nonZeroTopics {
//		fmt.Printf("- From: %s, To: %s\n", info.From, info.To)
//	}
//
//}
//
//func getTransferLogs(contractAddress string, nonZeroTopics *[]LogInfo) {
//	apiURL := fmt.Sprintf("%s?module=logs&action=getLogs&address=%s&apikey=%s", bscScanAPIBaseURL, contractAddress, apiKey)
//
//	response, err := http.Get(apiURL)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer response.Body.Close()
//
//	body, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var result map[string]interface{}
//	err = json.Unmarshal(body, &result)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	status, ok := result["status"].(string)
//	if !ok || status != "1" {
//		log.Println("API request failed")
//		return
//	}
//
//	logs, ok := result["result"].([]interface{})
//	if !ok {
//		log.Println("Failed to parse logs")
//		return
//	}
//
//	for _, log := range logs {
//		logInfo := log.(map[string]interface{})
//		topics, ok := logInfo["topics"].([]interface{})
//
//		// Check if the log is a Transfer event
//		if ok && len(topics) > 0 && topics[0].(string) == transferEventHash {
//			counter := 0
//			hasZero := false
//			loginfo := &LogInfo{}
//			for _, topic := range topics {
//				trimmedTopic := removeLeadingZeros(topic.(string))
//
//				// Check if the trimmed topic is "0x"
//				if trimmedTopic == "0x" {
//					hasZero = true
//					break
//				}
//				switch counter {
//				case 1:
//					loginfo.From = trimmedTopic
//				case 2:
//					loginfo.To = trimmedTopic
//				}
//				counter++
//			}
//
//			// If none of the topics is "0x", add to the nonZeroTopics array
//			if !hasZero {
//				*nonZeroTopics = append(*nonZeroTopics, *loginfo)
//			}
//		}
//	}
//}
//
//func removeLeadingZeros(hexString string) string {
//	// Remove the '0x' prefix
//	hexString = hexString[2:]
//
//	// Trim leading zeros
//	trimmedString := strings.TrimLeft(hexString, "0")
//
//	// Add '0x' prefix back
//	trimmedString = "0x" + trimmedString
//
//	return trimmedString
//}
