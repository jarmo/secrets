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
  return `secrets COMMAND [OPTIONS]

Usage:
  secrets list [FILTER] [--vault-path=VAULT_PATH]
  secrets add NAME [--vault-path=VAULT_PATH]
  secrets edit ID [--vault-path=VAULT_PATH]
  secrets --delete ID [--vault-path=VAULT_PATH]
  secrets --change-password [--vault-path=VAULT_PATH]
  secrets --init-vault --vault-path=VAULT_PATH

Options:
  -d --delete              Delete secret from the vault by id.
  --change-password        Change the vault password.
  --vault-path VAULT_PATH  Optional vault path. Defaults to the path in configuration.
  --init-vault             Initialize vault to specified path.
  -h --help                Show this screen.
  -v --version             Show version.`
}

func createCommand(arguments map[string]interface {}) interface{} {
  vaultPath := vaultPath(arguments)

  if arguments["list"].(bool) {
    if filter, hasValue := arguments["FILTER"].(string); !hasValue {
      return command.List{Filter: "", VaultPath: vaultPath}
    } else {
      return command.List{Filter: filter, VaultPath: vaultPath}
    }
  } else if arguments["add"].(bool) {
    return command.Add{Name: arguments["NAME"].(string), VaultPath: vaultPath}
  } else if arguments["edit"].(bool) {
    id, _ := uuid.FromString(arguments["ID"].(string))
    return command.Edit{Id: id, VaultPath: vaultPath}
  } else if arguments["--delete"].(bool) {
    id, _ := uuid.FromString(arguments["ID"].(string))
    return command.Delete{Id: id, VaultPath: vaultPath}
  } else if arguments["--change-password"].(bool) {
    return command.ChangePassword{VaultPath: vaultPath}
  } else if arguments["--init-vault"].(bool) {
    return command.Initialize{VaultPath: vaultPath}
  } else {
    return nil
  }
}

func vaultPath(arguments map[string]interface {}) string {
  if vaultPath, hasValue := arguments["--vault-path"].(string); hasValue {
    return vaultPath
  } else {
    return ""
  }
}
