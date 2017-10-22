package path

import (
  "path/filepath"
  "os"
  "os/user"
  "io/ioutil"
  "encoding/json"
  "fmt"
  "github.com/jarmo/secrets/input"
)

type config struct {
  Path string
}

type dropBoxConfig struct {
  Personal config
  Business config
}

func Get() string {
  currentUser, _ := user.Current()
  currentUserHome := currentUser.HomeDir

  secretsConfigurationPath := filepath.Join(currentUserHome, "/.secrets.conf.json")
  if configJSON, err := ioutil.ReadFile(secretsConfigurationPath); os.IsNotExist(err) {
    if path := dropBoxPath(currentUserHome); path != "" {
      return filepath.Join(path, ".secrets.json")
    }
  } else {
    var conf config
    if err := json.Unmarshal(configJSON, &conf); err == nil {
      return conf.Path
    }
  }

  fmt.Println("Vault was not found!")
  vaultPath := input.Ask("Enter vault absolute path: ")
  fmt.Println()

  return vaultPath
}

func dropBoxPath(currentUserHome string) string {
  paths := []string {
    filepath.Join(os.Getenv("APPDATA"), "/Dropbox/info.json"),
    filepath.Join(os.Getenv("LOCALAPPDATA"), "/Dropbox/info.json"),
    filepath.Join(currentUserHome, "/.dropbox/info.json")}

  for _, path := range paths {
    if configJSON, err := ioutil.ReadFile(path); os.IsNotExist(err) {
      continue
    } else {
      var conf dropBoxConfig
      if err := json.Unmarshal(configJSON, &conf); err == nil {
        if conf.Business.Path != "" {
          return conf.Business.Path
        } else if conf.Personal.Path != "" {
          return conf.Personal.Path
        }
      }
    }
  }

  return ""
}
