// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/kato114/byte/blob/main/LICENSE)

package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/kato114/byte/v15/x/vesting/client/cli"
)

var RegisterClawbackProposalHandler = govclient.NewProposalHandler(cli.NewClawbackProposalCmd)
