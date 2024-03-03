package global

import "github.com/ethereum/go-ethereum/ethclient"

var (
	Client *ethclient.Client
)

func ConnectToEthereumNode(grpc string) error {
	var err error
	Client, err = ethclient.Dial(grpc) // 替换为您的以太坊节点的地址
	if err != nil {
		return err
	}
	return nil
}
