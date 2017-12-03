package path

import (
  "path/filepath"
  "os"
  "os/user"
  "io/ioutil"
  "encoding/json"
  "errors"
)

type config struct {
  Path string
}

func Get() (string, error) {
  if configJSON, err := ioutil.ReadFile(configurationPath()); os.IsNotExist(err) {
    return "", errors.New("Vault not found! Create or specify one.")
  } else {
    var conf config
    if err := json.Unmarshal(configJSON, &conf); err == nil {
      return conf.Path, nil
    } else {
      return "", err
    }
  }
}

func Store(vaultPath string) string {
  configurationPath := configurationPath()

  if configJSON, err := json.Marshal(config{vaultPath}); err != nil {
    panic(err)
  } else if err := ioutil.WriteFile(configurationPath, configJSON, 0600); err != nil {
    panic(err)
  }

  return configurationPath
}

func configurationPath() string {
  currentUser, _ := user.Current()
  currentUserHome := currentUser.HomeDir
  return filepath.Join(currentUserHome, "/.secrets.conf.json")
}
