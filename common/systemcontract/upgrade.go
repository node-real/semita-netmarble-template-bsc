package systemcontract

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

type UpgradeConfig struct {
	BeforeUpgrade upgradeHook
	AfterUpgrade  upgradeHook
	ContractAddr  common.Address
	CommitUrl     string
	Code          string
}

type Upgrade struct {
	UpgradeName string
	Configs     []*UpgradeConfig
}

type upgradeHook func(blockNumber *big.Int, contractAddr common.Address, statedb *state.StateDB) error

const (
	mainNet    = "Mainnet"
	testNet    = "Testnet"
	defaultNet = "Default"
)

var (
	GenesisHash common.Hash
	//upgrade config
	ramanujanUpgrade = make(map[string]*Upgrade)

	nielsUpgrade = make(map[string]*Upgrade)

	mirrorUpgrade = make(map[string]*Upgrade)

	brunoUpgrade = make(map[string]*Upgrade)

	h2Upgrade = make(map[string]*Upgrade)
)

func init() {
	h2Upgrade[defaultNet] = &Upgrade{
		UpgradeName: "h2",
		Configs: []*UpgradeConfig{
			{
				ContractAddr: common.HexToAddress(H2Contract),
				CommitUrl:    "https://github.com/binance-chain/bsc-genesis-contract/commit/f4bc161dac5937b8cbd4fe3089c7514c415430f9 fix: bas",
				Code:         "",
			},
		},
	}
}

func UpgradeBuildInSystemContract(config *params.ChainConfig, blockNumber *big.Int, statedb *state.StateDB) {
	if config == nil || blockNumber == nil || statedb == nil {
		return
	}
	var network string
	switch GenesisHash {
	/* Add mainnet genesis hash */
	case params.MainnetGenesisHash:
		network = mainNet
	case params.TestnetGenesisHash:
		network = testNet
	default:
		network = defaultNet
	}

	logger := log.New("system-contract-upgrade", network)
	if config.IsOnRamanujan(blockNumber) {
		applySystemContractUpgrade(ramanujanUpgrade[network], blockNumber, statedb, logger)
	}

	if config.IsOnNiels(blockNumber) {
		applySystemContractUpgrade(nielsUpgrade[network], blockNumber, statedb, logger)
	}

	if config.IsOnMirrorSync(blockNumber) {
		applySystemContractUpgrade(mirrorUpgrade[network], blockNumber, statedb, logger)
	}

	if config.IsOnBruno(blockNumber) {
		applySystemContractUpgrade(brunoUpgrade[network], blockNumber, statedb, logger)
	}

	if config.IsOnH2(blockNumber) {
		applySystemContractUpgrade(h2Upgrade[network], blockNumber, statedb, logger)
	}

	/*
		apply other upgrades
	*/
}

func applySystemContractUpgrade(upgrade *Upgrade, blockNumber *big.Int, statedb *state.StateDB, logger log.Logger) {
	if upgrade == nil {
		logger.Info("Empty upgrade config", "height", blockNumber.String())
		return
	}

	logger.Info(fmt.Sprintf("Apply upgrade %s at height %d", upgrade.UpgradeName, blockNumber.Int64()))
	for _, cfg := range upgrade.Configs {
		logger.Info(fmt.Sprintf("Upgrade contract %s to commit %s", cfg.ContractAddr.String(), cfg.CommitUrl))

		if cfg.BeforeUpgrade != nil {
			err := cfg.BeforeUpgrade(blockNumber, cfg.ContractAddr, statedb)
			if err != nil {
				panic(fmt.Errorf("contract address: %s, execute beforeUpgrade error: %s", cfg.ContractAddr.String(), err.Error()))
			}
		}

		newContractCode, err := hex.DecodeString(cfg.Code)
		if err != nil {
			panic(fmt.Errorf("failed to decode new contract code: %s", err.Error()))
		}
		statedb.SetCode(cfg.ContractAddr, newContractCode)

		if cfg.AfterUpgrade != nil {
			err := cfg.AfterUpgrade(blockNumber, cfg.ContractAddr, statedb)
			if err != nil {
				panic(fmt.Errorf("contract address: %s, execute afterUpgrade error: %s", cfg.ContractAddr.String(), err.Error()))
			}
		}
	}
}
