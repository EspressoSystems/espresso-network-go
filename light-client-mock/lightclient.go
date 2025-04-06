// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lightclientmock

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// IPlonkVerifierPlonkProof is an auto generated low-level Go binding around an user-defined struct.
type IPlonkVerifierPlonkProof struct {
	Wire0                 BN254G1Point
	Wire1                 BN254G1Point
	Wire2                 BN254G1Point
	Wire3                 BN254G1Point
	Wire4                 BN254G1Point
	ProdPerm              BN254G1Point
	Split0                BN254G1Point
	Split1                BN254G1Point
	Split2                BN254G1Point
	Split3                BN254G1Point
	Split4                BN254G1Point
	Zeta                  BN254G1Point
	ZetaOmega             BN254G1Point
	WireEval0             *big.Int
	WireEval1             *big.Int
	WireEval2             *big.Int
	WireEval3             *big.Int
	WireEval4             *big.Int
	SigmaEval0            *big.Int
	SigmaEval1            *big.Int
	SigmaEval2            *big.Int
	SigmaEval3            *big.Int
	ProdPermZetaOmegaEval *big.Int
}

// LightClientLightClientState is an auto generated low-level Go binding around an user-defined struct.
type LightClientLightClientState struct {
	ViewNum       uint64
	BlockHeight   uint64
	BlockCommRoot *big.Int
}

// LightClientStakeTableState is an auto generated low-level Go binding around an user-defined struct.
type LightClientStakeTableState struct {
	Threshold      *big.Int
	BlsKeyComm     *big.Int
	SchnorrKeyComm *big.Int
	AmountComm     *big.Int
}

// LightClientStateHistoryCommitment is an auto generated low-level Go binding around an user-defined struct.
type LightClientStateHistoryCommitment struct {
	L1BlockHeight        uint64
	L1BlockTimestamp     uint64
	HotShotBlockHeight   uint64
	HotShotBlockCommRoot *big.Int
}

// LightclientmockMetaData contains all meta data concerning the Lightclientmock contract.
var LightclientmockMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"genesis\",\"type\":\"tuple\",\"internalType\":\"structLightClient.LightClientState\",\"components\":[{\"name\":\"viewNum\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockCommRoot\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"}]},{\"name\":\"genesisStakeTableState\",\"type\":\"tuple\",\"internalType\":\"structLightClient.StakeTableState\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blsKeyComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"schnorrKeyComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"amountComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"}]},{\"name\":\"maxHistorySeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentBlockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disablePermissionedProverMode\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizedState\",\"inputs\":[],\"outputs\":[{\"name\":\"viewNum\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockCommRoot\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"genesisStakeTableState\",\"inputs\":[],\"outputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blsKeyComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"schnorrKeyComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"amountComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"genesisState\",\"inputs\":[],\"outputs\":[{\"name\":\"viewNum\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockCommRoot\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHotShotCommitment\",\"inputs\":[{\"name\":\"hotShotBlockHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"hotShotBlockCommRoot\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"hotshotBlockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStateHistoryCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"majorVersion\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"minorVersion\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"patchVersion\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_genesis\",\"type\":\"tuple\",\"internalType\":\"structLightClient.LightClientState\",\"components\":[{\"name\":\"viewNum\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockCommRoot\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"}]},{\"name\":\"_genesisStakeTableState\",\"type\":\"tuple\",\"internalType\":\"structLightClient.StakeTableState\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blsKeyComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"schnorrKeyComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"amountComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"}]},{\"name\":\"_stateHistoryRetentionPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPermissionedProverEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lagOverEscapeHatchThreshold\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"newFinalizedState\",\"inputs\":[{\"name\":\"newState\",\"type\":\"tuple\",\"internalType\":\"structLightClient.LightClientState\",\"components\":[{\"name\":\"viewNum\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockCommRoot\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIPlonkVerifier.PlonkProof\",\"components\":[{\"name\":\"wire0\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"wire1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"wire2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"wire3\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"wire4\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"prodPerm\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"split0\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"split1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"split2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"split3\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"split4\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"zeta\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"zetaOmega\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"wireEval0\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"wireEval1\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"wireEval2\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"wireEval3\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"wireEval4\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"sigmaEval0\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"sigmaEval1\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"sigmaEval2\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"sigmaEval3\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"prodPermZetaOmegaEval\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionedProver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFinalizedState\",\"inputs\":[{\"name\":\"state\",\"type\":\"tuple\",\"internalType\":\"structLightClient.LightClientState\",\"components\":[{\"name\":\"viewNum\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockCommRoot\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setHotShotDownSince\",\"inputs\":[{\"name\":\"l1Height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setHotShotUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPermissionedProver\",\"inputs\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStateHistory\",\"inputs\":[{\"name\":\"_stateHistoryCommitments\",\"type\":\"tuple[]\",\"internalType\":\"structLightClient.StateHistoryCommitment[]\",\"components\":[{\"name\":\"l1BlockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"l1BlockTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"hotShotBlockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"hotShotBlockCommRoot\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setstateHistoryRetentionPeriod\",\"inputs\":[{\"name\":\"historySeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stateHistoryCommitments\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"l1BlockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"l1BlockTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"hotShotBlockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"hotShotBlockCommRoot\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stateHistoryFirstIndex\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stateHistoryRetentionPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewState\",\"inputs\":[{\"name\":\"viewNum\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"blockCommRoot\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"BN254.ScalarField\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PermissionedProverNotRequired\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PermissionedProverRequired\",\"inputs\":[{\"name\":\"permissionedProver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgrade\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientSnapshotHistory\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArgs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidHotShotBlockForCommitmentCheck\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMaxStateHistory\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoChangeRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutdatedState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ProverNotPermissioned\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"WrongStakeTableUsed\",\"inputs\":[]}]",
}

// LightclientmockABI is the input ABI used to generate the binding from.
// Deprecated: Use LightclientmockMetaData.ABI instead.
var LightclientmockABI = LightclientmockMetaData.ABI

// Lightclientmock is an auto generated Go binding around an Ethereum contract.
type Lightclientmock struct {
	LightclientmockCaller     // Read-only binding to the contract
	LightclientmockTransactor // Write-only binding to the contract
	LightclientmockFilterer   // Log filterer for contract events
}

// LightclientmockCaller is an auto generated read-only Go binding around an Ethereum contract.
type LightclientmockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightclientmockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LightclientmockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightclientmockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LightclientmockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightclientmockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LightclientmockSession struct {
	Contract     *Lightclientmock  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LightclientmockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LightclientmockCallerSession struct {
	Contract *LightclientmockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// LightclientmockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LightclientmockTransactorSession struct {
	Contract     *LightclientmockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// LightclientmockRaw is an auto generated low-level Go binding around an Ethereum contract.
type LightclientmockRaw struct {
	Contract *Lightclientmock // Generic contract binding to access the raw methods on
}

// LightclientmockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LightclientmockCallerRaw struct {
	Contract *LightclientmockCaller // Generic read-only contract binding to access the raw methods on
}

// LightclientmockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LightclientmockTransactorRaw struct {
	Contract *LightclientmockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLightclientmock creates a new instance of Lightclientmock, bound to a specific deployed contract.
func NewLightclientmock(address common.Address, backend bind.ContractBackend) (*Lightclientmock, error) {
	contract, err := bindLightclientmock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lightclientmock{LightclientmockCaller: LightclientmockCaller{contract: contract}, LightclientmockTransactor: LightclientmockTransactor{contract: contract}, LightclientmockFilterer: LightclientmockFilterer{contract: contract}}, nil
}

// NewLightclientmockCaller creates a new read-only instance of Lightclientmock, bound to a specific deployed contract.
func NewLightclientmockCaller(address common.Address, caller bind.ContractCaller) (*LightclientmockCaller, error) {
	contract, err := bindLightclientmock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LightclientmockCaller{contract: contract}, nil
}

// NewLightclientmockTransactor creates a new write-only instance of Lightclientmock, bound to a specific deployed contract.
func NewLightclientmockTransactor(address common.Address, transactor bind.ContractTransactor) (*LightclientmockTransactor, error) {
	contract, err := bindLightclientmock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LightclientmockTransactor{contract: contract}, nil
}

// NewLightclientmockFilterer creates a new log filterer instance of Lightclientmock, bound to a specific deployed contract.
func NewLightclientmockFilterer(address common.Address, filterer bind.ContractFilterer) (*LightclientmockFilterer, error) {
	contract, err := bindLightclientmock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LightclientmockFilterer{contract: contract}, nil
}

// bindLightclientmock binds a generic wrapper to an already deployed contract.
func bindLightclientmock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LightclientmockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lightclientmock *LightclientmockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lightclientmock.Contract.LightclientmockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lightclientmock *LightclientmockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lightclientmock.Contract.LightclientmockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lightclientmock *LightclientmockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lightclientmock.Contract.LightclientmockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lightclientmock *LightclientmockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lightclientmock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lightclientmock *LightclientmockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lightclientmock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lightclientmock *LightclientmockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lightclientmock.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Lightclientmock *LightclientmockCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Lightclientmock.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Lightclientmock *LightclientmockSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Lightclientmock.Contract.UPGRADEINTERFACEVERSION(&_Lightclientmock.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Lightclientmock *LightclientmockCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Lightclientmock.Contract.UPGRADEINTERFACEVERSION(&_Lightclientmock.CallOpts)
}

// CurrentBlockNumber is a free data retrieval call binding the contract method 0x378ec23b.
//
// Solidity: function currentBlockNumber() view returns(uint256)
func (_Lightclientmock *LightclientmockCaller) CurrentBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lightclientmock.contract.Call(opts, &out, "currentBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentBlockNumber is a free data retrieval call binding the contract method 0x378ec23b.
//
// Solidity: function currentBlockNumber() view returns(uint256)
func (_Lightclientmock *LightclientmockSession) CurrentBlockNumber() (*big.Int, error) {
	return _Lightclientmock.Contract.CurrentBlockNumber(&_Lightclientmock.CallOpts)
}

// CurrentBlockNumber is a free data retrieval call binding the contract method 0x378ec23b.
//
// Solidity: function currentBlockNumber() view returns(uint256)
func (_Lightclientmock *LightclientmockCallerSession) CurrentBlockNumber() (*big.Int, error) {
	return _Lightclientmock.Contract.CurrentBlockNumber(&_Lightclientmock.CallOpts)
}

// FinalizedState is a free data retrieval call binding the contract method 0x9fdb54a7.
//
// Solidity: function finalizedState() view returns(uint64 viewNum, uint64 blockHeight, uint256 blockCommRoot)
func (_Lightclientmock *LightclientmockCaller) FinalizedState(opts *bind.CallOpts) (struct {
	ViewNum       uint64
	BlockHeight   uint64
	BlockCommRoot *big.Int
}, error) {
	var out []interface{}
	err := _Lightclientmock.contract.Call(opts, &out, "finalizedState")

	outstruct := new(struct {
		ViewNum       uint64
		BlockHeight   uint64
		BlockCommRoot *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ViewNum = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.BlockHeight = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.BlockCommRoot = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FinalizedState is a free data retrieval call binding the contract method 0x9fdb54a7.
//
// Solidity: function finalizedState() view returns(uint64 viewNum, uint64 blockHeight, uint256 blockCommRoot)
func (_Lightclientmock *LightclientmockSession) FinalizedState() (struct {
	ViewNum       uint64
	BlockHeight   uint64
	BlockCommRoot *big.Int
}, error) {
	return _Lightclientmock.Contract.FinalizedState(&_Lightclientmock.CallOpts)
}

// FinalizedState is a free data retrieval call binding the contract method 0x9fdb54a7.
//
// Solidity: function finalizedState() view returns(uint64 viewNum, uint64 blockHeight, uint256 blockCommRoot)
func (_Lightclientmock *LightclientmockCallerSession) FinalizedState() (struct {
	ViewNum       uint64
	BlockHeight   uint64
	BlockCommRoot *big.Int
}, error) {
	return _Lightclientmock.Contract.FinalizedState(&_Lightclientmock.CallOpts)
}

// GenesisStakeTableState is a free data retrieval call binding the contract method 0x426d3194.
//
// Solidity: function genesisStakeTableState() view returns(uint256 threshold, uint256 blsKeyComm, uint256 schnorrKeyComm, uint256 amountComm)
func (_Lightclientmock *LightclientmockCaller) GenesisStakeTableState(opts *bind.CallOpts) (struct {
	Threshold      *big.Int
	BlsKeyComm     *big.Int
	SchnorrKeyComm *big.Int
	AmountComm     *big.Int
}, error) {
	var out []interface{}
	err := _Lightclientmock.contract.Call(opts, &out, "genesisStakeTableState")

	outstruct := new(struct {
		Threshold      *big.Int
		BlsKeyComm     *big.Int
		SchnorrKeyComm *big.Int
		AmountComm     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Threshold = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlsKeyComm = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SchnorrKeyComm = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AmountComm = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GenesisStakeTableState is a free data retrieval call binding the contract method 0x426d3194.
//
// Solidity: function genesisStakeTableState() view returns(uint256 threshold, uint256 blsKeyComm, uint256 schnorrKeyComm, uint256 amountComm)
func (_Lightclientmock *LightclientmockSession) GenesisStakeTableState() (struct {
	Threshold      *big.Int
	BlsKeyComm     *big.Int
	SchnorrKeyComm *big.Int
	AmountComm     *big.Int
}, error) {
	return _Lightclientmock.Contract.GenesisStakeTableState(&_Lightclientmock.CallOpts)
}

// GenesisStakeTableState is a free data retrieval call binding the contract method 0x426d3194.
//
// Solidity: function genesisStakeTableState() view returns(uint256 threshold, uint256 blsKeyComm, uint256 schnorrKeyComm, uint256 amountComm)
func (_Lightclientmock *LightclientmockCallerSession) GenesisStakeTableState() (struct {
	Threshold      *big.Int
	BlsKeyComm     *big.Int
	SchnorrKeyComm *big.Int
	AmountComm     *big.Int
}, error) {
	return _Lightclientmock.Contract.GenesisStakeTableState(&_Lightclientmock.CallOpts)
}

// GenesisState is a free data retrieval call binding the contract method 0xd24d933d.
//
// Solidity: function genesisState() view returns(uint64 viewNum, uint64 blockHeight, uint256 blockCommRoot)
func (_Lightclientmock *LightclientmockCaller) GenesisState(opts *bind.CallOpts) (struct {
	ViewNum       uint64
	BlockHeight   uint64
	BlockCommRoot *big.Int
}, error) {
	var out []interface{}
	err := _Lightclientmock.contract.Call(opts, &out, "genesisState")

	outstruct := new(struct {
		ViewNum       uint64
		BlockHeight   uint64
		BlockCommRoot *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ViewNum = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.BlockHeight = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.BlockCommRoot = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GenesisState is a free data retrieval call binding the contract method 0xd24d933d.
//
// Solidity: function genesisState() view returns(uint64 viewNum, uint64 blockHeight, uint256 blockCommRoot)
func (_Lightclientmock *LightclientmockSession) GenesisState() (struct {
	ViewNum       uint64
	BlockHeight   uint64
	BlockCommRoot *big.Int
}, error) {
	return _Lightclientmock.Contract.GenesisState(&_Lightclientmock.CallOpts)
}

// GenesisState is a free data retrieval call binding the contract method 0xd24d933d.
//
// Solidity: function genesisState() view returns(uint64 viewNum, uint64 blockHeight, uint256 blockCommRoot)
func (_Lightclientmock *LightclientmockCallerSession) GenesisState() (struct {
	ViewNum       uint64
	BlockHeight   uint64
	BlockCommRoot *big.Int
}, error) {
	return _Lightclientmock.Contract.GenesisState(&_Lightclientmock.CallOpts)
}

// GetHotShotCommitment is a free data retrieval call binding the contract method 0x8584d23f.
//
// Solidity: function getHotShotCommitment(uint256 hotShotBlockHeight) view returns(uint256 hotShotBlockCommRoot, uint64 hotshotBlockHeight)
func (_Lightclientmock *LightclientmockCaller) GetHotShotCommitment(opts *bind.CallOpts, hotShotBlockHeight *big.Int) (struct {
	HotShotBlockCommRoot *big.Int
	HotshotBlockHeight   uint64
}, error) {
	var out []interface{}
	err := _Lightclientmock.contract.Call(opts, &out, "getHotShotCommitment", hotShotBlockHeight)

	outstruct := new(struct {
		HotShotBlockCommRoot *big.Int
		HotshotBlockHeight   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.HotShotBlockCommRoot = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.HotshotBlockHeight = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetHotShotCommitment is a free data retrieval call binding the contract method 0x8584d23f.
//
// Solidity: function getHotShotCommitment(uint256 hotShotBlockHeight) view returns(uint256 hotShotBlockCommRoot, uint64 hotshotBlockHeight)
func (_Lightclientmock *LightclientmockSession) GetHotShotCommitment(hotShotBlockHeight *big.Int) (struct {
	HotShotBlockCommRoot *big.Int
	HotshotBlockHeight   uint64
}, error) {
	return _Lightclientmock.Contract.GetHotShotCommitment(&_Lightclientmock.CallOpts, hotShotBlockHeight)
}

// GetHotShotCommitment is a free data retrieval call binding the contract method 0x8584d23f.
//
// Solidity: function getHotShotCommitment(uint256 hotShotBlockHeight) view returns(uint256 hotShotBlockCommRoot, uint64 hotshotBlockHeight)
func (_Lightclientmock *LightclientmockCallerSession) GetHotShotCommitment(hotShotBlockHeight *big.Int) (struct {
	HotShotBlockCommRoot *big.Int
	HotshotBlockHeight   uint64
}, error) {
	return _Lightclientmock.Contract.GetHotShotCommitment(&_Lightclientmock.CallOpts, hotShotBlockHeight)
}

// GetStateHistoryCount is a free data retrieval call binding the contract method 0xf9e50d19.
//
// Solidity: function getStateHistoryCount() view returns(uint256)
func (_Lightclientmock *LightclientmockCaller) GetStateHistoryCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lightclientmock.contract.Call(opts, &out, "getStateHistoryCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStateHistoryCount is a free data retrieval call binding the contract method 0xf9e50d19.
//
// Solidity: function getStateHistoryCount() view returns(uint256)
func (_Lightclientmock *LightclientmockSession) GetStateHistoryCount() (*big.Int, error) {
	return _Lightclientmock.Contract.GetStateHistoryCount(&_Lightclientmock.CallOpts)
}

// GetStateHistoryCount is a free data retrieval call binding the contract method 0xf9e50d19.
//
// Solidity: function getStateHistoryCount() view returns(uint256)
func (_Lightclientmock *LightclientmockCallerSession) GetStateHistoryCount() (*big.Int, error) {
	return _Lightclientmock.Contract.GetStateHistoryCount(&_Lightclientmock.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(uint8 majorVersion, uint8 minorVersion, uint8 patchVersion)
func (_Lightclientmock *LightclientmockCaller) GetVersion(opts *bind.CallOpts) (struct {
	MajorVersion uint8
	MinorVersion uint8
	PatchVersion uint8
}, error) {
	var out []interface{}
	err := _Lightclientmock.contract.Call(opts, &out, "getVersion")

	outstruct := new(struct {
		MajorVersion uint8
		MinorVersion uint8
		PatchVersion uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MajorVersion = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.MinorVersion = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.PatchVersion = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(uint8 majorVersion, uint8 minorVersion, uint8 patchVersion)
func (_Lightclientmock *LightclientmockSession) GetVersion() (struct {
	MajorVersion uint8
	MinorVersion uint8
	PatchVersion uint8
}, error) {
	return _Lightclientmock.Contract.GetVersion(&_Lightclientmock.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(uint8 majorVersion, uint8 minorVersion, uint8 patchVersion)
func (_Lightclientmock *LightclientmockCallerSession) GetVersion() (struct {
	MajorVersion uint8
	MinorVersion uint8
	PatchVersion uint8
}, error) {
	return _Lightclientmock.Contract.GetVersion(&_Lightclientmock.CallOpts)
}

// IsPermissionedProverEnabled is a free data retrieval call binding the contract method 0x826e41fc.
//
// Solidity: function isPermissionedProverEnabled() view returns(bool)
func (_Lightclientmock *LightclientmockCaller) IsPermissionedProverEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lightclientmock.contract.Call(opts, &out, "isPermissionedProverEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPermissionedProverEnabled is a free data retrieval call binding the contract method 0x826e41fc.
//
// Solidity: function isPermissionedProverEnabled() view returns(bool)
func (_Lightclientmock *LightclientmockSession) IsPermissionedProverEnabled() (bool, error) {
	return _Lightclientmock.Contract.IsPermissionedProverEnabled(&_Lightclientmock.CallOpts)
}

// IsPermissionedProverEnabled is a free data retrieval call binding the contract method 0x826e41fc.
//
// Solidity: function isPermissionedProverEnabled() view returns(bool)
func (_Lightclientmock *LightclientmockCallerSession) IsPermissionedProverEnabled() (bool, error) {
	return _Lightclientmock.Contract.IsPermissionedProverEnabled(&_Lightclientmock.CallOpts)
}

// LagOverEscapeHatchThreshold is a free data retrieval call binding the contract method 0xe0303301.
//
// Solidity: function lagOverEscapeHatchThreshold(uint256 blockNumber, uint256 threshold) view returns(bool)
func (_Lightclientmock *LightclientmockCaller) LagOverEscapeHatchThreshold(opts *bind.CallOpts, blockNumber *big.Int, threshold *big.Int) (bool, error) {
	var out []interface{}
	err := _Lightclientmock.contract.Call(opts, &out, "lagOverEscapeHatchThreshold", blockNumber, threshold)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LagOverEscapeHatchThreshold is a free data retrieval call binding the contract method 0xe0303301.
//
// Solidity: function lagOverEscapeHatchThreshold(uint256 blockNumber, uint256 threshold) view returns(bool)
func (_Lightclientmock *LightclientmockSession) LagOverEscapeHatchThreshold(blockNumber *big.Int, threshold *big.Int) (bool, error) {
	return _Lightclientmock.Contract.LagOverEscapeHatchThreshold(&_Lightclientmock.CallOpts, blockNumber, threshold)
}

// LagOverEscapeHatchThreshold is a free data retrieval call binding the contract method 0xe0303301.
//
// Solidity: function lagOverEscapeHatchThreshold(uint256 blockNumber, uint256 threshold) view returns(bool)
func (_Lightclientmock *LightclientmockCallerSession) LagOverEscapeHatchThreshold(blockNumber *big.Int, threshold *big.Int) (bool, error) {
	return _Lightclientmock.Contract.LagOverEscapeHatchThreshold(&_Lightclientmock.CallOpts, blockNumber, threshold)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lightclientmock *LightclientmockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lightclientmock.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lightclientmock *LightclientmockSession) Owner() (common.Address, error) {
	return _Lightclientmock.Contract.Owner(&_Lightclientmock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lightclientmock *LightclientmockCallerSession) Owner() (common.Address, error) {
	return _Lightclientmock.Contract.Owner(&_Lightclientmock.CallOpts)
}

// PermissionedProver is a free data retrieval call binding the contract method 0x313df7b1.
//
// Solidity: function permissionedProver() view returns(address)
func (_Lightclientmock *LightclientmockCaller) PermissionedProver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lightclientmock.contract.Call(opts, &out, "permissionedProver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionedProver is a free data retrieval call binding the contract method 0x313df7b1.
//
// Solidity: function permissionedProver() view returns(address)
func (_Lightclientmock *LightclientmockSession) PermissionedProver() (common.Address, error) {
	return _Lightclientmock.Contract.PermissionedProver(&_Lightclientmock.CallOpts)
}

// PermissionedProver is a free data retrieval call binding the contract method 0x313df7b1.
//
// Solidity: function permissionedProver() view returns(address)
func (_Lightclientmock *LightclientmockCallerSession) PermissionedProver() (common.Address, error) {
	return _Lightclientmock.Contract.PermissionedProver(&_Lightclientmock.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Lightclientmock *LightclientmockCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lightclientmock.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Lightclientmock *LightclientmockSession) ProxiableUUID() ([32]byte, error) {
	return _Lightclientmock.Contract.ProxiableUUID(&_Lightclientmock.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Lightclientmock *LightclientmockCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Lightclientmock.Contract.ProxiableUUID(&_Lightclientmock.CallOpts)
}

// StateHistoryCommitments is a free data retrieval call binding the contract method 0x02b592f3.
//
// Solidity: function stateHistoryCommitments(uint256 ) view returns(uint64 l1BlockHeight, uint64 l1BlockTimestamp, uint64 hotShotBlockHeight, uint256 hotShotBlockCommRoot)
func (_Lightclientmock *LightclientmockCaller) StateHistoryCommitments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	L1BlockHeight        uint64
	L1BlockTimestamp     uint64
	HotShotBlockHeight   uint64
	HotShotBlockCommRoot *big.Int
}, error) {
	var out []interface{}
	err := _Lightclientmock.contract.Call(opts, &out, "stateHistoryCommitments", arg0)

	outstruct := new(struct {
		L1BlockHeight        uint64
		L1BlockTimestamp     uint64
		HotShotBlockHeight   uint64
		HotShotBlockCommRoot *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.L1BlockHeight = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.L1BlockTimestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.HotShotBlockHeight = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.HotShotBlockCommRoot = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StateHistoryCommitments is a free data retrieval call binding the contract method 0x02b592f3.
//
// Solidity: function stateHistoryCommitments(uint256 ) view returns(uint64 l1BlockHeight, uint64 l1BlockTimestamp, uint64 hotShotBlockHeight, uint256 hotShotBlockCommRoot)
func (_Lightclientmock *LightclientmockSession) StateHistoryCommitments(arg0 *big.Int) (struct {
	L1BlockHeight        uint64
	L1BlockTimestamp     uint64
	HotShotBlockHeight   uint64
	HotShotBlockCommRoot *big.Int
}, error) {
	return _Lightclientmock.Contract.StateHistoryCommitments(&_Lightclientmock.CallOpts, arg0)
}

// StateHistoryCommitments is a free data retrieval call binding the contract method 0x02b592f3.
//
// Solidity: function stateHistoryCommitments(uint256 ) view returns(uint64 l1BlockHeight, uint64 l1BlockTimestamp, uint64 hotShotBlockHeight, uint256 hotShotBlockCommRoot)
func (_Lightclientmock *LightclientmockCallerSession) StateHistoryCommitments(arg0 *big.Int) (struct {
	L1BlockHeight        uint64
	L1BlockTimestamp     uint64
	HotShotBlockHeight   uint64
	HotShotBlockCommRoot *big.Int
}, error) {
	return _Lightclientmock.Contract.StateHistoryCommitments(&_Lightclientmock.CallOpts, arg0)
}

// StateHistoryFirstIndex is a free data retrieval call binding the contract method 0x2f79889d.
//
// Solidity: function stateHistoryFirstIndex() view returns(uint64)
func (_Lightclientmock *LightclientmockCaller) StateHistoryFirstIndex(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Lightclientmock.contract.Call(opts, &out, "stateHistoryFirstIndex")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// StateHistoryFirstIndex is a free data retrieval call binding the contract method 0x2f79889d.
//
// Solidity: function stateHistoryFirstIndex() view returns(uint64)
func (_Lightclientmock *LightclientmockSession) StateHistoryFirstIndex() (uint64, error) {
	return _Lightclientmock.Contract.StateHistoryFirstIndex(&_Lightclientmock.CallOpts)
}

// StateHistoryFirstIndex is a free data retrieval call binding the contract method 0x2f79889d.
//
// Solidity: function stateHistoryFirstIndex() view returns(uint64)
func (_Lightclientmock *LightclientmockCallerSession) StateHistoryFirstIndex() (uint64, error) {
	return _Lightclientmock.Contract.StateHistoryFirstIndex(&_Lightclientmock.CallOpts)
}

// StateHistoryRetentionPeriod is a free data retrieval call binding the contract method 0xc23b9e9e.
//
// Solidity: function stateHistoryRetentionPeriod() view returns(uint32)
func (_Lightclientmock *LightclientmockCaller) StateHistoryRetentionPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Lightclientmock.contract.Call(opts, &out, "stateHistoryRetentionPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// StateHistoryRetentionPeriod is a free data retrieval call binding the contract method 0xc23b9e9e.
//
// Solidity: function stateHistoryRetentionPeriod() view returns(uint32)
func (_Lightclientmock *LightclientmockSession) StateHistoryRetentionPeriod() (uint32, error) {
	return _Lightclientmock.Contract.StateHistoryRetentionPeriod(&_Lightclientmock.CallOpts)
}

// StateHistoryRetentionPeriod is a free data retrieval call binding the contract method 0xc23b9e9e.
//
// Solidity: function stateHistoryRetentionPeriod() view returns(uint32)
func (_Lightclientmock *LightclientmockCallerSession) StateHistoryRetentionPeriod() (uint32, error) {
	return _Lightclientmock.Contract.StateHistoryRetentionPeriod(&_Lightclientmock.CallOpts)
}

// DisablePermissionedProverMode is a paid mutator transaction binding the contract method 0x69cc6a04.
//
// Solidity: function disablePermissionedProverMode() returns()
func (_Lightclientmock *LightclientmockTransactor) DisablePermissionedProverMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lightclientmock.contract.Transact(opts, "disablePermissionedProverMode")
}

// DisablePermissionedProverMode is a paid mutator transaction binding the contract method 0x69cc6a04.
//
// Solidity: function disablePermissionedProverMode() returns()
func (_Lightclientmock *LightclientmockSession) DisablePermissionedProverMode() (*types.Transaction, error) {
	return _Lightclientmock.Contract.DisablePermissionedProverMode(&_Lightclientmock.TransactOpts)
}

// DisablePermissionedProverMode is a paid mutator transaction binding the contract method 0x69cc6a04.
//
// Solidity: function disablePermissionedProverMode() returns()
func (_Lightclientmock *LightclientmockTransactorSession) DisablePermissionedProverMode() (*types.Transaction, error) {
	return _Lightclientmock.Contract.DisablePermissionedProverMode(&_Lightclientmock.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x9baa3cc9.
//
// Solidity: function initialize((uint64,uint64,uint256) _genesis, (uint256,uint256,uint256,uint256) _genesisStakeTableState, uint32 _stateHistoryRetentionPeriod, address owner) returns()
func (_Lightclientmock *LightclientmockTransactor) Initialize(opts *bind.TransactOpts, _genesis LightClientLightClientState, _genesisStakeTableState LightClientStakeTableState, _stateHistoryRetentionPeriod uint32, owner common.Address) (*types.Transaction, error) {
	return _Lightclientmock.contract.Transact(opts, "initialize", _genesis, _genesisStakeTableState, _stateHistoryRetentionPeriod, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x9baa3cc9.
//
// Solidity: function initialize((uint64,uint64,uint256) _genesis, (uint256,uint256,uint256,uint256) _genesisStakeTableState, uint32 _stateHistoryRetentionPeriod, address owner) returns()
func (_Lightclientmock *LightclientmockSession) Initialize(_genesis LightClientLightClientState, _genesisStakeTableState LightClientStakeTableState, _stateHistoryRetentionPeriod uint32, owner common.Address) (*types.Transaction, error) {
	return _Lightclientmock.Contract.Initialize(&_Lightclientmock.TransactOpts, _genesis, _genesisStakeTableState, _stateHistoryRetentionPeriod, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x9baa3cc9.
//
// Solidity: function initialize((uint64,uint64,uint256) _genesis, (uint256,uint256,uint256,uint256) _genesisStakeTableState, uint32 _stateHistoryRetentionPeriod, address owner) returns()
func (_Lightclientmock *LightclientmockTransactorSession) Initialize(_genesis LightClientLightClientState, _genesisStakeTableState LightClientStakeTableState, _stateHistoryRetentionPeriod uint32, owner common.Address) (*types.Transaction, error) {
	return _Lightclientmock.Contract.Initialize(&_Lightclientmock.TransactOpts, _genesis, _genesisStakeTableState, _stateHistoryRetentionPeriod, owner)
}

// NewFinalizedState is a paid mutator transaction binding the contract method 0x2063d4f7.
//
// Solidity: function newFinalizedState((uint64,uint64,uint256) newState, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) proof) returns()
func (_Lightclientmock *LightclientmockTransactor) NewFinalizedState(opts *bind.TransactOpts, newState LightClientLightClientState, proof IPlonkVerifierPlonkProof) (*types.Transaction, error) {
	return _Lightclientmock.contract.Transact(opts, "newFinalizedState", newState, proof)
}

// NewFinalizedState is a paid mutator transaction binding the contract method 0x2063d4f7.
//
// Solidity: function newFinalizedState((uint64,uint64,uint256) newState, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) proof) returns()
func (_Lightclientmock *LightclientmockSession) NewFinalizedState(newState LightClientLightClientState, proof IPlonkVerifierPlonkProof) (*types.Transaction, error) {
	return _Lightclientmock.Contract.NewFinalizedState(&_Lightclientmock.TransactOpts, newState, proof)
}

// NewFinalizedState is a paid mutator transaction binding the contract method 0x2063d4f7.
//
// Solidity: function newFinalizedState((uint64,uint64,uint256) newState, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) proof) returns()
func (_Lightclientmock *LightclientmockTransactorSession) NewFinalizedState(newState LightClientLightClientState, proof IPlonkVerifierPlonkProof) (*types.Transaction, error) {
	return _Lightclientmock.Contract.NewFinalizedState(&_Lightclientmock.TransactOpts, newState, proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lightclientmock *LightclientmockTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lightclientmock.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lightclientmock *LightclientmockSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lightclientmock.Contract.RenounceOwnership(&_Lightclientmock.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lightclientmock *LightclientmockTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lightclientmock.Contract.RenounceOwnership(&_Lightclientmock.TransactOpts)
}

// SetFinalizedState is a paid mutator transaction binding the contract method 0xb5adea3c.
//
// Solidity: function setFinalizedState((uint64,uint64,uint256) state) returns()
func (_Lightclientmock *LightclientmockTransactor) SetFinalizedState(opts *bind.TransactOpts, state LightClientLightClientState) (*types.Transaction, error) {
	return _Lightclientmock.contract.Transact(opts, "setFinalizedState", state)
}

// SetFinalizedState is a paid mutator transaction binding the contract method 0xb5adea3c.
//
// Solidity: function setFinalizedState((uint64,uint64,uint256) state) returns()
func (_Lightclientmock *LightclientmockSession) SetFinalizedState(state LightClientLightClientState) (*types.Transaction, error) {
	return _Lightclientmock.Contract.SetFinalizedState(&_Lightclientmock.TransactOpts, state)
}

// SetFinalizedState is a paid mutator transaction binding the contract method 0xb5adea3c.
//
// Solidity: function setFinalizedState((uint64,uint64,uint256) state) returns()
func (_Lightclientmock *LightclientmockTransactorSession) SetFinalizedState(state LightClientLightClientState) (*types.Transaction, error) {
	return _Lightclientmock.Contract.SetFinalizedState(&_Lightclientmock.TransactOpts, state)
}

// SetHotShotDownSince is a paid mutator transaction binding the contract method 0x2d52aad6.
//
// Solidity: function setHotShotDownSince(uint256 l1Height) returns()
func (_Lightclientmock *LightclientmockTransactor) SetHotShotDownSince(opts *bind.TransactOpts, l1Height *big.Int) (*types.Transaction, error) {
	return _Lightclientmock.contract.Transact(opts, "setHotShotDownSince", l1Height)
}

// SetHotShotDownSince is a paid mutator transaction binding the contract method 0x2d52aad6.
//
// Solidity: function setHotShotDownSince(uint256 l1Height) returns()
func (_Lightclientmock *LightclientmockSession) SetHotShotDownSince(l1Height *big.Int) (*types.Transaction, error) {
	return _Lightclientmock.Contract.SetHotShotDownSince(&_Lightclientmock.TransactOpts, l1Height)
}

// SetHotShotDownSince is a paid mutator transaction binding the contract method 0x2d52aad6.
//
// Solidity: function setHotShotDownSince(uint256 l1Height) returns()
func (_Lightclientmock *LightclientmockTransactorSession) SetHotShotDownSince(l1Height *big.Int) (*types.Transaction, error) {
	return _Lightclientmock.Contract.SetHotShotDownSince(&_Lightclientmock.TransactOpts, l1Height)
}

// SetHotShotUp is a paid mutator transaction binding the contract method 0xc8e5e498.
//
// Solidity: function setHotShotUp() returns()
func (_Lightclientmock *LightclientmockTransactor) SetHotShotUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lightclientmock.contract.Transact(opts, "setHotShotUp")
}

// SetHotShotUp is a paid mutator transaction binding the contract method 0xc8e5e498.
//
// Solidity: function setHotShotUp() returns()
func (_Lightclientmock *LightclientmockSession) SetHotShotUp() (*types.Transaction, error) {
	return _Lightclientmock.Contract.SetHotShotUp(&_Lightclientmock.TransactOpts)
}

// SetHotShotUp is a paid mutator transaction binding the contract method 0xc8e5e498.
//
// Solidity: function setHotShotUp() returns()
func (_Lightclientmock *LightclientmockTransactorSession) SetHotShotUp() (*types.Transaction, error) {
	return _Lightclientmock.Contract.SetHotShotUp(&_Lightclientmock.TransactOpts)
}

// SetPermissionedProver is a paid mutator transaction binding the contract method 0x013fa5fc.
//
// Solidity: function setPermissionedProver(address prover) returns()
func (_Lightclientmock *LightclientmockTransactor) SetPermissionedProver(opts *bind.TransactOpts, prover common.Address) (*types.Transaction, error) {
	return _Lightclientmock.contract.Transact(opts, "setPermissionedProver", prover)
}

// SetPermissionedProver is a paid mutator transaction binding the contract method 0x013fa5fc.
//
// Solidity: function setPermissionedProver(address prover) returns()
func (_Lightclientmock *LightclientmockSession) SetPermissionedProver(prover common.Address) (*types.Transaction, error) {
	return _Lightclientmock.Contract.SetPermissionedProver(&_Lightclientmock.TransactOpts, prover)
}

// SetPermissionedProver is a paid mutator transaction binding the contract method 0x013fa5fc.
//
// Solidity: function setPermissionedProver(address prover) returns()
func (_Lightclientmock *LightclientmockTransactorSession) SetPermissionedProver(prover common.Address) (*types.Transaction, error) {
	return _Lightclientmock.Contract.SetPermissionedProver(&_Lightclientmock.TransactOpts, prover)
}

// SetStateHistory is a paid mutator transaction binding the contract method 0xf5676160.
//
// Solidity: function setStateHistory((uint64,uint64,uint64,uint256)[] _stateHistoryCommitments) returns()
func (_Lightclientmock *LightclientmockTransactor) SetStateHistory(opts *bind.TransactOpts, _stateHistoryCommitments []LightClientStateHistoryCommitment) (*types.Transaction, error) {
	return _Lightclientmock.contract.Transact(opts, "setStateHistory", _stateHistoryCommitments)
}

// SetStateHistory is a paid mutator transaction binding the contract method 0xf5676160.
//
// Solidity: function setStateHistory((uint64,uint64,uint64,uint256)[] _stateHistoryCommitments) returns()
func (_Lightclientmock *LightclientmockSession) SetStateHistory(_stateHistoryCommitments []LightClientStateHistoryCommitment) (*types.Transaction, error) {
	return _Lightclientmock.Contract.SetStateHistory(&_Lightclientmock.TransactOpts, _stateHistoryCommitments)
}

// SetStateHistory is a paid mutator transaction binding the contract method 0xf5676160.
//
// Solidity: function setStateHistory((uint64,uint64,uint64,uint256)[] _stateHistoryCommitments) returns()
func (_Lightclientmock *LightclientmockTransactorSession) SetStateHistory(_stateHistoryCommitments []LightClientStateHistoryCommitment) (*types.Transaction, error) {
	return _Lightclientmock.Contract.SetStateHistory(&_Lightclientmock.TransactOpts, _stateHistoryCommitments)
}

// SetstateHistoryRetentionPeriod is a paid mutator transaction binding the contract method 0x96c1ca61.
//
// Solidity: function setstateHistoryRetentionPeriod(uint32 historySeconds) returns()
func (_Lightclientmock *LightclientmockTransactor) SetstateHistoryRetentionPeriod(opts *bind.TransactOpts, historySeconds uint32) (*types.Transaction, error) {
	return _Lightclientmock.contract.Transact(opts, "setstateHistoryRetentionPeriod", historySeconds)
}

// SetstateHistoryRetentionPeriod is a paid mutator transaction binding the contract method 0x96c1ca61.
//
// Solidity: function setstateHistoryRetentionPeriod(uint32 historySeconds) returns()
func (_Lightclientmock *LightclientmockSession) SetstateHistoryRetentionPeriod(historySeconds uint32) (*types.Transaction, error) {
	return _Lightclientmock.Contract.SetstateHistoryRetentionPeriod(&_Lightclientmock.TransactOpts, historySeconds)
}

// SetstateHistoryRetentionPeriod is a paid mutator transaction binding the contract method 0x96c1ca61.
//
// Solidity: function setstateHistoryRetentionPeriod(uint32 historySeconds) returns()
func (_Lightclientmock *LightclientmockTransactorSession) SetstateHistoryRetentionPeriod(historySeconds uint32) (*types.Transaction, error) {
	return _Lightclientmock.Contract.SetstateHistoryRetentionPeriod(&_Lightclientmock.TransactOpts, historySeconds)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lightclientmock *LightclientmockTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lightclientmock.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lightclientmock *LightclientmockSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lightclientmock.Contract.TransferOwnership(&_Lightclientmock.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lightclientmock *LightclientmockTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lightclientmock.Contract.TransferOwnership(&_Lightclientmock.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Lightclientmock *LightclientmockTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Lightclientmock.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Lightclientmock *LightclientmockSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Lightclientmock.Contract.UpgradeToAndCall(&_Lightclientmock.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Lightclientmock *LightclientmockTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Lightclientmock.Contract.UpgradeToAndCall(&_Lightclientmock.TransactOpts, newImplementation, data)
}

// LightclientmockInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Lightclientmock contract.
type LightclientmockInitializedIterator struct {
	Event *LightclientmockInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightclientmockInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightclientmockInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightclientmockInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightclientmockInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightclientmockInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightclientmockInitialized represents a Initialized event raised by the Lightclientmock contract.
type LightclientmockInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Lightclientmock *LightclientmockFilterer) FilterInitialized(opts *bind.FilterOpts) (*LightclientmockInitializedIterator, error) {

	logs, sub, err := _Lightclientmock.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LightclientmockInitializedIterator{contract: _Lightclientmock.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Lightclientmock *LightclientmockFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LightclientmockInitialized) (event.Subscription, error) {

	logs, sub, err := _Lightclientmock.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightclientmockInitialized)
				if err := _Lightclientmock.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Lightclientmock *LightclientmockFilterer) ParseInitialized(log types.Log) (*LightclientmockInitialized, error) {
	event := new(LightclientmockInitialized)
	if err := _Lightclientmock.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightclientmockNewStateIterator is returned from FilterNewState and is used to iterate over the raw logs and unpacked data for NewState events raised by the Lightclientmock contract.
type LightclientmockNewStateIterator struct {
	Event *LightclientmockNewState // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightclientmockNewStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightclientmockNewState)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightclientmockNewState)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightclientmockNewStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightclientmockNewStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightclientmockNewState represents a NewState event raised by the Lightclientmock contract.
type LightclientmockNewState struct {
	ViewNum       uint64
	BlockHeight   uint64
	BlockCommRoot *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewState is a free log retrieval operation binding the contract event 0xa04a773924505a418564363725f56832f5772e6b8d0dbd6efce724dfe803dae6.
//
// Solidity: event NewState(uint64 indexed viewNum, uint64 indexed blockHeight, uint256 blockCommRoot)
func (_Lightclientmock *LightclientmockFilterer) FilterNewState(opts *bind.FilterOpts, viewNum []uint64, blockHeight []uint64) (*LightclientmockNewStateIterator, error) {

	var viewNumRule []interface{}
	for _, viewNumItem := range viewNum {
		viewNumRule = append(viewNumRule, viewNumItem)
	}
	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	logs, sub, err := _Lightclientmock.contract.FilterLogs(opts, "NewState", viewNumRule, blockHeightRule)
	if err != nil {
		return nil, err
	}
	return &LightclientmockNewStateIterator{contract: _Lightclientmock.contract, event: "NewState", logs: logs, sub: sub}, nil
}

// WatchNewState is a free log subscription operation binding the contract event 0xa04a773924505a418564363725f56832f5772e6b8d0dbd6efce724dfe803dae6.
//
// Solidity: event NewState(uint64 indexed viewNum, uint64 indexed blockHeight, uint256 blockCommRoot)
func (_Lightclientmock *LightclientmockFilterer) WatchNewState(opts *bind.WatchOpts, sink chan<- *LightclientmockNewState, viewNum []uint64, blockHeight []uint64) (event.Subscription, error) {

	var viewNumRule []interface{}
	for _, viewNumItem := range viewNum {
		viewNumRule = append(viewNumRule, viewNumItem)
	}
	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	logs, sub, err := _Lightclientmock.contract.WatchLogs(opts, "NewState", viewNumRule, blockHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightclientmockNewState)
				if err := _Lightclientmock.contract.UnpackLog(event, "NewState", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewState is a log parse operation binding the contract event 0xa04a773924505a418564363725f56832f5772e6b8d0dbd6efce724dfe803dae6.
//
// Solidity: event NewState(uint64 indexed viewNum, uint64 indexed blockHeight, uint256 blockCommRoot)
func (_Lightclientmock *LightclientmockFilterer) ParseNewState(log types.Log) (*LightclientmockNewState, error) {
	event := new(LightclientmockNewState)
	if err := _Lightclientmock.contract.UnpackLog(event, "NewState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightclientmockOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Lightclientmock contract.
type LightclientmockOwnershipTransferredIterator struct {
	Event *LightclientmockOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightclientmockOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightclientmockOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightclientmockOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightclientmockOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightclientmockOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightclientmockOwnershipTransferred represents a OwnershipTransferred event raised by the Lightclientmock contract.
type LightclientmockOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lightclientmock *LightclientmockFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LightclientmockOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lightclientmock.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LightclientmockOwnershipTransferredIterator{contract: _Lightclientmock.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lightclientmock *LightclientmockFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LightclientmockOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lightclientmock.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightclientmockOwnershipTransferred)
				if err := _Lightclientmock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lightclientmock *LightclientmockFilterer) ParseOwnershipTransferred(log types.Log) (*LightclientmockOwnershipTransferred, error) {
	event := new(LightclientmockOwnershipTransferred)
	if err := _Lightclientmock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightclientmockPermissionedProverNotRequiredIterator is returned from FilterPermissionedProverNotRequired and is used to iterate over the raw logs and unpacked data for PermissionedProverNotRequired events raised by the Lightclientmock contract.
type LightclientmockPermissionedProverNotRequiredIterator struct {
	Event *LightclientmockPermissionedProverNotRequired // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightclientmockPermissionedProverNotRequiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightclientmockPermissionedProverNotRequired)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightclientmockPermissionedProverNotRequired)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightclientmockPermissionedProverNotRequiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightclientmockPermissionedProverNotRequiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightclientmockPermissionedProverNotRequired represents a PermissionedProverNotRequired event raised by the Lightclientmock contract.
type LightclientmockPermissionedProverNotRequired struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPermissionedProverNotRequired is a free log retrieval operation binding the contract event 0x9a5f57de856dd668c54dd95e5c55df93432171cbca49a8776d5620ea59c02450.
//
// Solidity: event PermissionedProverNotRequired()
func (_Lightclientmock *LightclientmockFilterer) FilterPermissionedProverNotRequired(opts *bind.FilterOpts) (*LightclientmockPermissionedProverNotRequiredIterator, error) {

	logs, sub, err := _Lightclientmock.contract.FilterLogs(opts, "PermissionedProverNotRequired")
	if err != nil {
		return nil, err
	}
	return &LightclientmockPermissionedProverNotRequiredIterator{contract: _Lightclientmock.contract, event: "PermissionedProverNotRequired", logs: logs, sub: sub}, nil
}

// WatchPermissionedProverNotRequired is a free log subscription operation binding the contract event 0x9a5f57de856dd668c54dd95e5c55df93432171cbca49a8776d5620ea59c02450.
//
// Solidity: event PermissionedProverNotRequired()
func (_Lightclientmock *LightclientmockFilterer) WatchPermissionedProverNotRequired(opts *bind.WatchOpts, sink chan<- *LightclientmockPermissionedProverNotRequired) (event.Subscription, error) {

	logs, sub, err := _Lightclientmock.contract.WatchLogs(opts, "PermissionedProverNotRequired")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightclientmockPermissionedProverNotRequired)
				if err := _Lightclientmock.contract.UnpackLog(event, "PermissionedProverNotRequired", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePermissionedProverNotRequired is a log parse operation binding the contract event 0x9a5f57de856dd668c54dd95e5c55df93432171cbca49a8776d5620ea59c02450.
//
// Solidity: event PermissionedProverNotRequired()
func (_Lightclientmock *LightclientmockFilterer) ParsePermissionedProverNotRequired(log types.Log) (*LightclientmockPermissionedProverNotRequired, error) {
	event := new(LightclientmockPermissionedProverNotRequired)
	if err := _Lightclientmock.contract.UnpackLog(event, "PermissionedProverNotRequired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightclientmockPermissionedProverRequiredIterator is returned from FilterPermissionedProverRequired and is used to iterate over the raw logs and unpacked data for PermissionedProverRequired events raised by the Lightclientmock contract.
type LightclientmockPermissionedProverRequiredIterator struct {
	Event *LightclientmockPermissionedProverRequired // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightclientmockPermissionedProverRequiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightclientmockPermissionedProverRequired)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightclientmockPermissionedProverRequired)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightclientmockPermissionedProverRequiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightclientmockPermissionedProverRequiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightclientmockPermissionedProverRequired represents a PermissionedProverRequired event raised by the Lightclientmock contract.
type LightclientmockPermissionedProverRequired struct {
	PermissionedProver common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPermissionedProverRequired is a free log retrieval operation binding the contract event 0x8017bb887fdf8fca4314a9d40f6e73b3b81002d67e5cfa85d88173af6aa46072.
//
// Solidity: event PermissionedProverRequired(address permissionedProver)
func (_Lightclientmock *LightclientmockFilterer) FilterPermissionedProverRequired(opts *bind.FilterOpts) (*LightclientmockPermissionedProverRequiredIterator, error) {

	logs, sub, err := _Lightclientmock.contract.FilterLogs(opts, "PermissionedProverRequired")
	if err != nil {
		return nil, err
	}
	return &LightclientmockPermissionedProverRequiredIterator{contract: _Lightclientmock.contract, event: "PermissionedProverRequired", logs: logs, sub: sub}, nil
}

// WatchPermissionedProverRequired is a free log subscription operation binding the contract event 0x8017bb887fdf8fca4314a9d40f6e73b3b81002d67e5cfa85d88173af6aa46072.
//
// Solidity: event PermissionedProverRequired(address permissionedProver)
func (_Lightclientmock *LightclientmockFilterer) WatchPermissionedProverRequired(opts *bind.WatchOpts, sink chan<- *LightclientmockPermissionedProverRequired) (event.Subscription, error) {

	logs, sub, err := _Lightclientmock.contract.WatchLogs(opts, "PermissionedProverRequired")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightclientmockPermissionedProverRequired)
				if err := _Lightclientmock.contract.UnpackLog(event, "PermissionedProverRequired", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePermissionedProverRequired is a log parse operation binding the contract event 0x8017bb887fdf8fca4314a9d40f6e73b3b81002d67e5cfa85d88173af6aa46072.
//
// Solidity: event PermissionedProverRequired(address permissionedProver)
func (_Lightclientmock *LightclientmockFilterer) ParsePermissionedProverRequired(log types.Log) (*LightclientmockPermissionedProverRequired, error) {
	event := new(LightclientmockPermissionedProverRequired)
	if err := _Lightclientmock.contract.UnpackLog(event, "PermissionedProverRequired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightclientmockUpgradeIterator is returned from FilterUpgrade and is used to iterate over the raw logs and unpacked data for Upgrade events raised by the Lightclientmock contract.
type LightclientmockUpgradeIterator struct {
	Event *LightclientmockUpgrade // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightclientmockUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightclientmockUpgrade)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightclientmockUpgrade)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightclientmockUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightclientmockUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightclientmockUpgrade represents a Upgrade event raised by the Lightclientmock contract.
type LightclientmockUpgrade struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgrade is a free log retrieval operation binding the contract event 0xf78721226efe9a1bb678189a16d1554928b9f2192e2cb93eeda83b79fa40007d.
//
// Solidity: event Upgrade(address implementation)
func (_Lightclientmock *LightclientmockFilterer) FilterUpgrade(opts *bind.FilterOpts) (*LightclientmockUpgradeIterator, error) {

	logs, sub, err := _Lightclientmock.contract.FilterLogs(opts, "Upgrade")
	if err != nil {
		return nil, err
	}
	return &LightclientmockUpgradeIterator{contract: _Lightclientmock.contract, event: "Upgrade", logs: logs, sub: sub}, nil
}

// WatchUpgrade is a free log subscription operation binding the contract event 0xf78721226efe9a1bb678189a16d1554928b9f2192e2cb93eeda83b79fa40007d.
//
// Solidity: event Upgrade(address implementation)
func (_Lightclientmock *LightclientmockFilterer) WatchUpgrade(opts *bind.WatchOpts, sink chan<- *LightclientmockUpgrade) (event.Subscription, error) {

	logs, sub, err := _Lightclientmock.contract.WatchLogs(opts, "Upgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightclientmockUpgrade)
				if err := _Lightclientmock.contract.UnpackLog(event, "Upgrade", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgrade is a log parse operation binding the contract event 0xf78721226efe9a1bb678189a16d1554928b9f2192e2cb93eeda83b79fa40007d.
//
// Solidity: event Upgrade(address implementation)
func (_Lightclientmock *LightclientmockFilterer) ParseUpgrade(log types.Log) (*LightclientmockUpgrade, error) {
	event := new(LightclientmockUpgrade)
	if err := _Lightclientmock.contract.UnpackLog(event, "Upgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightclientmockUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Lightclientmock contract.
type LightclientmockUpgradedIterator struct {
	Event *LightclientmockUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightclientmockUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightclientmockUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightclientmockUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightclientmockUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightclientmockUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightclientmockUpgraded represents a Upgraded event raised by the Lightclientmock contract.
type LightclientmockUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Lightclientmock *LightclientmockFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LightclientmockUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Lightclientmock.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LightclientmockUpgradedIterator{contract: _Lightclientmock.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Lightclientmock *LightclientmockFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LightclientmockUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Lightclientmock.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightclientmockUpgraded)
				if err := _Lightclientmock.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Lightclientmock *LightclientmockFilterer) ParseUpgraded(log types.Log) (*LightclientmockUpgraded, error) {
	event := new(LightclientmockUpgraded)
	if err := _Lightclientmock.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
