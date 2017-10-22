package vault

import (
  "errors"
  "github.com/jarmo/secrets/input"
  "github.com/jarmo/secrets/secret"
  "github.com/jarmo/secrets/vault/storage"
  "github.com/jarmo/secrets/vault/add"
  "github.com/jarmo/secrets/vault/list"
  "github.com/jarmo/secrets/vault/delete"
  "github.com/jarmo/secrets/vault/edit"
  "github.com/satori/go.uuid"
)

func List(filter string) []secret.Secret {
  return list.Execute(storage.Read(input.AskPassword(), storagePath()), filter)
}

func Add(name string) secret.Secret {
  password := input.AskPassword()
  existingSecrets := storage.Read(password, storagePath())
  newSecret, newSecrets := add.Execute(existingSecrets, name)
  storage.Write(password, storagePath(), newSecrets)
  return newSecret
}

func Delete(id uuid.UUID) (secret.Secret, error) {
  password := input.AskPassword()
  existingSecrets := storage.Read(password, storagePath())
  deletedSecret, newSecrets, err := delete.Execute(existingSecrets, id)
  if err != nil {
    return deletedSecret, err
  }
  storage.Write(password, storagePath(), newSecrets)

  return deletedSecret, nil
}

func Edit(id uuid.UUID) (*secret.Secret, error) {
  password := input.AskPassword()
  existingSecrets := storage.Read(password, storagePath())
  existingSecretIndex := findIndexById(existingSecrets, id)
  if existingSecretIndex == -1 {
    return nil, errors.New("Secret by specified id not found!")
  }

  editedSecret, newSecrets := edit.Execute(existingSecrets, existingSecretIndex)
  storage.Write(password, storagePath(), newSecrets)

  return &editedSecret, nil
}

func storagePath() string {
  return "/Users/jarmo/.secrets.json"
}

func findIndexById(secrets []secret.Secret, id uuid.UUID) int {
  for index, secret := range secrets {
    if secret.Id == id {
      return index
    }
  }

  return -1
}
