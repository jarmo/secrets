package path

import (
  "fmt"
  "path/filepath"
  "os"
  "os/user"
  "io/ioutil"
  "encoding/json"
  "errors"
  "github.com/pinzolo/xdgdir"
)

type config struct {
  Path string
  Alias string
}

func Get(alias string) (string, error) {
  if configs, err := configurations(configurationPath()); err == nil {
    if confByAlias := findByAlias(configs, alias); confByAlias != nil {
      return confByAlias.Path, nil
    } else {
      return configs[0].Path, nil
    }
  } else {
    return "", err
  }
}

func Store(vaultPath string, vaultAlias string) string {
  configurationPath := configurationPath()
  conf, _ := configurations(configurationPath)
  conf = append(conf, config{Path: vaultPath, Alias: vaultAlias})

  if configJSON, err := json.MarshalIndent(conf, "", " "); err != nil {
    panic(err)
  } else if err := ioutil.WriteFile(configurationPath, configJSON, 0600); err != nil {
    panic(err)
  }

  return configurationPath
}

func configurations(path string) ([]config, error) {
  if configJSON, err := ioutil.ReadFile(path); os.IsNotExist(err) {
    return make([]config, 0), errors.New("Vault not found! Create one with initialize command or specify with --path or --alias switches.")
  } else {
    var conf []config
    if err := json.Unmarshal(configJSON, &conf); err == nil {
      return conf, nil
    } else {
      return make([]config, 0), err
    }
  }
}

func findByAlias(configs []config, alias string) *config {
  for _, config := range configs {
    if config.Alias == alias {
      return &config
    }
  }

  return nil
}

func configurationPath() string {
  xdgApp := xdgdir.NewApp("secrets")
  xdgConfigurationFilePath, err := xdgApp.ConfigFile("config.json")
  if err != nil {
    panic(err)
  }

  currentUser, _ := user.Current()
  currentUserHome := currentUser.HomeDir
  deprecatedConfigurationPath := filepath.Join(currentUserHome, "/.secrets.conf.json")

  if _, err := os.Stat(deprecatedConfigurationPath); err == nil {
    xdgConfigurationDir, err := xdgApp.ConfigDir()

    if err != nil {
      panic(err)
    }

    fmt.Println(fmt.Sprintf(`
[WARN] current secrets configuration will stop working in the future major version!
[WARN] XDG Base Directory (https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html) needs to be used!

Fix it with the following commands:
mkdir -p %s
mv %s %s
`, xdgConfigurationDir, deprecatedConfigurationPath, xdgConfigurationFilePath))
    return deprecatedConfigurationPath
  } else {
    return xdgConfigurationFilePath
  }
}
