package crypto

import (
  "encoding/base64"
  "crypto/rand"
  "encoding/json"
  "errors"
  "golang.org/x/crypto/scrypt"
  "golang.org/x/crypto/nacl/secretbox"
  "github.com/jarmo/secrets/secret"
)

type Encrypted struct {
  Data string
  Nonce string
  Salt string
  N int
  R int
  P int
}

func Encrypt(password []byte, secrets []secret.Secret) Encrypted {
  if encryptedSecretJSON, err := json.Marshal(secrets); err != nil {
    panic(err)
  } else {
    salt := generateRandomBytes(32)
    N := 32768
    r := 8
    p := 2
    secretKey := calculateSecretKey(password, salt, N, r, p)
    var nonce [24]byte
    copy(nonce[:], generateRandomBytes(24))

    data := secretbox.Seal(nil, encryptedSecretJSON, &nonce, &secretKey)
    return Encrypted{Data: base64.StdEncoding.EncodeToString(data), Nonce: base64.StdEncoding.EncodeToString(nonce[:]), Salt: base64.StdEncoding.EncodeToString(salt), N: N, R: r, P: p}
  }
}

func Decrypt(password []byte, encryptedSecrets Encrypted) ([]secret.Secret, error) {
  salt, _ := base64.StdEncoding.DecodeString(encryptedSecrets.Salt)
  secretKey := calculateSecretKey(password, []byte(salt), encryptedSecrets.N, encryptedSecrets.R, encryptedSecrets.P)
  data, _ := base64.StdEncoding.DecodeString(encryptedSecrets.Data)
  nonceBytes, _ := base64.StdEncoding.DecodeString(encryptedSecrets.Nonce)
  var nonce [24]byte
  copy(nonce[:], nonceBytes)
  var decryptedSecrets []secret.Secret

  if decryptedSecretsJSON, ok := secretbox.Open(nil, data, &nonce, &secretKey); !ok {
    return make([]secret.Secret, 0), errors.New("Invalid vault password!")
  } else if err := json.Unmarshal(decryptedSecretsJSON, &decryptedSecrets); err != nil {
    panic(err)
  }

  return decryptedSecrets, nil
}

func calculateSecretKey(password, salt []byte, N, r, p int) [32]byte {
  keyLength := 32

  secretKeyBytes, err := scrypt.Key(password, salt, N, r, p, keyLength)
  if err != nil {
    panic(err)
  }

  var secretKey [32]byte
  copy(secretKey[:], secretKeyBytes)

  return secretKey
}

func generateRandomBytes(length int) []byte {
  result := make([]byte, length)
  _, err := rand.Read(result)
  if err != nil {
    panic(err)
  }

  return result
}
