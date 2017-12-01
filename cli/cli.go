package cli

import (
  "github.com/jarmo/secrets/cli/command"
  "github.com/docopt/docopt-go"
  "github.com/satori/go.uuid"
)

func Execute(version string, args []string) interface{} {
  arguments, _ := docopt.Parse(createUsage(), args, true, version, false)
  return createCommand(arguments)
}

func createUsage() string {
  return `secrets COMMAND [OPTION] [VAULT_PATH]

Usage:
  secrets --list [ID|NAME|VALUE] [VAULT_PATH]
  secrets --add NAME [VAULT_PATH]
  secrets --edit ID [VAULT_PATH]
  secrets --delete ID [VAULT_PATH]
  secrets --change-password [VAULT_PATH]

Arguments:
  VAULT_PATH           Optional parameter to specify vault absolute path.
                       When not specified, path is asked from stdin or read from configuration file.

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
  vaultPath := vaultPath(arguments)

  if arguments["--list"].(bool) {
    if filter, hasValue := arguments["ID"].(string); !hasValue {
      return command.List{Filter: "", VaultPath: vaultPath}
    } else {
      return command.List{Filter: filter, VaultPath: vaultPath}
    }
  } else if arguments["--add"].(bool) {
    return command.Add{Name: arguments["NAME"].(string), VaultPath: vaultPath}
  } else if arguments["--edit"].(bool) {
    id, _ := uuid.FromString(arguments["ID"].(string))
    return command.Edit{Id: id, VaultPath: vaultPath}
  } else if arguments["--delete"].(bool) {
    id, _ := uuid.FromString(arguments["ID"].(string))
    return command.Delete{Id: id, VaultPath: vaultPath}
  } else if arguments["--change-password"].(bool) {
    return command.ChangePassword{VaultPath: vaultPath}
  } else {
    return nil
  }
}

func vaultPath(arguments map[string]interface {}) string {
  if vaultPath, hasValue := arguments["VAULT_PATH"].(string); hasValue {
    return vaultPath
  } else {
    return ""
  }
}
