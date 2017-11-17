package add

import (
  "testing"
  "fmt"
  "github.com/jarmo/secrets/secret"
)

func TestExecute(t *testing.T) {
  secret1 := secret.New("name-1", "value-1")
  secret2 := secret.New("name-2", "value-2")
  secrets := [...]secret.Secret{secret1, secret2}

  addedSecret, newSecrets := Execute(secrets[:], "new-secret-name", "new-secret-value")
  expectedAddedSecret := secret.Secret{addedSecret.Id, "new-secret-name", "new-secret-value"}

  if addedSecret != expectedAddedSecret {
    t.Fatal(fmt.Sprintf("Expected %s to be added, but was %s", expectedAddedSecret, addedSecret))
  }

  expectedNewSecrets := [...]secret.Secret{secret1, secret2, addedSecret}

  if fmt.Sprintf("%v", newSecrets) != fmt.Sprintf("%v", expectedNewSecrets) {
    t.Fatal(fmt.Sprintf("Expected new secrets to be %s, but got %s", expectedNewSecrets, newSecrets))
  }
}
