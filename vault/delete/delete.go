package delete

import (
  "github.com/jarmo/secrets/secret"
  "github.com/satori/go.uuid"
  "errors"
)

func Execute(secrets []secret.Secret, id uuid.UUID) (secret.Secret, []secret.Secret, error) {
  index, secret := findById(secrets, id)
  if index == -1 {
    return secret, secrets, errors.New("Secret by specified id not found!")
  }

  return secret, append(secrets[:index], secrets[index + 1:]...), nil
}

func findById(secrets []secret.Secret, id uuid.UUID) (int, secret.Secret) {
  for index, secret := range secrets {
    if secret.Id == id {
      return index, secret
    }
  }

  return -1, secret.Secret{}
}
