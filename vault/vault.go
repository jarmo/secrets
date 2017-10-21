package vault

import (
  "fmt"
  "syscall"
  "github.com/jarmo/secrets/secret"
  "github.com/jarmo/secrets/vault/storage"
  "github.com/jarmo/secrets/vault/add"
  "github.com/jarmo/secrets/vault/list"
  "golang.org/x/crypto/ssh/terminal"
)

func List(filter string) []secret.Secret {
  return list.Execute(storage.Read(askPassword(), storagePath()), filter)
}

func Add(name string) secret.Secret {
  password := askPassword()
  existingSecrets := storage.Read(password, storagePath())
  newSecret := add.Execute(existingSecrets, name)
  storage.Write(password, storagePath(), append(existingSecrets, newSecret))
  return newSecret
}

func storagePath() string {
  return "/Users/jarmo/.secrets.json"
}

func askPassword() []byte {
  fmt.Print("Enter vault password: ")
  password, err := terminal.ReadPassword(int(syscall.Stdin))
  if err != nil {
    panic(err)
  }

  return password
}
