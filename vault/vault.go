package vault

import (
  "os"
  "io/ioutil"
  "encoding/json"
  "github.com/jarmo/secrets/secret"
)

func List(filter string) []secret.Secret {
  return read()
}

func Add(name string) secret.Secret {
  secrets := read()
  secret := secret.Create(name, "value")
  write(append(secrets, secret))
  return secret
}

func read() []secret.Secret {
  if data, err := ioutil.ReadFile(storagePath()); os.IsNotExist(err) {
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

func write(secrets []secret.Secret) {
  secretsJSON, _ := json.Marshal(secrets)
  if err := ioutil.WriteFile(storagePath(), secretsJSON, 0600); err != nil {
    panic(err)
  }
}

func storagePath() string {
  return "/Users/jarmo/.secrets.json"
}
