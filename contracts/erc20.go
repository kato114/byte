// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/kato114/byte/blob/main/LICENSE)

package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/kato114/byte/v15/x/evm/types"

	"github.com/kato114/byte/v15/x/erc20/types"
)

var (
	//go:embed compiled_contracts/ERC20MinterBurnerDecimals.json
	ERC20MinterBurnerDecimalsJSON []byte //nolint: golint

	// ERC20MinterBurnerDecimalsContract is the compiled erc20 contract
	ERC20MinterBurnerDecimalsContract evmtypes.CompiledContract

	// ERC20MinterBurnerDecimalsAddress is the erc20 module address
	ERC20MinterBurnerDecimalsAddress common.Address
)

func init() {
	ERC20MinterBurnerDecimalsAddress = types.ModuleAddress

	err := json.Unmarshal(ERC20MinterBurnerDecimalsJSON, &ERC20MinterBurnerDecimalsContract)
	if err != nil {
		panic(err)
	}

	if len(ERC20MinterBurnerDecimalsContract.Bin) == 0 {
		panic("load contract failed")
	}
}
