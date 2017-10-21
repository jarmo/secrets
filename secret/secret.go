package secret

import (
  "github.com/satori/go.uuid"
)

type Secret struct {
  Id uuid.UUID
  Name string
  Value string
}

func Create(name, value string) Secret {
  return Secret{uuid.NewV4(), name, value}
}
