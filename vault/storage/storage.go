package storage

import(
  "os"
  "io/ioutil"
  "encoding/json"
  "github.com/jarmo/secrets/secret"
)

func Read(path string) []secret.Secret {
  if data, err := ioutil.ReadFile(path); os.IsNotExist(err) {
    return make([]secret.Secret, 0)
  } else {
    var secrets []secret.Secret
    if err := json.Unmarshal(data, &secrets); err != nil {
      panic(err)
    } else {
      return secrets
    }
  }
}

func Write(path string, secrets []secret.Secret) {
  secretsJSON, _ := json.Marshal(secrets)
  if err := ioutil.WriteFile(path, secretsJSON, 0600); err != nil {
    panic(err)
  }
}

