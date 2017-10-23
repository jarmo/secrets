package add

import (
  "fmt"
  "github.com/jarmo/secrets/secret"
  "github.com/jarmo/secrets/input"
)

func Execute(secrets []secret.Secret, name string) (secret.Secret, []secret.Secret) {
  newSecret := secret.New(name, input.AskMultiline(fmt.Sprintf("Enter value for '%s':\n", name)))
  return newSecret, append(secrets, newSecret)
}
