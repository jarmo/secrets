package main

import (
  "os"
  "fmt"
  "github.com/jarmo/secrets/cli"
  "github.com/jarmo/secrets/cli/command"
  "github.com/jarmo/secrets/vault"
  "github.com/jarmo/secrets/vault/storage/path"
  "github.com/jarmo/secrets/input"
)

const VERSION = "0.0.1"

func main() {
  switch parsedCommand := cli.Execute(VERSION, os.Args[1:]).(type) {
    case command.List:
      secrets := vault.List(parsedCommand.Filter, path.Get(), askPassword())
      for _, secret := range secrets {
        fmt.Println(secret)
      }
    case command.Add:
      password := askPassword()
      secretName := parsedCommand.Name
      secretValue := input.AskMultiline(fmt.Sprintf("Enter value for '%s':\n", parsedCommand.Name))
      fmt.Println("Added:", vault.Add(secretName, secretValue, path.Get(), password))
    case command.Delete:
      deletedSecret, err := vault.Delete(parsedCommand.Id, path.Get(), askPassword())
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      } else {
        fmt.Println("Deleted:", deletedSecret)
      }
    case command.Edit:
      password := askPassword()
      newName := input.Ask(fmt.Sprintf("Enter new name: "))
      newValue := input.AskMultiline("Enter new value:\n")
      editedSecret, err := vault.Edit(parsedCommand.Id, newName, newValue, path.Get(), password)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      } else {
        fmt.Println("Edited:", editedSecret)
      }
    case command.ChangePassword:
      currentPassword := askPassword()
      newPassword := input.AskPassword("Enter new vault password: ")
      newPasswordConfirmation := input.AskPassword("Enter new vault password again: ")

      if err := vault.ChangePassword(path.Get(), currentPassword, newPassword, newPasswordConfirmation); err != nil {
        fmt.Println(err)
        os.Exit(1)
      } else {
        fmt.Println("Vault password successfully changed!")
      }
    default:
      fmt.Printf("Unhandled command: %T\n", parsedCommand)
  }
}

func askPassword() []byte {
  return input.AskPassword("Enter vault password: ")
}
