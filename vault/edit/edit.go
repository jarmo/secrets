package edit

import (
  "fmt"
  "github.com/jarmo/secrets/secret"
  "github.com/jarmo/secrets/input"
)

func Execute(secrets []secret.Secret, index int) (secret.Secret, []secret.Secret) {
  editedSecret := secrets[index]
  newSecret := secret.Create(input.Ask(fmt.Sprintf("Enter new name for '%s': ", editedSecret.Name)), input.AskMultiline("Enter new value:\n"))
  newSecret.Id = editedSecret.Id
  secrets[index] = newSecret
  return editedSecret, secrets
}

