// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/kato114/byte/blob/main/LICENSE)
package upgrade

import (
	"fmt"
)

// CreateModuleQueryExec creates a Evmos module query
func (m *Manager) CreateModuleQueryExec(moduleName, subCommand, chainID string) (string, error) {
	cmd := []string{
		"byted",
		"q",
		moduleName,
		subCommand,
		fmt.Sprintf("--chain-id=%s", chainID),
		"--keyring-backend=test",
		"--log_format=json",
	}
	return m.CreateExec(cmd, m.ContainerID())
}
