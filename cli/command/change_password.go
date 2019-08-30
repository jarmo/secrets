package command

import (
  "fmt"
  "os"
  "github.com/jarmo/secrets/v5/vault"
  "github.com/jarmo/secrets/v5/cli/vaultfile"
  "github.com/jarmo/secrets/v5/input"
)

type ChangePassword struct {
  VaultPath string
  VaultAlias string
}

func (command ChangePassword) Execute() {
  currentPassword := input.AskVaultPassword()
  newPassword := input.AskPassword("Enter new vault password: ")
  newPasswordConfirmation := input.AskPassword("Enter new vault password again: ")

  if err := vault.ChangePassword(vaultfile.Path(command.VaultAlias, command.VaultPath), currentPassword, newPassword, newPasswordConfirmation); err != nil {
    fmt.Println(err)
    os.Exit(1)
  } else {
    fmt.Println("Vault password successfully changed!")
  }
}

