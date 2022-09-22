package zecreyLegend

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"math/big"
)

func ConstructTransactOpts(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, gasPrice *big.Int, gasLimit uint64) (transactOpts *bind.TransactOpts, err error) {
	transactOpts, err = bind.NewKeyedTransactorWithChainID(authCli.PrivateKey, authCli.ChainId)
	if err != nil {
		return nil, err
	}
	nonce, err := cli.GetPendingNonce(authCli.Address.Hex())
	if err != nil {
		return nil, err
	}
	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = gasLimit
	transactOpts.From = authCli.Address
	transactOpts.Value = big.NewInt(0)
	return transactOpts, nil
}


func ConstructTransactOptsWithValue(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, gasPrice *big.Int, gasLimit uint64, value int64) (transactOpts *bind.TransactOpts, err error) {
	transactOpts, err = bind.NewKeyedTransactorWithChainID(authCli.PrivateKey, authCli.ChainId)
	if err != nil {
		return nil, err
	}
	nonce, err := cli.GetPendingNonce(authCli.Address.Hex())
	if err != nil {
		return nil, err
	}
	if value < 0 {
		return nil, errors.New("value is negative")
	}

	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = gasLimit
	transactOpts.From = authCli.Address
	transactOpts.Value = big.NewInt(value)
	return transactOpts, nil
}
