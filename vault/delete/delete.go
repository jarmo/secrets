package delete

import (
  "github.com/jarmo/secrets/v5/secret"
)

func Execute(secrets []secret.Secret, index int) (secret.Secret, []secret.Secret) {
  deletedSecret := secrets[index]
  return deletedSecret, append(secrets[:index], secrets[index + 1:]...)
}
