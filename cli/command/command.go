package command

import (
  "github.com/satori/go.uuid"
)

type List struct {
  Filter string
  VaultPath string
}

type Add struct {
  Name string
  VaultPath string
}

type Edit struct {
  Id uuid.UUID
  VaultPath string
}

type Delete struct {
  Id uuid.UUID
  VaultPath string
}

type ChangePassword struct {
  VaultPath string
}

type Initialize struct {
  VaultPath string
}
