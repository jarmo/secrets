package cli

import (
  "github.com/jarmo/secrets/command"
  "github.com/docopt/docopt-go"
  "github.com/satori/go.uuid"
)

func Execute(version string, args []string) interface{} {
  arguments, _ := docopt.Parse(createUsage(), args, true, version, false)
  return createCommand(arguments)
}

func createUsage() string {
  return `secrets

Usage:
  secrets --list [ID|NAME|VALUE]
  secrets --add NAME
  secrets --edit ID
  secrets --delete ID
  secrets --change-password

Options:
  -l --list            List all secrets in the vault or filter by id, partial name or value.
  -a --add             Add a new secret to the vault.
  -e --edit            Edit secret in the vault by id.
  -d --delete          Delete secret from the vault by id.
  --change-password    Change the vault password.
  -h --help            Show this screen.
  -v --version         Show version.`
}

func createCommand(arguments map[string]interface {}) interface{} {
  if arguments["--list"].(bool) {
    if filter, hasValue := arguments["ID"].(string); !hasValue {
      return command.List{Filter: ""}
    } else {
      return command.List{Filter: filter}
    }
  } else if arguments["--add"].(bool) {
    return command.Add{Name: arguments["NAME"].(string)}
  } else if arguments["--edit"].(bool) {
    id, _ := uuid.FromString(arguments["ID"].(string))
    return command.Edit{Id: id}
  } else if arguments["--delete"].(bool) {
    id, _ := uuid.FromString(arguments["ID"].(string))
    return command.Delete{Id: id}
  } else if arguments["--change-password"].(bool) {
    return command.ChangePassword{}
  } else {
    return nil
  }
}
