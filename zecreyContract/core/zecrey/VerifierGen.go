// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zecrey

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VerifierABI is the input ABI used to generate the binding from.
const VerifierABI = "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"input\",\"type\":\"uint256[2]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VerifierBin is the compiled bytecode used for deploying new contracts.
var VerifierBin = "0x60806040523480156100115760006000fd5b50610017565b6116a8806100266000396000f3fe60806040523480156100115760006000fd5b50600436106100305760003560e01c8063f5c9d69e1461003657610030565b60006000fd5b6101a6600480360361014081101561004e5760006000fd5b8101908080604001906002806020026040519081016040528092919082600260200280828437600081840152601f19601f8201169050808301925050505050509090919290909192908060800190600280602002604051908101604052809291906000905b82821015610107578382604002016002806020026040519081016040528092919082600260200280828437600081840152601f19601f820116905080830192505050505050815260200190600101906100b3565b5050505090909192909091929080604001906002806020026040519081016040528092919082600260200280828437600081840152601f19601f82011690508083019250505050505090909192909091929080604001906002806020026040519081016040528092919082600260200280828437600081840152601f19601f8201169050808301925050505050509090919290909192905050506101be565b60405180821515815260200191505060405180910390f35b60006101c86114f1565b60405180604001604052808760006002811015156101e257fe5b602002015181526020018760016002811015156101fb57fe5b602002015181526020015081600001819052506040518060400160405280604051806040016040528088600060028110151561023357fe5b6020020151600060028110151561024657fe5b6020020151815260200188600060028110151561025f57fe5b6020020151600160028110151561027257fe5b60200201518152602001508152602001604051806040016040528088600160028110151561029c57fe5b602002015160006002811015156102af57fe5b602002015181526020018860016002811015156102c857fe5b602002015160016002811015156102db57fe5b60200201518152602001508152602001508160200181905250604051806040016040528085600060028110151561030e57fe5b6020020151815260200185600160028110151561032757fe5b60200201518152602001508160400181905250600061034a610a6163ffffffff16565b90506000604051806040016040528060008152602001600081526020015090507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4783600001516000015110151561040c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f76657269666965722d61582d6774652d7072696d652d7100000000000000000081526020015060200191505060405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478360000151602001511015156104ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f76657269666965722d61592d6774652d7072696d652d7100000000000000000081526020015060200191505060405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4783602001516000015160006002811015156104e657fe5b6020020151101515610563576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f76657269666965722d6258302d6774652d7072696d652d71000000000000000081526020015060200191505060405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47836020015160200151600060028110151561059b57fe5b6020020151101515610618576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f76657269666965722d6259302d6774652d7072696d652d71000000000000000081526020015060200191505060405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47836020015160000151600160028110151561065057fe5b60200201511015156106cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f76657269666965722d6258312d6774652d7072696d652d71000000000000000081526020015060200191505060405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47836020015160200151600160028110151561070557fe5b6020020151101515610782576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f76657269666965722d6259312d6774652d7072696d652d71000000000000000081526020015060200191505060405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47836040015160000151101515610824576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f76657269666965722d63582d6774652d7072696d652d7100000000000000000081526020015060200191505060405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478360400151602001511015156108c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f76657269666965722d63592d6774652d7072696d652d7100000000000000000081526020015060200191505060405180910390fd5b6000600090505b855060028110156109e0577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001868260028110151561090757fe5b6020020151101515610984576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f76657269666965722d6774652d736e61726b2d7363616c61722d6669656c640081526020015060200191505060405180910390fd5b6109ce826109c38560800151600185016003811015156109a057fe5b602002015189856002811015156109b357fe5b6020020151610e8e63ffffffff16565b610fa463ffffffff16565b915081505b80806001019150506108cd565b50610a0881836080015160006003811015156109f857fe5b6020020151610fa463ffffffff16565b90508050610a4c610a2284600001516110e063ffffffff16565b846020015184600001518560200151858760400151896040015189606001516111a463ffffffff16565b9350505050610a59565050505b949350505050565b610a69611527565b60405180604001604052807f0c7509c153ce226f510b4f599eaf17a5ae1a59c95d40a240b370ab38ca75370481526020017f1aa5ff7f04e3b8132032ced67a5013b5cc2b508c75e58e6fc28aa364596191bc8152602001508160000181905250604051806040016040528060405180604001604052807f022bd68e4b74a5a21cd3978d57a7f5d41443e32ad47eba3acf504d6f352b295c81526020017f26e1f310e560cd871efcfda76ee3e23eee1f1fd08cde0cd323080a08b765195f815260200150815260200160405180604001604052807f2b74c16cec250634d82672a28e161a80fd796bfa6c554ac72e9b0696d7925a0c81526020017f0cda2e7448fd8312f189ac1b17dd60cc7752782215659f0604ad30308377c5678152602001508152602001508160200181905250604051806040016040528060405180604001604052807f1510119c86df39d9a77da91550581aab38f8828a1eabce6285b380cbf998570381526020017ec77eb3c390b84ad18b8ba5c2af50fb4792af36b1c2fd3dbdc6ee5b5ec51506815260200150815260200160405180604001604052807f1f97654a2ad8ec22053a824fe6e796b8c2ffbc224f835b79bee2b591fe73e6f581526020017f139665663c7eab792cceb88de58ba7f3a83171beebe888daf54c8875ce53464b8152602001508152602001508160400181905250604051806040016040528060405180604001604052807f0efd410606bc9137ee5f38e08f6ac2ec2b3cc35b37a0c63598891d0156f0966581526020017f2b456c7aed96a7ed9f2349ff15bbb4012b65215663d2173cb732b5d018722498815260200150815260200160405180604001604052807e92a4155a9ed4a07b53f6320b93637c36e954f8970e92e2fa06075750c42a3281526020017f1f03264a4f869e2108d5bbb392f88186657d02751ff79a31bac99ec4a9de1153815260200150815260200150816060018190525060405180604001604052807f0f9bd37b3a0136a23f716a8f8d9ccb179ffe7cd6fb247405696207bfce026ff581526020017f12ae39135bf9e01fb9bd89af169a65ba465eff07b396e27332a1345c72733aee81526020015081608001516000600381101515610d9c57fe5b602002018190525060405180604001604052807f228d861a32ecca69d116a73b9862f8bc69f8c7b6d718b0aa422c6c70e54c857581526020017f242ec2d5211a3a3a3435178c300a100f049c839b77a97807aa2747338926e88c81526020015081608001516001600381101515610e0f57fe5b602002018190525060405180604001604052807f1b69b846c07bdb7a0502f49742171e7f55b5c4d62f566fe51c0756489b27f77581526020017f2fa40c58a1d862c05b8efdcedf80609f3099f925cad35a91ae730e938c1562f281526020015081608001516002600381101515610e8257fe5b60200201819052505b90565b610e96611577565b610e9e611594565b8360000151816000600381101515610eb257fe5b60200201909081815260200150508360200151816001600381101515610ed457fe5b602002019090818152602001505082816002600381101515610ef257fe5b6020020190908181526020015050600060608360808460076107d05a03fa90508060008114610f2057610f22565bfe5b50801515610f9b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f70616972696e672d6d756c2d6661696c6564000000000000000000000000000081526020015060200191505060405180910390fd5b50505b92915050565b610fac611577565b610fb46115b6565b8360000151816000600481101515610fc857fe5b60200201909081815260200150508360200151816001600481101515610fea57fe5b6020020190908181526020015050826000015181600260048110151561100c57fe5b6020020190908181526020015050826020015181600360048110151561102e57fe5b6020020190908181526020015050600060608360c08460066107d05a03fa9050806000811461105c5761105e565bfe5b508015156110d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f70616972696e672d6164642d6661696c6564000000000000000000000000000081526020015060200191505060405180910390fd5b50505b92915050565b6110e8611577565b60008260000151148015611100575060008260200151145b1561112a576040518060400160405280600081526020016000815260200150905061119f5661119e565b6040518060400160405280836000015181526020017f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47846020015181151561116e57fe5b067f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4703815260200150905061119f565b5b919050565b6000600060405180608001604052808b8152602001898152602001878152602001858152602001509050600060405180608001604052808b815260200189815260200187815260200185815260200150905060006018905060008167ffffffffffffffff811180156112165760006000fd5b506040519080825280602002602001820160405280156112455781602001602082028036833780820191505090505b5090506000600090505b6004811015611411576000600682029050858260048110151561126e57fe5b602002015160000151836000830181518110151561128857fe5b602002602001019090818152602001505085826004811015156112a757fe5b60200201516020015183600183018151811015156112c157fe5b602002602001019090818152602001505084826004811015156112e057fe5b60200201516000015160006002811015156112f757fe5b6020020151836002830181518110151561130d57fe5b6020026020010190908181526020015050848260048110151561132c57fe5b602002015160000151600160028110151561134357fe5b6020020151836003830181518110151561135957fe5b6020026020010190908181526020015050848260048110151561137857fe5b602002015160200151600060028110151561138f57fe5b602002015183600483018151811015156113a557fe5b602002602001019090818152602001505084826004811015156113c457fe5b60200201516020015160016002811015156113db57fe5b602002015183600583018151811015156113f157fe5b6020026020010190908181526020015050505b808060010191505061124f565b5061141a6115d8565b6000602082602086026020860160086107d05a03fa9050806000811461143f57611441565bfe5b508015156114ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f70616972696e672d6f70636f64652d6661696c6564000000000000000000000081526020015060200191505060405180910390fd5b60008260006001811015156114cb57fe5b6020020151141596505050505050506114e5565050505050505b98975050505050505050565b6040518060600160405280611504611577565b81526020016115116115fa565b815260200161151e611577565b81526020015090565b6040518060a0016040528061153a611577565b81526020016115476115fa565b81526020016115546115fa565b81526020016115616115fa565b815260200161156e611623565b81526020015090565b604051806040016040528060008152602001600081526020015090565b6040518060600160405280600390602082028036833780820191505090505090565b6040518060800160405280600490602082028036833780820191505090505090565b6040518060200160405280600190602082028036833780820191505090505090565b604051806040016040528061160d611650565b815260200161161a611650565b81526020015090565b60405180606001604052806003905b61163a611577565b8152602001906001900390816116325790505090565b604051806040016040528060029060208202803683378082019150509050509056fea2646970667358221220e9307b43b94b1c47033a71a0f04f0cba35d5fc1ee7d0fde8fb6b09aea753fbef64736f6c63430007060033"

// DeployVerifier deploys a new Ethereum contract, binding an instance of Verifier to it.
func DeployVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verifier, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// Verifier is an auto generated Go binding around an Ethereum contract.
type Verifier struct {
	VerifierCaller     // Read-only binding to the contract
	VerifierTransactor // Write-only binding to the contract
	VerifierFilterer   // Log filterer for contract events
}

// VerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierSession struct {
	Contract     *Verifier         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierCallerSession struct {
	Contract *VerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// VerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierTransactorSession struct {
	Contract     *VerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierRaw struct {
	Contract *Verifier // Generic contract binding to access the raw methods on
}

// VerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierCallerRaw struct {
	Contract *VerifierCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierTransactorRaw struct {
	Contract *VerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifier creates a new instance of Verifier, bound to a specific deployed contract.
func NewVerifier(address common.Address, backend bind.ContractBackend) (*Verifier, error) {
	contract, err := bindVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// NewVerifierCaller creates a new read-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierCaller(address common.Address, caller bind.ContractCaller) (*VerifierCaller, error) {
	contract, err := bindVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierCaller{contract: contract}, nil
}

// NewVerifierTransactor creates a new write-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierTransactor, error) {
	contract, err := bindVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierTransactor{contract: contract}, nil
}

// NewVerifierFilterer creates a new log filterer instance of Verifier, bound to a specific deployed contract.
func NewVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierFilterer, error) {
	contract, err := bindVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierFilterer{contract: contract}, nil
}

// bindVerifier binds a generic wrapper to an already deployed contract.
func bindVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.VerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf5c9d69e.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] input) view returns(bool r)
func (_Verifier *VerifierCaller) VerifyProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [2]*big.Int) (bool, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "verifyProof", a, b, c, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0xf5c9d69e.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] input) view returns(bool r)
func (_Verifier *VerifierSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [2]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf5c9d69e.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] input) view returns(bool r)
func (_Verifier *VerifierCallerSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [2]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}