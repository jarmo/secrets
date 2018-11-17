package cli

import (
  "github.com/docopt/docopt-go"
  "github.com/satori/go.uuid"
  "github.com/jarmo/secrets/cli/command"
)

func Command(version string, args []string) command.Executable {
  arguments, _ := docopt.Parse(createUsage(), args, true, version, false)
  return createCommand(arguments)
}

func createUsage() string {
  return `secrets COMMAND [OPTIONS]

Usage:
  secrets list [FILTER] [--alias=VAULT_ALIAS | --path=VAULT_PATH]
  secrets add NAME [--alias=VAULT_ALIAS | --path=VAULT_PATH]
  secrets edit ID [--alias=VAULT_ALIAS | --path=VAULT_PATH]
  secrets delete ID [--alias=VAULT_PATH | --path=VAULT_PATH]
  secrets change-password [--alias=VAULT_PATH | --path=VAULT_PATH]
  secrets initialize --path=VAULT_PATH --alias=VAULT_ALIAS

Options:
  --alias VAULT_ALIAS    Optional vault alias.
  --path VAULT_PATH      Optional vault path. Defaults to the path in configuration.
  -h --help              Show this screen.
  -v --version           Show version.`
}

func createCommand(arguments map[string]interface {}) command.Executable {
  vaultAlias := vaultAlias(arguments)
  vaultPath := vaultPath(arguments)

  if arguments["list"].(bool) {
    if filter, hasValue := arguments["FILTER"].(string); !hasValue {
      return command.List{Filter: "", VaultAlias: vaultAlias, VaultPath: vaultPath}
    } else {
      return command.List{Filter: filter, VaultAlias: vaultAlias, VaultPath: vaultPath}
    }
  } else if arguments["add"].(bool) {
    return command.Add{Name: arguments["NAME"].(string), VaultAlias: vaultAlias, VaultPath: vaultPath}
  } else if arguments["edit"].(bool) {
    id, _ := uuid.FromString(arguments["ID"].(string))
    return command.Edit{Id: id, VaultAlias: vaultAlias, VaultPath: vaultPath}
  } else if arguments["delete"].(bool) {
    id, _ := uuid.FromString(arguments["ID"].(string))
    return command.Delete{Id: id, VaultAlias: vaultAlias, VaultPath: vaultPath}
  } else if arguments["change-password"].(bool) {
    return command.ChangePassword{VaultAlias: vaultAlias, VaultPath: vaultPath}
  } else if arguments["initialize"].(bool) {
    return command.Initialize{VaultAlias: vaultAlias, VaultPath: vaultPath}
  } else {
    return nil
  }
}

func vaultPath(arguments map[string]interface {}) string {
  if vaultPath, hasValue := arguments["--path"].(string); hasValue {
    return vaultPath
  } else {
    return ""
  }
}

func vaultAlias(arguments map[string]interface {}) string {
  if vaultAlias, hasValue := arguments["--alias"].(string); hasValue {
    return vaultAlias
  } else {
    return ""
  }
}
