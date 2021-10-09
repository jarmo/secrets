package list

import (
  "testing"
  "fmt"
  "github.com/jarmo/secrets/secret"
)

func TestExecute_WithNameFilter(t *testing.T) {
  secret1 := secret.New("foo", "goo")
  secret2 := secret.New("Baz", "boo")
  secret3 := secret.New("bar", "brr")
  secrets := [...]secret.Secret{secret1, secret2, secret3}

  listedSecrets := Execute(secrets[:], "ba")

  expectedListedSecrets := [...]secret.Secret{secret3, secret2}

  if fmt.Sprintf("%v", listedSecrets) != fmt.Sprintf("%v", expectedListedSecrets) {
    t.Fatal(fmt.Sprintf("Expected listed secrets to be %s, but got %s", expectedListedSecrets, listedSecrets))
  }
}

func TestExecute_WithValueFilter(t *testing.T) {
  secret1 := secret.New("foo", "goo")
  secret2 := secret.New("Baz", "boo")
  secret3 := secret.New("bar", "brr")
  secrets := [...]secret.Secret{secret1, secret2, secret3}

  listedSecrets := Execute(secrets[:], "boo")

  expectedListedSecrets := [...]secret.Secret{secret2}

  if fmt.Sprintf("%v", listedSecrets) != fmt.Sprintf("%v", expectedListedSecrets) {
    t.Fatal(fmt.Sprintf("Expected listed secrets to be %s, but got %s", expectedListedSecrets, listedSecrets))
  }
}

func TestExecute_WithIdFilter(t *testing.T) {
  secret1 := secret.New("foo", "goo")
  secret2 := secret.New("Baz", "boo")
  secret3 := secret.New("bar", "brr")
  secrets := [...]secret.Secret{secret1, secret2, secret3}

  listedSecrets := Execute(secrets[:], secret3.Id.String())

  expectedListedSecrets := [...]secret.Secret{secret3}

  if fmt.Sprintf("%v", listedSecrets) != fmt.Sprintf("%v", expectedListedSecrets) {
    t.Fatal(fmt.Sprintf("Expected listed secrets to be %s, but got %s", expectedListedSecrets, listedSecrets))
  }
}

func TestExecute_WithoutAnyFilter(t *testing.T) {
  secret1 := secret.New("foo", "goo")
  secret2 := secret.New("Baz", "boo")
  secret3 := secret.New("bar", "brr")
  secrets := [...]secret.Secret{secret1, secret2, secret3}

  listedSecrets := Execute(secrets[:], "")

  expectedListedSecrets := [...]secret.Secret{secret3, secret2, secret1}

  if fmt.Sprintf("%v", listedSecrets) != fmt.Sprintf("%v", expectedListedSecrets) {
    t.Fatal(fmt.Sprintf("Expected listed secrets to be %s, but got %s", expectedListedSecrets, listedSecrets))
  }
}
