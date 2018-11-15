package crypto

import (
  "testing"
  "fmt"
  "github.com/jarmo/secrets/secret"
)

func TestEncryption(t *testing.T) {
  secrets := []secret.Secret{secret.New("foo", "bar")}
  password := "secret-password"

  encryptedSecrets := Encrypt([]byte(password), secrets)
  decryptedSecrets, err := Decrypt([]byte(password), encryptedSecrets)

  if err != nil {
    t.Fatal(err)
  }

  if fmt.Sprintf("%v", decryptedSecrets) != fmt.Sprintf("%v", secrets) {
    t.Fatalf("Expected decrypted secrets to be '%s', but got '%s'", secrets, decryptedSecrets)
  }
}

func TestDecryption_WithInvalidPassword(t *testing.T) {
  secrets := []secret.Secret{secret.New("foo", "bar")}
  password := "secret-password"

  encryptedSecrets := Encrypt([]byte(password), secrets)
  decryptedSecrets, err := Decrypt([]byte("wrong-password"), encryptedSecrets)

  if len(decryptedSecrets) != 0 {
    t.Fatalf("Expected no secrets, but got: %v", decryptedSecrets)
  }

  if err.Error() != "Invalid vault password!" {
    t.Fatalf("Expected invalid password error message but got: %v", err)
  }
}
