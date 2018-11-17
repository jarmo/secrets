package command

import (
  "fmt"
  "os"
  "github.com/satori/go.uuid"
  "github.com/jarmo/secrets/cli/vaultfile"
  "github.com/jarmo/secrets/storage"
  "github.com/jarmo/secrets/vault"
  "github.com/jarmo/secrets/input"
)

type Edit struct {
  Id uuid.UUID
  VaultPath string
  VaultAlias string
}

func (command Edit) Execute() {
  secrets, path, password := vaultfile.Read(vaultfile.Path(command.VaultAlias, command.VaultPath))
  newName := input.Ask(fmt.Sprintf("Enter new name: "))
  newValue := input.AskMultiline("Enter new value:\n")
  editedSecret, newSecrets, err := vault.Edit(secrets, command.Id, newName, newValue)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  } else {
    storage.Write(path, password, newSecrets)
    fmt.Println("Edited:", editedSecret)
  }
}

