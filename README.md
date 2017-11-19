# secrets
[![CircleCI](https://circleci.com/gh/jarmo/secrets.svg?style=svg&circle-token=596c25b873ed12dd07c3df358afbf7e0c0cdf806)](https://circleci.com/gh/jarmo/secrets)

**Secure** passwords manager written in Go. It aims to be *NYAPM* (Not Yet Another Password Manager), but tries to be different from others by using a different kind of easy to use secure cryptographic functions provided by [libsodium](https://download.libsodium.org/doc/) - you won't see any *NSA-approved* cryptography in here.

Other big difference from other solutions is that **secrets** does only one thing and does it well - it stores *secrets* encrypted at rest. It does not sync between machines, it does generate new passwords, it does not auto-fill any passwords, etc.

Secrets can be anything from passwords, 2FA backup codes, secret diary entries to private keys!

## Installation

Download latest binary from [releases](https://github.com/jarmo/secrets/releases) and add it to somewhere in your **PATH**. That's it.

*Of course, you can compile your own version of binary to be 100% sure that it has not been tampered with.*

## Usage

**secrets** has a *CLI* (command-line interface), but can be also used as a Go library if there's any need.

Here's an output from `secrets --help` command.

```
$ secrets --help
secrets

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
  -v --version         Show version.
```

### Adding a New Secret

Add your first secret will also create a vault:

```
$ secrets -a "my secret" 
Enter vault password: [enter secure passphrase and remember it]
Enter value for 'my secret':
my secret value
Added: 
[299ed462-b171-4d67-ba21-264b221d9913]
my secret
my secret value
```

Because values can have multiple lines, you can enter whatever you want. Use ctrl+d on **macOS** and **Linux** or ctrl+z on **Windows** to complete entering multi-line values.

### Listing All Secrets

```
$ secrets -l
Enter vault password: [your secure passphrase]

[299ed462-b171-4d67-ba21-264b221d9913]
my secret
my secret value
```

### Listing Specific Secrets

```
$ secrets -l "secret"
Enter vault password: [your secure passphrase]

[299ed462-b171-4d67-ba21-264b221d9913]
my secret
my secret value
```

### Editing a Secret

```
$ secrets -e 299ed462-b171-4d67-ba21-264b221d9913                                                                    
Enter vault password: [your secure passphrase]
Enter new name: different secret name
Enter new value:
different secret value
yet another secret value line
Edited: 
[299ed462-b171-4d67-ba21-264b221d9913]
different secret name
different secret value
yet another secret value line
```

### Deleting a Secret

```
$ secrets -d 299ed462-b171-4d67-ba21-264b221d9913
Enter vault password: 
Deleted: 
[299ed462-b171-4d67-ba21-264b221d9913]
different secret name
different secret value
yet another secret value line
```

## Vault Location

Vault will be created by default into your Dropbox folder if you have it installed as a **.secrets.json**. This acts as a possible backup/syncing solution when a need arises. You can however use a custom path by creating a configuration file similar to this:

```
echo '{"Path": "/home/foo/my-secrets.json"}' > ~/.secrets.conf.json
```

## Development

1. Install [dep](https://github.com/golang/dep) for dependency management.

2. Retrieve, build and install binaries to `$GOPATH/bin/`

```
go get -u github.com/jarmo/secrets
cd $GOPATH/src/github.com/jarmo/secrets
dep ensure
make
make install
```

## Background Story

I've used [LastPass](https://www.lastpass.com/) and [mitro](http://www.mitro.co/) in the past to store my secrets, but didn't feel too secure with either of them due to security vulnerabilities and/or one of them being shut down. I've got enough of switching between different managers and decided to write my own. I did write a version of **secrets** in Ruby a few years ago, but decided to give Go a try due to its portability features and here's the result. I've also decided to use a cryptographic library called libsodium.
