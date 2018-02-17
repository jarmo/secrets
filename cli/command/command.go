package command

import (
  "github.com/satori/go.uuid"
)

type List struct {
  Filter string
  VaultPath string
  VaultAlias string
}

type Add struct {
  Name string
  VaultPath string
  VaultAlias string
}

type Edit struct {
  Id uuid.UUID
  VaultPath string
  VaultAlias string
}

type Delete struct {
  Id uuid.UUID
  VaultPath string
  VaultAlias string
}

type ChangePassword struct {
  VaultPath string
  VaultAlias string
}

type Initialize struct {
  VaultPath string
  VaultAlias string
}
