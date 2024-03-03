package global

import (
	"NFT1218Admin/contract"
	"github.com/ethereum/go-ethereum/common"
)

var (
	LpContract    *contract.LpToken
	BTLToken      *contract.BTLToken
	NFTToken      *contract.NFTToken
	PancakeRouter *contract.PancakeRouter
	USDTToken     *contract.ERC
)

func LoadLPContract(contractAddress common.Address) error {
	var err error
	LpContract, err = contract.NewLpToken(contractAddress, Client)
	if err != nil {
		return err
	}
	return nil
}

func LoadBTLContract(contractAddress common.Address) error {
	var err error
	BTLToken, err = contract.NewBTLToken(contractAddress, Client)
	if err != nil {
		return err
	}
	return nil
}

func LoadNFTContract(contractAddress common.Address) error {
	var err error
	NFTToken, err = contract.NewNFTToken(contractAddress, Client)
	if err != nil {
		return err
	}
	return nil
}

func LoadPancakeRouterContract(contractAddress common.Address) error {
	var err error
	PancakeRouter, err = contract.NewPancakeRouter(contractAddress, Client)
	if err != nil {
		return err
	}
	return nil
}

func LoadUSDTContract(contractAddress common.Address) error {
	var err error
	USDTToken, err = contract.NewERC(contractAddress, Client)
	if err != nil {
		return err
	}
	return nil
}
