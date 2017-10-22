package vault

import (
  "fmt"
  "syscall"
  "github.com/jarmo/secrets/secret"
  "github.com/jarmo/secrets/vault/storage"
  "github.com/jarmo/secrets/vault/add"
  "github.com/jarmo/secrets/vault/list"
  "github.com/jarmo/secrets/vault/delete"
  "golang.org/x/crypto/ssh/terminal"
  "github.com/satori/go.uuid"
)

func List(filter string) []secret.Secret {
  return list.Execute(storage.Read(askPassword(), storagePath()), filter)
}

func Add(name string) secret.Secret {
  password := askPassword()
  existingSecrets := storage.Read(password, storagePath())
  newSecret, newSecrets := add.Execute(existingSecrets, name)
  storage.Write(password, storagePath(), newSecrets)
  return newSecret
}

func Delete(id uuid.UUID) (secret.Secret, error) {
  password := askPassword()
  existingSecrets := storage.Read(password, storagePath())
  deletedSecret, newSecrets, err := delete.Execute(existingSecrets, id)
  if err != nil {
    return deletedSecret, err
  }
  storage.Write(password, storagePath(), newSecrets)

  return deletedSecret, nil
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
  fmt.Println()

  return password
}
