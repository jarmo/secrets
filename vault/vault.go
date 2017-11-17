package vault

import (
  "errors"
  "bytes"
  "github.com/jarmo/secrets/secret"
  "github.com/jarmo/secrets/vault/storage"
  "github.com/jarmo/secrets/vault/add"
  "github.com/jarmo/secrets/vault/list"
  "github.com/jarmo/secrets/vault/delete"
  "github.com/jarmo/secrets/vault/edit"
  "github.com/satori/go.uuid"
)

func List(filter, storagePath string, password []byte) []secret.Secret {
  return list.Execute(storage.Read(password, storagePath), filter)
}

func Add(name, value, storagePath string, password []byte) secret.Secret {
  existingSecrets := storage.Read(password, storagePath)
  newSecret, newSecrets := add.Execute(existingSecrets, name, value)
  storage.Write(password, storagePath, newSecrets)
  return newSecret
}

func Delete(id uuid.UUID, storagePath string, password []byte) (*secret.Secret, error) {
  existingSecrets := storage.Read(password, storagePath)
  existingSecretIndex := findIndexById(existingSecrets, id)
  if existingSecretIndex == -1 {
    return nil, errors.New("Secret by specified id not found!")
  }

  deletedSecret, newSecrets := delete.Execute(existingSecrets, existingSecretIndex)
  storage.Write(password, storagePath, newSecrets)

  return &deletedSecret, nil
}

func Edit(id uuid.UUID, newName, newValue, storagePath string, password []byte) (*secret.Secret, error) {
  existingSecrets := storage.Read(password, storagePath)
  existingSecretIndex := findIndexById(existingSecrets, id)
  if existingSecretIndex == -1 {
    return nil, errors.New("Secret by specified id not found!")
  }

  editedSecret, newSecrets := edit.Execute(existingSecrets, existingSecretIndex, newName, newValue)
  storage.Write(password, storagePath, newSecrets)

  return &editedSecret, nil
}

func ChangePassword(storagePath string, currentPassword, newPassword, newPasswordConfirmation []byte) error {
  secrets := storage.Read(currentPassword, storagePath)

  if !bytes.Equal(newPassword, newPasswordConfirmation) {
    return errors.New("Passwords do not match!")
  }

  storage.Write(newPassword, storagePath, secrets)

  return nil
}

func findIndexById(secrets []secret.Secret, id uuid.UUID) int {
  for index, secret := range secrets {
    if secret.Id == id {
      return index
    }
  }

  return -1
}
