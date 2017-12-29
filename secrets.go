package main

import (
  "os"
  "fmt"
  "github.com/jarmo/secrets/cli"
  "github.com/jarmo/secrets/cli/command"
  "github.com/jarmo/secrets/vault"
  "github.com/jarmo/secrets/storage"
  "github.com/jarmo/secrets/storage/path"
  "github.com/jarmo/secrets/secret"
  "github.com/jarmo/secrets/input"
)

const VERSION = "2.0.0"

func main() {
  switch parsedCommand := cli.Execute(VERSION, os.Args[1:]).(type) {
    case command.List:
      secrets, _, _ := loadVault(vaultPath(parsedCommand.VaultPath))
      for _, secret := range vault.List(secrets, parsedCommand.Filter) {
        fmt.Println(secret)
      }
    case command.Add:
      secrets, path, password := loadVault(vaultPath(parsedCommand.VaultPath))
      secretName := parsedCommand.Name
      secretValue := input.AskMultiline(fmt.Sprintf("Enter value for '%s':\n", parsedCommand.Name))
      newSecret, newSecrets := vault.Add(secrets, secretName, secretValue)
      storage.Write(path, password, newSecrets)
      fmt.Println("Added:", newSecret)
    case command.Delete:
      secrets, path, password := loadVault(vaultPath(parsedCommand.VaultPath))
      deletedSecret, newSecrets, err := vault.Delete(secrets, parsedCommand.Id)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      } else {
        storage.Write(path, password, newSecrets)
        fmt.Println("Deleted:", deletedSecret)
      }
    case command.Edit:
      secrets, path, password := loadVault(vaultPath(parsedCommand.VaultPath))
      newName := input.Ask(fmt.Sprintf("Enter new name: "))
      newValue := input.AskMultiline("Enter new value:\n")
      editedSecret, newSecrets, err := vault.Edit(secrets, parsedCommand.Id, newName, newValue)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      } else {
        storage.Write(path, password, newSecrets)
        fmt.Println("Edited:", editedSecret)
      }
    case command.ChangePassword:
      currentPassword := askPassword()
      newPassword := input.AskPassword("Enter new vault password: ")
      newPasswordConfirmation := input.AskPassword("Enter new vault password again: ")

      if err := vault.ChangePassword(vaultPath(parsedCommand.VaultPath), currentPassword, newPassword, newPasswordConfirmation); err != nil {
        fmt.Println(err)
        os.Exit(1)
      } else {
        fmt.Println("Vault password successfully changed!")
      }
    case command.Initialize:
      configurationPath := path.Store(parsedCommand.VaultPath)
      fmt.Println(fmt.Sprintf("Vault successfully configured at %s and is ready to store your secrets!", configurationPath))
    default:
      fmt.Printf("Unhandled command: %T[%v]\n", parsedCommand, parsedCommand)
  }
}

func loadVault(vaultPath string) ([]secret.Secret, string, []byte) {
  password := askPassword()
  secrets, err := storage.Read(vaultPath, password)

  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }

  return secrets, vaultPath, password
}

func vaultPath(customPath string) string {
  if customPath != "" {
    return customPath
  } else {
    vaultPath, err := path.Get()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    return vaultPath
  }
}

func askPassword() []byte {
  return input.AskPassword("Enter vault password: ")
}
