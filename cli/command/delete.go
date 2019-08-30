package command

import (
  "fmt"
  "os"
  "github.com/satori/go.uuid"
  "github.com/jarmo/secrets/v5/cli/vaultfile"
  "github.com/jarmo/secrets/v5/storage"
  "github.com/jarmo/secrets/v5/vault"
)

type Delete struct {
  Id uuid.UUID
  VaultPath string
  VaultAlias string
}

func (command Delete) Execute() {
  secrets, path, password := vaultfile.Read(vaultfile.Path(command.VaultAlias, command.VaultPath))
  deletedSecret, newSecrets, err := vault.Delete(secrets, command.Id)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  } else {
    storage.Write(path, password, newSecrets)
    fmt.Println("Deleted:", deletedSecret)
  }
}

