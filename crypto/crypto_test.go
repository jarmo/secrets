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
  decryptedSecrets := Decrypt([]byte(password), encryptedSecrets)

  if fmt.Sprintf("%v", decryptedSecrets) != fmt.Sprintf("%v", secrets) {
    t.Fatal(fmt.Sprintf("Expected decrypted secrets to be '%s', but got '%s'", secrets, decryptedSecrets))
  }
}
