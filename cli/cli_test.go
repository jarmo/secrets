package cli

import (
  "testing"
  "fmt"
  "github.com/jarmo/secrets/cli/command"
  "github.com/satori/go.uuid"
)

const version = "1.3.3.7"

func TestExecute_ListWithoutFilter(t *testing.T) {
  filter := ""

  switch parsedCommand := Execute(version, []string{"--list", filter}).(type) {
    case command.List:
      if parsedCommand.Filter != filter {
        t.Fatal(fmt.Sprintf("Expected filter to be '%v' but was: '%v'", filter, parsedCommand.Filter))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_ListWithFilter(t *testing.T) {
  filter := "custom-filter"

  switch parsedCommand := Execute(version, []string{"--list", filter}).(type) {
    case command.List:
      if parsedCommand.Filter != filter {
        t.Fatal(fmt.Sprintf("Expected filter to be '%v', but was '%v'", filter, parsedCommand.Filter))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_Add(t *testing.T) {
  name := "custom-name"

  switch parsedCommand := Execute(version, []string{"--add", name}).(type) {
    case command.Add:
      if parsedCommand.Name != name {
        t.Fatal(fmt.Sprintf("Expected name to be '%v', but was '%v'", name, parsedCommand.Name))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_Edit(t *testing.T) {
  id := uuid.NewV4()

  switch parsedCommand := Execute(version, []string{"--edit", id.String()}).(type) {
    case command.Edit:
      if parsedCommand.Id != id {
        t.Fatal(fmt.Sprintf("Expected id to be '%v', but was '%v'", id, parsedCommand.Id))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_Delete(t *testing.T) {
  id := uuid.NewV4()

  switch parsedCommand := Execute(version, []string{"--delete", id.String()}).(type) {
    case command.Delete:
      if parsedCommand.Id != id {
        t.Fatal(fmt.Sprintf("Expected id to be '%v', but was '%v'", id, parsedCommand.Id))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_ChangePassword(t *testing.T) {
  switch parsedCommand := Execute(version, []string{"--change-password"}).(type) {
    case command.ChangePassword:
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_Initialize(t *testing.T) {
  vaultPath := "/foo/bar/baz"
  switch parsedCommand := Execute(version, []string{"--init-vault", "--vault-path", vaultPath}).(type) {
    case command.Initialize:
      if parsedCommand.VaultPath != "/foo/bar/baz" {
        t.Fatal("Expected vault path to be '%v', but was '%v'", vaultPath, parsedCommand.VaultPath)
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}
