// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/kato114/byte/blob/main/LICENSE)
package v14_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	evmosapp "github.com/kato114/byte/v15/app"
	"github.com/kato114/byte/v15/precompiles/vesting"
	"github.com/kato114/byte/v15/x/evm/statedb"
	evmtypes "github.com/kato114/byte/v15/x/evm/types"
	"github.com/stretchr/testify/suite"
)

var s *UpgradesTestSuite

type UpgradesTestSuite struct {
	suite.Suite

	ctx        sdk.Context
	app        *evmosapp.Evmos
	address    common.Address
	validators []stakingtypes.Validator
	ethSigner  ethtypes.Signer
	bondDenom  string

	precompile *vesting.Precompile
	stateDB    *statedb.StateDB

	queryClientEVM evmtypes.QueryClient
}

func TestUpgradeTestSuite(t *testing.T) {
	s = new(UpgradesTestSuite)
	suite.Run(t, s)
}

func (s *UpgradesTestSuite) SetupTest() {
	s.DoSetupTest()
}
