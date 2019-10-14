package path

import (
  "os"
  "io/ioutil"
  "encoding/json"
  "errors"
)

type Config struct {
  Path string
  Alias string
}

func Configurations(path string) ([]Config, error) {
  if configJSON, err := ioutil.ReadFile(path); os.IsNotExist(err) {
    return make([]Config, 0), errors.New("Vault is not configured!")
  } else {
    var conf []Config
    if err := json.Unmarshal(configJSON, &conf); err == nil {
      return conf, nil
    } else {
      return make([]Config, 0), err
    }
  }
}

func FindByAlias(configs []Config, alias string) *Config {
  for _, config := range configs {
    if config.Alias == alias {
      return &config
    }
  }

  return nil
}

