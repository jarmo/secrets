package vault

import (
  "errors"
  "bytes"
  "github.com/jarmo/secrets/input"
  "github.com/jarmo/secrets/secret"
  "github.com/jarmo/secrets/vault/storage"
  "github.com/jarmo/secrets/vault/storage/path"
  "github.com/jarmo/secrets/vault/add"
  "github.com/jarmo/secrets/vault/list"
  "github.com/jarmo/secrets/vault/delete"
  "github.com/jarmo/secrets/vault/edit"
  "github.com/satori/go.uuid"
)

func List(filter string) []secret.Secret {
  storagePath := path.Get()
  return list.Execute(storage.Read(askPassword(), storagePath), filter)
}

func Add(name string) secret.Secret {
  storagePath := path.Get()
  password := askPassword()
  existingSecrets := storage.Read(password, storagePath)
  newSecret, newSecrets := add.Execute(existingSecrets, name)
  storage.Write(password, storagePath, newSecrets)
  return newSecret
}

func Delete(id uuid.UUID) (*secret.Secret, error) {
  storagePath := path.Get()
  password := askPassword()
  existingSecrets := storage.Read(password, storagePath)
  existingSecretIndex := findIndexById(existingSecrets, id)
  if existingSecretIndex == -1 {
    return nil, errors.New("Secret by specified id not found!")
  }

  deletedSecret, newSecrets := delete.Execute(existingSecrets, existingSecretIndex)
  storage.Write(password, storagePath, newSecrets)

  return &deletedSecret, nil
}

func Edit(id uuid.UUID) (*secret.Secret, error) {
  storagePath := path.Get()
  password := askPassword()
  existingSecrets := storage.Read(password, storagePath)
  existingSecretIndex := findIndexById(existingSecrets, id)
  if existingSecretIndex == -1 {
    return nil, errors.New("Secret by specified id not found!")
  }

  editedSecret, newSecrets := edit.Execute(existingSecrets, existingSecretIndex)
  storage.Write(password, storagePath, newSecrets)

  return &editedSecret, nil
}

func ChangePassword() error {
  storagePath := path.Get()
  currentPassword := input.AskPassword("Enter vault password: ")
  secrets := storage.Read(currentPassword, storagePath)

  newPassword := input.AskPassword("Enter new vault password: ")
  newPasswordConfirmation := input.AskPassword("Enter new vault password again: ")

  if !bytes.Equal(newPassword, newPasswordConfirmation) {
    return errors.New("Passwords do not match!")
  }

  storage.Write(newPassword, storagePath, secrets)

  return nil
}

func askPassword() []byte {
  return input.AskPassword("Enter vault password: ")
}

func findIndexById(secrets []secret.Secret, id uuid.UUID) int {
  for index, secret := range secrets {
    if secret.Id == id {
      return index
    }
  }

  return -1
}
