package commands

type List struct {
  Filter string
}

type Add struct {
  Name string
}

type Edit struct {
  Id string
}

type Delete struct {
  Id string
}

type ChangePassword struct {
}
