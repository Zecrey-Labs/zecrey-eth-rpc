package zecreyLegend

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"testing"
)

var (
	cli, _      = _rpc.NewClient("http://tf-dex-preview-validator-nlb-6fd109ac8b9d390a.elb.ap-northeast-1.amazonaws.com:8545")
	instance, _ = LoadZecreyLegendInstance(cli, "0x630267531f9E89D93a6e117F03D38c62e37a9F82")
)

func TestLoadZecreyLegendInstance(t *testing.T) {
	totalBlocksCommitted, err := instance.TotalBlocksCommitted(EmptyCallOpts())
	if err != nil {
		panic(err)
	}
	fmt.Println(totalBlocksCommitted)
}

func TestLoadZecreyLegendInstance2(t *testing.T) {
	storedBlockHashes, err := instance.StoredBlockHashes(EmptyCallOpts(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(common.Bytes2Hex(storedBlockHashes[:]))
}
