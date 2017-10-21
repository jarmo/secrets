package cli

import (
  "github.com/jarmo/secrets/commands"
  "github.com/docopt/docopt-go"
)

func Execute(version string, args []string) interface{} {
  arguments, _ := docopt.Parse(createUsage(), args, true, version, false)
  return createCommand(arguments)
}

func createUsage() string {
  return `secrets

Usage:
  secrets --list [ID|NAME]
  secrets --add NAME
  secrets --edit ID
  secrets --delete ID
  secrets --change-password

Options:
  -l --list            List all secrets in the vault or filter by id or name.
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
      return commands.List{Filter: ""}
    } else {
      return commands.List{Filter: filter}
    }
  } else if arguments["--add"].(bool) {
    return commands.Add{Name: arguments["NAME"].(string)}
  } else if arguments["--edit"].(bool) {
    return commands.Edit{Id: arguments["ID"].(string)}
  } else if arguments["--edit"].(bool) {
    return commands.Delete{Id: arguments["ID"].(string)}
  } else if arguments["--change-password"].(bool) {
    return commands.ChangePassword{}
  } else {
    return nil
  }
}
