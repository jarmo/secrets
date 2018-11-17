package command

import (
  "fmt"
  "github.com/jarmo/secrets/vault"
  "github.com/jarmo/secrets/cli/vaultfile"
)

type List struct {
  Filter string
  VaultPath string
  VaultAlias string
}

func (command List) Execute() {
  secrets, _, _ := vaultfile.Read(vaultfile.Path(command.VaultAlias, command.VaultPath))
  for _, secret := range vault.List(secrets, command.Filter) {
    fmt.Println(secret)
  }
}

