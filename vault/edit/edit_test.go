package edit

import (
  "testing"
  "fmt"
  "github.com/jarmo/secrets/v5/secret"
)

func TestExecute(t *testing.T) {
  secret1 := secret.New("name-1", "value-1")
  secret2 := secret.New("name-2", "value-2")
  secret3 := secret.New("name-3", "value-3")
  secrets := [...]secret.Secret{secret1, secret2, secret3}

  editedSecret, newSecrets := Execute(secrets[:], 1, "new-secret-name", "new-secret-value")

  expectedEditedSecret := secret.Secret{secret2.Id, "new-secret-name", "new-secret-value"}

  if editedSecret != expectedEditedSecret {
    t.Fatal(fmt.Sprintf("Expected %s to be edited, but was %s", expectedEditedSecret, editedSecret))
  }

  expectedNewSecrets := [...]secret.Secret{secret1, expectedEditedSecret, secret3}

  if fmt.Sprintf("%v", newSecrets) != fmt.Sprintf("%v", expectedNewSecrets) {
    t.Fatal(fmt.Sprintf("Expected new secrets to be %s, but got %s", expectedNewSecrets, newSecrets))
  }
}
