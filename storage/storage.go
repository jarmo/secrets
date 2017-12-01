package storage

import(
  "os"
  "io/ioutil"
  "encoding/json"
  "github.com/jarmo/secrets/crypto"
  "github.com/jarmo/secrets/secret"
)

func Read(path string, password []byte) []secret.Secret {
  if encryptedSecretsJSON, err := ioutil.ReadFile(path); os.IsNotExist(err) {
    return make([]secret.Secret, 0)
  } else {
    var encryptedSecrets crypto.Encrypted
    if err := json.Unmarshal(encryptedSecretsJSON, &encryptedSecrets); err != nil {
      panic(err)
    }

    return crypto.Decrypt(password, encryptedSecrets)
  }
}

func Write(path string, password []byte, decryptedSecrets []secret.Secret) {
  encryptedSecrets := crypto.Encrypt(password, decryptedSecrets)

  if encryptedSecretsJSON, err := json.Marshal(encryptedSecrets); err != nil {
    panic(err)
  } else if err := ioutil.WriteFile(path, encryptedSecretsJSON, 0600); err != nil {
    panic(err)
  }
}
