package util

import (
	"NFT1218Admin/global"
	"fmt"
	"github.com/GoPlusSecurity/goplus-sdk-go/api/token"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/errorcode"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Admin struct {
	BTLPrice     string
	PoolsNum     string
	LpCount      int
	HolderCount  string
	LPDividends  string
	NftDividends string
}

func GetBalanceBTL(BTLContract, USDTContract, lptoken, lpaddress, nftprivate string) *Admin {
	var admin = &Admin{}
	path := []common.Address{common.HexToAddress(BTLContract), common.HexToAddress(USDTContract)}
	num2 := big.NewInt(1000000000000000000)
	OutAmount, _ := global.PancakeRouter.GetAmountsOut(nil, num2, path)

	amountOutMin := OutAmount[1]

	admin.BTLPrice = amountOutMin.String()

	lpbanlance, _ := global.USDTToken.BalanceOf(nil, common.HexToAddress(lptoken))

	admin.PoolsNum = lpbanlance.String()

	tokenSecurity := token.NewTokenSecurity(nil)
	chainId := "56"
	contractAddresses := []string{BTLContract}
	data, err := tokenSecurity.Run(chainId, contractAddresses)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	if data.Payload.Code != errorcode.SUCCESS {
		panic(data.Payload.Message)
	}
	//data.Payload.Result[0].HolderCount
	for _, value := range data.Payload.Result {
		admin.HolderCount = value.HolderCount
	}

	btlbalance, err := global.BTLToken.BalanceOf(nil, common.HexToAddress(lpaddress))
	lpUsdtBalance, err := global.USDTToken.BalanceOf(nil, common.HexToAddress(lpaddress))
	if err != nil {
		fmt.Println(" global.BTLToken.BalanceOf err:", err.Error())
	}
	if btlbalance.Cmp(big.NewInt(0)) == 0 {
		admin.LPDividends = lpUsdtBalance.String()
	} else {
		OutAmount1, err := global.PancakeRouter.GetAmountsOut(nil, btlbalance, path)
		if err != nil {
			fmt.Println("  global.PancakeRouter.GetAmountsOut err:", err.Error())
		}

		admin.LPDividends = OutAmount1[1].Add(OutAmount1[1], lpUsdtBalance).String()
	}

	btlbalance1, err := global.BTLToken.BalanceOf(nil, common.HexToAddress(nftprivate))
	nftUsdtBalance, err := global.USDTToken.BalanceOf(nil, common.HexToAddress(nftprivate))

	if err != nil {
		fmt.Println(" global.BTLToken.BalanceOf err:", err.Error())
	}
	if btlbalance1.Cmp(big.NewInt(0)) == 0 {
		admin.NftDividends = nftUsdtBalance.String()
	} else {
		OutAmount2, err := global.PancakeRouter.GetAmountsOut(nil, btlbalance1, path)
		if err != nil {
			fmt.Println("  global.PancakeRouter.GetAmountsOut err:", err.Error())
		}

		admin.NftDividends = OutAmount2[1].Add(OutAmount2[1], nftUsdtBalance).String()
	}

	return admin

}
