package storage

import (
  "testing"
  "fmt"
  "io/ioutil"
  "os"
  "encoding/json"
  "github.com/jarmo/secrets/secret"
  "github.com/jarmo/secrets/crypto"
  "github.com/satori/go.uuid"
)

func TestWrite(t *testing.T) {
  vaultPath, err := ioutil.TempFile("", "test-secrets-vault")
  if err != nil {
    t.Fatal(err)
  }
  vaultPathStr := vaultPath.Name()
  defer os.Remove(vaultPathStr)

  Write(vaultPathStr, []byte("secret-password"), secrets())

  fileInfo, err := os.Stat(vaultPathStr)
  if fileInfo.Mode() != 0600 {
    t.Fatal("Expected vault permissions to be -rw-------, but are:", fileInfo.Mode())
  }

  encryptedSecretsJSON, err := ioutil.ReadFile(vaultPathStr)
  if err != nil {
    t.Fatal(err)
  }

  var encryptedSecrets crypto.Encrypted
  err = json.Unmarshal(encryptedSecretsJSON, &encryptedSecrets)

  if err != nil {
    t.Fatal(err)
  }
}

func TestRead_UsingArgon2idKey(t *testing.T) {
  decryptedSecrets, err := Read("storage_test_argon2id_input.json", []byte("secret-password"))

  if err != nil {
    t.Fatal(err)
  }

  expectedSecrets := secrets()
  id1, _ := uuid.FromString("910fe994-f102-407e-a4bd-947444a2ecb1")
  expectedSecrets[0].Id = id1
  id2, _ := uuid.FromString("6acd0733-0f2c-4138-a4ad-9cc98b8b61cc")
  expectedSecrets[1].Id = id2

  if fmt.Sprintf("%v", decryptedSecrets) != fmt.Sprintf("%v", expectedSecrets) {
    t.Fatal(fmt.Sprintf("Expected decrypted secrets to be %s, but got %s", expectedSecrets, decryptedSecrets))
  }
}

func TestRead_WithInvalidPasswordUsingArgon2idKey(t *testing.T) {
  decryptedSecrets, err := Read("storage_test_argon2id_input.json", []byte("wrong-password"))

  if len(decryptedSecrets) != 0 {
    t.Fatal(fmt.Sprintf("Expected no secrets, but got: %v", decryptedSecrets))
  }

  if err.Error() != "Invalid vault password!" {
    t.Fatal(fmt.Sprintf("Expected invalid password error message but got: %v", err))
  }
}

func TestRead_NoVault(t *testing.T) {
  decryptedSecrets, err := Read("non-existing-file", []byte("secret-password"))

  if err != nil {
    t.Fatal(err)
  }

  if len(decryptedSecrets) != 0 {
    t.Fatal(fmt.Sprintf("Expected to have 0 secrets, but got: %d", len(decryptedSecrets)))
  }
}

func secrets() []secret.Secret {
  secret1 := secret.New("foo", "goo")
  secret2 := secret.New("baz", "boo")
  return []secret.Secret{secret1, secret2}
}
