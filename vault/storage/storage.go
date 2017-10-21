package storage

import(
  "os"
  "io/ioutil"
  "encoding/json"
  "github.com/jarmo/secrets/crypto"
  "github.com/jarmo/secrets/secret"
  "github.com/jarmo/secrets/secret/encrypted"
)

func Read(password []byte, path string) []secret.Secret {
  if encryptedSecretsJSON, err := ioutil.ReadFile(path); os.IsNotExist(err) {
    return make([]secret.Secret, 0)
  } else {
    var encryptedSecrets []encrypted.Secret
    if err := json.Unmarshal(encryptedSecretsJSON, &encryptedSecrets); err != nil {
      panic(err)
    }

    var decryptedSecrets []secret.Secret
    for _, encryptedSecret := range encryptedSecrets {
      decryptedSecrets = append(decryptedSecrets, crypto.Decrypt(password, encryptedSecret))
    }

    return decryptedSecrets
  }
}

func Write(password []byte, path string, decryptedSecrets []secret.Secret) {
  var encryptedSecrets []encrypted.Secret
  for _, decryptedSecret := range decryptedSecrets {
    encryptedSecrets = append(encryptedSecrets, crypto.Encrypt(password, decryptedSecret))
  }

  if encryptedSecretsJSON, err := json.Marshal(encryptedSecrets); err != nil {
    panic(err)
  } else if err := ioutil.WriteFile(path, encryptedSecretsJSON, 0600); err != nil {
    panic(err)
  }
}

