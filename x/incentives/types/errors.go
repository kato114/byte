// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/kato114/byte/blob/main/LICENSE)

package types

import (
	errorsmod "cosmossdk.io/errors"
)

// errors
var (
	ErrInternalIncentive = errorsmod.Register(ModuleName, 2, "internal incentives error")
)
