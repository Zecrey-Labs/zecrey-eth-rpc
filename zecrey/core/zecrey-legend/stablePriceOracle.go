package zecreyLegend

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"math/big"
)

func EmptyCallOpts() *bind.CallOpts {
	return &bind.CallOpts{}
}

/*
	LoadZecreyLegendInstance: load zecrey legend instance if it is already deployed
*/
func LoadStablePriceOracleInstance(cli *_rpc.ProviderClient, addr string) (instance *StablePriceOracle, err error) {
	instance, err = NewStablePriceOracle(common.HexToAddress(addr), *cli)
	return instance, err
}

func Price(
	instance *StablePriceOracle,
	name string,
) (amount *big.Int, err error) {
	amount, err = instance.Price(EmptyCallOpts(), name)
	if err != nil {
		return nil, err
	}
	return amount, nil
}
