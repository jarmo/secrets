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

  switch parsedCommand := Execute(version, []string{"list", filter}).(type) {
    case command.List:
      if parsedCommand.Filter != filter {
        t.Fatal(fmt.Sprintf("Expected filter to be '%v' but was: '%v'", filter, parsedCommand.Filter))
      }
      if parsedCommand.VaultAlias != "" {
        t.Fatal(fmt.Sprintf("Expected VaultAlias to be empty but was: '%v'", parsedCommand.VaultAlias))
      }
      if parsedCommand.VaultPath != "" {
        t.Fatal(fmt.Sprintf("Expected VaultPath to be empty but was: '%v'", parsedCommand.VaultPath))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_ListWithoutFilterAndWithCustomVaultPath(t *testing.T) {
  filter := ""
  vaultPath := "/foo/bar/baz"

  switch parsedCommand := Execute(version, []string{"list", "--path", vaultPath}).(type) {
    case command.List:
      if parsedCommand.Filter != filter {
        t.Fatal(fmt.Sprintf("Expected filter to be '%v' but was: '%v'", filter, parsedCommand.Filter))
      }
      if parsedCommand.VaultAlias != "" {
        t.Fatal(fmt.Sprintf("Expected VaultAlias to be empty but was: '%v'", parsedCommand.VaultAlias))
      }
      if parsedCommand.VaultPath != vaultPath {
        t.Fatal(fmt.Sprintf("Expected VaultPath to be '%v' but was: '%v'", vaultPath, parsedCommand.VaultPath))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_ListWithoutFilterAndWithVaultAlias(t *testing.T) {
  filter := ""
  vaultAlias := "foo-bar"

  switch parsedCommand := Execute(version, []string{"list", "--alias", vaultAlias}).(type) {
    case command.List:
      if parsedCommand.Filter != filter {
        t.Fatal(fmt.Sprintf("Expected filter to be '%v' but was: '%v'", filter, parsedCommand.Filter))
      }
      if parsedCommand.VaultAlias != vaultAlias {
        t.Fatal(fmt.Sprintf("Expected VaultAlias to be '%v' but was: '%v'", vaultAlias, parsedCommand.VaultAlias))
      }
      if parsedCommand.VaultPath != "" {
        t.Fatal(fmt.Sprintf("Expected VaultPath to be empty but was: '%v'", parsedCommand.VaultPath))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_ListWithFilter(t *testing.T) {
  filter := "custom-filter"

  switch parsedCommand := Execute(version, []string{"list", filter}).(type) {
    case command.List:
      if parsedCommand.Filter != filter {
        t.Fatal(fmt.Sprintf("Expected filter to be '%v', but was '%v'", filter, parsedCommand.Filter))
      }
      if parsedCommand.VaultAlias != "" {
        t.Fatal(fmt.Sprintf("Expected VaultAlias to be empty but was: '%v'", parsedCommand.VaultAlias))
      }
      if parsedCommand.VaultPath != "" {
        t.Fatal(fmt.Sprintf("Expected VaultPath to be empty but was: '%v'", parsedCommand.VaultPath))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_ListWithFilterAndWithCustomVaultPath(t *testing.T) {
  filter := "custom-filter"
  vaultPath := "/foo/bar/baz"

  switch parsedCommand := Execute(version, []string{"list", filter, "--path", vaultPath}).(type) {
    case command.List:
      if parsedCommand.Filter != filter {
        t.Fatal(fmt.Sprintf("Expected filter to be '%v', but was '%v'", filter, parsedCommand.Filter))
      }
      if parsedCommand.VaultAlias != "" {
        t.Fatal(fmt.Sprintf("Expected VaultAlias to be empty but was: '%v'", parsedCommand.VaultAlias))
      }
      if parsedCommand.VaultPath != vaultPath {
        t.Fatal(fmt.Sprintf("Expected VaultPath to be '%v' but was: '%v'", vaultPath, parsedCommand.VaultPath))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_ListWithFilterAndWithVaultAlias(t *testing.T) {
  filter := "custom-filter"
  vaultAlias := "foo-bar"

  switch parsedCommand := Execute(version, []string{"list", filter, "--alias", vaultAlias}).(type) {
    case command.List:
      if parsedCommand.Filter != filter {
        t.Fatal(fmt.Sprintf("Expected filter to be '%v', but was '%v'", filter, parsedCommand.Filter))
      }
      if parsedCommand.VaultAlias != vaultAlias {
        t.Fatal(fmt.Sprintf("Expected VaultAlias to be '%v' but was: '%v'", vaultAlias, parsedCommand.VaultAlias))
      }
      if parsedCommand.VaultPath != "" {
        t.Fatal(fmt.Sprintf("Expected VaultPath to be empty but was: '%v'", parsedCommand.VaultPath))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_Add(t *testing.T) {
  name := "custom-name"

  switch parsedCommand := Execute(version, []string{"add", name}).(type) {
    case command.Add:
      if parsedCommand.Name != name {
        t.Fatal(fmt.Sprintf("Expected name to be '%v', but was '%v'", name, parsedCommand.Name))
      }
      if parsedCommand.VaultAlias != "" {
        t.Fatal(fmt.Sprintf("Expected VaultAlias to be empty but was: '%v'", parsedCommand.VaultAlias))
      }
      if parsedCommand.VaultPath != "" {
        t.Fatal(fmt.Sprintf("Expected VaultPath to be empty but was: '%v'", parsedCommand.VaultPath))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_AddWithCustomVaultPath(t *testing.T) {
  name := "custom-name"
  vaultPath := "/foo/bar/baz"

  switch parsedCommand := Execute(version, []string{"add", name, "--path", vaultPath}).(type) {
    case command.Add:
      if parsedCommand.Name != name {
        t.Fatal(fmt.Sprintf("Expected name to be '%v', but was '%v'", name, parsedCommand.Name))
      }
      if parsedCommand.VaultAlias != "" {
        t.Fatal(fmt.Sprintf("Expected VaultAlias to be empty but was: '%v'", parsedCommand.VaultAlias))
      }
      if parsedCommand.VaultPath != vaultPath {
        t.Fatal(fmt.Sprintf("Expected VaultPath to be '%v' but was: '%v'", vaultPath, parsedCommand.VaultPath))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_AddWithVaultAlias(t *testing.T) {
  name := "custom-name"
  vaultAlias := "foo-bar"

  switch parsedCommand := Execute(version, []string{"add", name, "--alias", vaultAlias}).(type) {
    case command.Add:
      if parsedCommand.Name != name {
        t.Fatal(fmt.Sprintf("Expected name to be '%v', but was '%v'", name, parsedCommand.Name))
      }
      if parsedCommand.VaultAlias != vaultAlias {
        t.Fatal(fmt.Sprintf("Expected VaultAlias to be '%v' but was: '%v'", vaultAlias, parsedCommand.VaultAlias))
      }
      if parsedCommand.VaultPath != "" {
        t.Fatal(fmt.Sprintf("Expected VaultPath to be empty but was: '%v'", parsedCommand.VaultPath))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_Edit(t *testing.T) {
  id := uuid.NewV4()

  switch parsedCommand := Execute(version, []string{"edit", id.String()}).(type) {
    case command.Edit:
      if parsedCommand.Id != id {
        t.Fatal(fmt.Sprintf("Expected id to be '%v', but was '%v'", id, parsedCommand.Id))
      }
      if parsedCommand.VaultAlias != "" {
        t.Fatal(fmt.Sprintf("Expected VaultAlias to be empty but was: '%v'", parsedCommand.VaultAlias))
      }
      if parsedCommand.VaultPath != "" {
        t.Fatal(fmt.Sprintf("Expected VaultPath to be empty but was: '%v'", parsedCommand.VaultPath))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_EditWithCustomVaultPath(t *testing.T) {
  id := uuid.NewV4()
  vaultPath := "/foo/bar/baz"

  switch parsedCommand := Execute(version, []string{"edit", id.String(), "--path", vaultPath}).(type) {
    case command.Edit:
      if parsedCommand.Id != id {
        t.Fatal(fmt.Sprintf("Expected id to be '%v', but was '%v'", id, parsedCommand.Id))
      }
      if parsedCommand.VaultAlias != "" {
        t.Fatal(fmt.Sprintf("Expected VaultAlias to be empty but was: '%v'", parsedCommand.VaultAlias))
      }
      if parsedCommand.VaultPath != vaultPath {
        t.Fatal(fmt.Sprintf("Expected VaultPath to be '%v' but was: '%v'", vaultPath, parsedCommand.VaultPath))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_EditWithVaultAlias(t *testing.T) {
  id := uuid.NewV4()
  vaultAlias := "foo-bar"

  switch parsedCommand := Execute(version, []string{"edit", id.String(), "--alias", vaultAlias}).(type) {
    case command.Edit:
      if parsedCommand.Id != id {
        t.Fatal(fmt.Sprintf("Expected id to be '%v', but was '%v'", id, parsedCommand.Id))
      }
      if parsedCommand.VaultAlias != vaultAlias {
        t.Fatal(fmt.Sprintf("Expected VaultAlias to be '%v' but was: '%v'", vaultAlias, parsedCommand.VaultAlias))
      }
      if parsedCommand.VaultPath != "" {
        t.Fatal(fmt.Sprintf("Expected VaultPath to be empty but was: '%v'", parsedCommand.VaultPath))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_Delete(t *testing.T) {
  id := uuid.NewV4()

  switch parsedCommand := Execute(version, []string{"delete", id.String()}).(type) {
    case command.Delete:
      if parsedCommand.Id != id {
        t.Fatal(fmt.Sprintf("Expected id to be '%v', but was '%v'", id, parsedCommand.Id))
      }
      if parsedCommand.VaultAlias != "" {
        t.Fatal(fmt.Sprintf("Expected VaultAlias to be empty but was: '%v'", parsedCommand.VaultAlias))
      }
      if parsedCommand.VaultPath != "" {
        t.Fatal(fmt.Sprintf("Expected VaultPath to be empty but was: '%v'", parsedCommand.VaultPath))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_DeleteWithCustomVaultPath(t *testing.T) {
  id := uuid.NewV4()
  vaultPath := "/foo/bar/baz"

  switch parsedCommand := Execute(version, []string{"delete", id.String(), "--path", vaultPath}).(type) {
    case command.Delete:
      if parsedCommand.Id != id {
        t.Fatal(fmt.Sprintf("Expected id to be '%v', but was '%v'", id, parsedCommand.Id))
      }
      if parsedCommand.VaultAlias != "" {
        t.Fatal(fmt.Sprintf("Expected VaultAlias to be empty but was: '%v'", parsedCommand.VaultAlias))
      }
      if parsedCommand.VaultPath != vaultPath {
        t.Fatal(fmt.Sprintf("Expected VaultPath to be '%v' but was: '%v'", vaultPath, parsedCommand.VaultPath))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_DeleteWithVaultAlias(t *testing.T) {
  id := uuid.NewV4()
  vaultAlias := "foo-bar"

  switch parsedCommand := Execute(version, []string{"delete", id.String(), "--alias", vaultAlias}).(type) {
    case command.Delete:
      if parsedCommand.Id != id {
        t.Fatal(fmt.Sprintf("Expected id to be '%v', but was '%v'", id, parsedCommand.Id))
      }
      if parsedCommand.VaultAlias != vaultAlias {
        t.Fatal(fmt.Sprintf("Expected VaultAlias to be '%v' but was: '%v'", vaultAlias, parsedCommand.VaultAlias))
      }
      if parsedCommand.VaultPath != "" {
        t.Fatal(fmt.Sprintf("Expected VaultPath to be empty but was: '%v'", parsedCommand.VaultPath))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_ChangePassword(t *testing.T) {
  switch parsedCommand := Execute(version, []string{"change-password"}).(type) {
    case command.ChangePassword:
      if parsedCommand.VaultAlias != "" {
        t.Fatal(fmt.Sprintf("Expected VaultAlias to be empty but was: '%v'", parsedCommand.VaultAlias))
      }
      if parsedCommand.VaultPath != "" {
        t.Fatal(fmt.Sprintf("Expected VaultPath to be empty but was: '%v'", parsedCommand.VaultPath))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_ChangePasswordWithCustomVaultPath(t *testing.T) {
  vaultPath := "/foo/bar/baz"

  switch parsedCommand := Execute(version, []string{"change-password", "--path", vaultPath}).(type) {
    case command.ChangePassword:
      if parsedCommand.VaultAlias != "" {
        t.Fatal(fmt.Sprintf("Expected VaultAlias to be empty but was: '%v'", parsedCommand.VaultAlias))
      }
      if parsedCommand.VaultPath != vaultPath {
        t.Fatal(fmt.Sprintf("Expected VaultPath to be '%v' but was: '%v'", vaultPath, parsedCommand.VaultPath))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_ChangePasswordWithVaultAlias(t *testing.T) {
  vaultAlias := "foo-bar"

  switch parsedCommand := Execute(version, []string{"change-password", "--alias", vaultAlias}).(type) {
    case command.ChangePassword:
      if parsedCommand.VaultAlias != vaultAlias {
        t.Fatal(fmt.Sprintf("Expected VaultAlias to be '%v' but was: '%v'", vaultAlias, parsedCommand.VaultAlias))
      }
      if parsedCommand.VaultPath != "" {
        t.Fatal(fmt.Sprintf("Expected VaultPath to be empty but was: '%v'", parsedCommand.VaultPath))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}

func TestExecute_Initialize(t *testing.T) {
  vaultPath := "/foo/bar/baz"
  vaultAlias := "foo-bar"

  switch parsedCommand := Execute(version, []string{"initialize", "--path", vaultPath, "--alias", vaultAlias}).(type) {
    case command.Initialize:
      if parsedCommand.VaultPath != vaultPath {
        t.Fatalf("Expected vault path to be '%v', but was '%v'", vaultPath, parsedCommand.VaultPath)
      }
      if parsedCommand.VaultAlias != vaultAlias {
        t.Fatal(fmt.Sprintf("Expected VaultAlias to be '%v' but was: '%v'", vaultAlias, parsedCommand.VaultAlias))
      }
    default:
      t.Fatal(fmt.Sprintf("Got unexpected command: %T", parsedCommand))
  }
}
