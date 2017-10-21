package crypto

import (
  "encoding/hex"
  "crypto/rand"
  "encoding/json"
  "golang.org/x/crypto/scrypt"
  "golang.org/x/crypto/nacl/secretbox"
  "github.com/jarmo/secrets/secret"
  "github.com/jarmo/secrets/secret/encrypted"
)

func Encrypt(password []byte, secret secret.Secret) encrypted.Secret {
  if encryptedSecretJSON, err := json.Marshal(secret); err != nil {
    panic(err)
  } else {
    salt := generateRandomBytes(32)
    secretKey := calculateSecretKey(password, salt)
    var nonce [24]byte
    copy(nonce[:], generateRandomBytes(24))

    encryptedSecretData := secretbox.Seal(nil, encryptedSecretJSON, &nonce, &secretKey)
    return encrypted.Create(encryptedSecretData, nonce, salt)
  }
}

func Decrypt(password []byte, encryptedSecret encrypted.Secret) secret.Secret {
  salt, _ := hex.DecodeString(encryptedSecret.Salt)
  secretKey := calculateSecretKey(password, []byte(salt))
  data, _ := hex.DecodeString(encryptedSecret.Data)
  nonceBytes, _ := hex.DecodeString(encryptedSecret.Nonce)
  var nonce [24]byte
  copy(nonce[:], nonceBytes)
  var decryptedSecret secret.Secret

  if decryptedSecretJSON, ok := secretbox.Open(nil, data, &nonce, &secretKey); !ok {
    panic("Invalid password!")
  } else if err := json.Unmarshal(decryptedSecretJSON, &decryptedSecret); err != nil {
    panic(err)
  }

  return decryptedSecret
}

func calculateSecretKey(password, salt []byte) [32]byte {
  cpuFactor := 32768
  memoryFactor := 8
  parallelFactor := 1
  keyLength := 64

  secretKeyBytes, err := scrypt.Key(password, salt, cpuFactor, memoryFactor, parallelFactor, keyLength)
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
