package vault

import (
  "github.com/jarmo/secrets/secret"
  "github.com/jarmo/secrets/vault/storage"
  "github.com/jarmo/secrets/vault/add"
  "github.com/jarmo/secrets/vault/list"
)

func List(filter string) []secret.Secret {
  return list.Execute(storage.Read(storagePath()), filter)
}

func Add(name string) secret.Secret {
  secrets := storage.Read(storagePath())
  secret := add.Execute(secrets, name)
  storage.Write(storagePath(), append(secrets, secret))
  return secret
}

func storagePath() string {
  return "/Users/jarmo/.secrets.json"
}
