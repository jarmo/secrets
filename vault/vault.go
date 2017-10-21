package vault

import (
  "fmt"
  "strings"
  "os"
  "bufio"
  "github.com/jarmo/secrets/secret"
  "github.com/jarmo/secrets/vault/storage"
)

func List(filter string) []secret.Secret {
  var secrets []secret.Secret
  for _, secret := range storage.Read(storagePath()) {
    if secret.Id.String() == filter ||
         strings.Index(strings.ToLower(secret.Name), strings.ToLower(filter)) != -1 {
      secrets = append(secrets, secret)
    }
  }

  return secrets
}

func Add(name string) secret.Secret {
  secrets := storage.Read(storagePath())

  fmt.Println("Enter value:")
  var value []string
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
      value = append(value, scanner.Text())
  }

  secret := secret.Create(name, strings.Join(value, "\n"))
  storage.Write(storagePath(), append(secrets, secret))
  return secret
}

func storagePath() string {
  return "/Users/jarmo/.secrets.json"
}
