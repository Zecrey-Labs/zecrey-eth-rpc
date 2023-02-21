package zecreyLegend

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecrey/core"
	"math/big"
)

/*
LoadZecreyLegendInstance: load zecrey legend instance if it is already deployed
*/
func LoadZecreyLegendInstance(cli *_rpc.ProviderClient, addr string) (instance *ZecreyLegend, err error) {
	instance, err = NewZecreyLegend(common.HexToAddress(addr), *cli)
	return instance, err
}

/*
CommitBlocks: commit blocks
*/
func CommitBlocks(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *ZecreyLegend,
	lastBlock StorageStoredBlockInfo, commitBlocksInfo []OldZecreyLegendCommitBlockInfo,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := core.ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.CommitBlocks(transactOpts, lastBlock, commitBlocksInfo)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
VerifyAndExecuteBlocks: verify and execute blocks
*/
func VerifyAndExecuteBlocks(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *ZecreyLegend,
	verifyAndExecuteBlocksInfo []OldZecreyLegendVerifyAndExecuteBlockInfo, proofs []*big.Int,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := core.ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.VerifyAndExecuteBlocks(transactOpts, verifyAndExecuteBlocksInfo, proofs)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func RevertBlocks(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *ZecreyLegend,
	revertBlocks []StorageStoredBlockInfo,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := core.ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	tx, err := instance.RevertBlocks(transactOpts, revertBlocks)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func RegisterZNS(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *ZecreyLegend, oracleInstance *StablePriceOracle,
	gasPrice *big.Int, gasLimit uint64,
	name string, ownerAddr common.Address, pkX [32]byte, pkY [32]byte,
) (txHash string, err error) {
	amount, err := Price(oracleInstance, name)
	if err != nil {
		return "", err
	}
	transactOpts, err := core.ConstructTransactOptsWithValue(cli, authCli, gasPrice, gasLimit, amount.Int64())
	if err != nil {
		return "", err
	}
	tx, err := instance.RegisterZNS(transactOpts, name, ownerAddr, pkX, pkY)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}
