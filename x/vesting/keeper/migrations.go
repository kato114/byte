// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/kato114/byte/blob/main/LICENSE)
package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	v2 "github.com/kato114/byte/v15/x/vesting/migrations/v2"
)

var _ module.MigrationHandler = Migrator{}.Migrate1to2

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{
		keeper: keeper,
	}
}

// Migrate1to2 migrates the store from consensus version 1 to 2
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v2.MigrateStore(ctx, m.keeper.accountKeeper)
}
