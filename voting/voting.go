// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package voting

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// VotingMetaData contains all meta data concerning the Voting contract.
var VotingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"candidateCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_candidateID\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506100546040518060400160405280600481526020017f6c6168680000000000000000000000000000000000000000000000000000000081525061009d60201b60201c565b6100986040518060400160405280600c81526020017f616e6f746865725f6c616868000000000000000000000000000000000000000081525061009d60201b60201c565b610491565b60025f8154809291906100af90610147565b9190505550604051806060016040528060025481526020018281526020015f8152505f5f60025481526020019081526020015f205f820151815f0155602082015181600101908161010091906103c2565b506040820151816002015590505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f819050919050565b5f6101518261013e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361018357610182610111565b5b600182019050919050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061020957607f821691505b60208210810361021c5761021b6101c5565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261027e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610243565b6102888683610243565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6102c36102be6102b98461013e565b6102a0565b61013e565b9050919050565b5f819050919050565b6102dc836102a9565b6102f06102e8826102ca565b84845461024f565b825550505050565b5f5f905090565b6103076102f8565b6103128184846102d3565b505050565b5b818110156103355761032a5f826102ff565b600181019050610318565b5050565b601f82111561037a5761034b81610222565b61035484610234565b81016020851015610363578190505b61037761036f85610234565b830182610317565b50505b505050565b5f82821c905092915050565b5f61039a5f198460080261037f565b1980831691505092915050565b5f6103b2838361038b565b9150826002028217905092915050565b6103cb8261018e565b67ffffffffffffffff8111156103e4576103e3610198565b5b6103ee82546101f2565b6103f9828285610339565b5f60209050601f83116001811461042a575f8415610418578287015190505b61042285826103a7565b865550610489565b601f19841661043886610222565b5f5b8281101561045f5784890151825560018201915060208501945060208101905061043a565b8683101561047c5784890151610478601f89168261038b565b8355505b6001600288020188555050505b505050505050565b6106d68061049e5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c80630121b93f1461004e57806309eef43e1461006a5780633477ee2e1461009a578063a9a981a3146100cc575b5f5ffd5b61006860048036038101906100639190610348565b6100ea565b005b610084600480360381019061007f91906103cd565b610243565b6040516100919190610412565b60405180910390f35b6100b460048036038101906100af9190610348565b610260565b6040516100c3939291906104aa565b60405180910390f35b6100d461030b565b6040516100e191906104e6565b60405180910390f35b60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610174576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016b90610549565b60405180910390fd5b5f8111801561018557506002548111155b6101c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101bb906105b1565b60405180910390fd5b6001805f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f5f8281526020019081526020015f206002015f81548092919061023b906105fc565b919050555050565b6001602052805f5260405f205f915054906101000a900460ff1681565b5f602052805f5260405f205f91509050805f01549080600101805461028490610670565b80601f01602080910402602001604051908101604052809291908181526020018280546102b090610670565b80156102fb5780601f106102d2576101008083540402835291602001916102fb565b820191905f5260205f20905b8154815290600101906020018083116102de57829003601f168201915b5050505050908060020154905083565b60025481565b5f5ffd5b5f819050919050565b61032781610315565b8114610331575f5ffd5b50565b5f813590506103428161031e565b92915050565b5f6020828403121561035d5761035c610311565b5b5f61036a84828501610334565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61039c82610373565b9050919050565b6103ac81610392565b81146103b6575f5ffd5b50565b5f813590506103c7816103a3565b92915050565b5f602082840312156103e2576103e1610311565b5b5f6103ef848285016103b9565b91505092915050565b5f8115159050919050565b61040c816103f8565b82525050565b5f6020820190506104255f830184610403565b92915050565b61043481610315565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61047c8261043a565b6104868185610444565b9350610496818560208601610454565b61049f81610462565b840191505092915050565b5f6060820190506104bd5f83018661042b565b81810360208301526104cf8185610472565b90506104de604083018461042b565b949350505050565b5f6020820190506104f95f83018461042b565b92915050565b7f596f7520616c726561647920766f746564206265666f726520210000000000005f82015250565b5f610533601a83610444565b915061053e826104ff565b602082019050919050565b5f6020820190508181035f83015261056081610527565b9050919050565b7f43616e64696461746520646f65736e27742065786973740000000000000000005f82015250565b5f61059b601783610444565b91506105a682610567565b602082019050919050565b5f6020820190508181035f8301526105c88161058f565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61060682610315565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610638576106376105cf565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061068757607f821691505b60208210810361069a57610699610643565b5b5091905056fea26469706673582212206c2dd3ce822da1b893c380fbb67f1a3d250b0f0c2e935f2462dd9f027ed784fd64736f6c634300081e0033",
}

// VotingABI is the input ABI used to generate the binding from.
// Deprecated: Use VotingMetaData.ABI instead.
var VotingABI = VotingMetaData.ABI

// VotingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VotingMetaData.Bin instead.
var VotingBin = VotingMetaData.Bin

// DeployVoting deploys a new Ethereum contract, binding an instance of Voting to it.
func DeployVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Voting, error) {
	parsed, err := VotingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VotingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// Voting is an auto generated Go binding around an Ethereum contract.
type Voting struct {
	VotingCaller     // Read-only binding to the contract
	VotingTransactor // Write-only binding to the contract
	VotingFilterer   // Log filterer for contract events
}

// VotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingSession struct {
	Contract     *Voting           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingCallerSession struct {
	Contract *VotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingTransactorSession struct {
	Contract     *VotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingRaw struct {
	Contract *Voting // Generic contract binding to access the raw methods on
}

// VotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingCallerRaw struct {
	Contract *VotingCaller // Generic read-only contract binding to access the raw methods on
}

// VotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingTransactorRaw struct {
	Contract *VotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoting creates a new instance of Voting, bound to a specific deployed contract.
func NewVoting(address common.Address, backend bind.ContractBackend) (*Voting, error) {
	contract, err := bindVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// NewVotingCaller creates a new read-only instance of Voting, bound to a specific deployed contract.
func NewVotingCaller(address common.Address, caller bind.ContractCaller) (*VotingCaller, error) {
	contract, err := bindVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingCaller{contract: contract}, nil
}

// NewVotingTransactor creates a new write-only instance of Voting, bound to a specific deployed contract.
func NewVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingTransactor, error) {
	contract, err := bindVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingTransactor{contract: contract}, nil
}

// NewVotingFilterer creates a new log filterer instance of Voting, bound to a specific deployed contract.
func NewVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingFilterer, error) {
	contract, err := bindVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingFilterer{contract: contract}, nil
}

// bindVoting binds a generic wrapper to an already deployed contract.
func bindVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VotingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.VotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transact(opts, method, params...)
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() view returns(uint256)
func (_Voting *VotingCaller) CandidateCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "candidateCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() view returns(uint256)
func (_Voting *VotingSession) CandidateCount() (*big.Int, error) {
	return _Voting.Contract.CandidateCount(&_Voting.CallOpts)
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() view returns(uint256)
func (_Voting *VotingCallerSession) CandidateCount() (*big.Int, error) {
	return _Voting.Contract.CandidateCount(&_Voting.CallOpts)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(uint256 id, string name, uint256 voteCount)
func (_Voting *VotingCaller) Candidates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	Name      string
	VoteCount *big.Int
}, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "candidates", arg0)

	outstruct := new(struct {
		Id        *big.Int
		Name      string
		VoteCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.VoteCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(uint256 id, string name, uint256 voteCount)
func (_Voting *VotingSession) Candidates(arg0 *big.Int) (struct {
	Id        *big.Int
	Name      string
	VoteCount *big.Int
}, error) {
	return _Voting.Contract.Candidates(&_Voting.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(uint256 id, string name, uint256 voteCount)
func (_Voting *VotingCallerSession) Candidates(arg0 *big.Int) (struct {
	Id        *big.Int
	Name      string
	VoteCount *big.Int
}, error) {
	return _Voting.Contract.Candidates(&_Voting.CallOpts, arg0)
}

// HasVoted is a free data retrieval call binding the contract method 0x09eef43e.
//
// Solidity: function hasVoted(address ) view returns(bool)
func (_Voting *VotingCaller) HasVoted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "hasVoted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x09eef43e.
//
// Solidity: function hasVoted(address ) view returns(bool)
func (_Voting *VotingSession) HasVoted(arg0 common.Address) (bool, error) {
	return _Voting.Contract.HasVoted(&_Voting.CallOpts, arg0)
}

// HasVoted is a free data retrieval call binding the contract method 0x09eef43e.
//
// Solidity: function hasVoted(address ) view returns(bool)
func (_Voting *VotingCallerSession) HasVoted(arg0 common.Address) (bool, error) {
	return _Voting.Contract.HasVoted(&_Voting.CallOpts, arg0)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _candidateID) returns()
func (_Voting *VotingTransactor) Vote(opts *bind.TransactOpts, _candidateID *big.Int) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "vote", _candidateID)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _candidateID) returns()
func (_Voting *VotingSession) Vote(_candidateID *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, _candidateID)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _candidateID) returns()
func (_Voting *VotingTransactorSession) Vote(_candidateID *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, _candidateID)
}
