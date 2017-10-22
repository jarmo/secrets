package command

import (
  "github.com/satori/go.uuid"
)

type List struct {
  Filter string
}

type Add struct {
  Name string
}

type Edit struct {
  Id uuid.UUID
}

type Delete struct {
  Id uuid.UUID
}

type ChangePassword struct {
}
