package delete

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

  deletedSecret, newSecrets := Execute(secrets[:], 1)

  if deletedSecret != secret2 {
    t.Fatal(fmt.Sprintf("Expected %s to be deleted, but was %s", secret2, deletedSecret))
  }

  expectedNewSecrets := [...]secret.Secret{secret1, secret3}

  if fmt.Sprintf("%v", newSecrets) != fmt.Sprintf("%v", expectedNewSecrets) {
    t.Fatal(fmt.Sprintf("Expected new secrets to be %s, but got %s", expectedNewSecrets, newSecrets))
  }
}
