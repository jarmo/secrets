package command

import (
  "fmt"
  "github.com/jarmo/secrets/cli/vaultfile"
  "github.com/jarmo/secrets/storage"
  "github.com/jarmo/secrets/vault"
  "github.com/jarmo/secrets/input"
)

type Add struct {
  Name string
  VaultPath string
  VaultAlias string
}

func (command Add) Execute() {
  secrets, path, password := vaultfile.Read(vaultfile.Path(command.VaultAlias, command.VaultPath))
  secretName := command.Name
  secretValue := input.AskMultiline(fmt.Sprintf("Enter value for '%s':\n", command.Name))
  newSecret, newSecrets := vault.Add(secrets, secretName, secretValue)
  storage.Write(path, password, newSecrets)
  fmt.Println("Added:", newSecret)
}

