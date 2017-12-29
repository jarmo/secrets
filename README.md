# secrets
[![CircleCI](https://circleci.com/gh/jarmo/secrets.svg?style=svg)](https://circleci.com/gh/jarmo/secrets)

**Secure** and simple passwords manager written in [Go](https://golang.org/). It aims to be *NYAPM* (Not Yet Another Password Manager), but tries to be different from others by following UNIX philosophy of doing only one thing and doing it well.

## Features

* stores your secrets encrypted at rest;
* secrets can be anything from passwords, 2FA backup codes, diary entries to private keys;
* does not leak count nor nature of your secrets;
* uses an alternative easy to use secure cryptography provided by [libsodium](https://download.libsodium.org/doc/);
* supports multiple vaults with different passwords;
* has [CLI](https://en.wikipedia.org/wiki/Command-line_interface) interface pre-built binaries for macOS, Linux and Windows, but can be compiled for many other platforms too due to usage of underlying Go language;
* may be used as an independent Go library.

### Anti-Features

* does not sync your secrets to any cloud - you have complete control over them;
* does not generate any passwords - use [proper tools](https://linux.die.net/man/1/pwgen) for that;
* does not auto-fill any passwords anywhere - it's up to you how you will fill your passwords;
* does not have any mobile apps nor browser plugins - less chance of your secrets to be leaked;

## Is it secure?

**Yes**, as long as its underlying cryptography is not broken. However, there are no 100% secure systems and there's no way to guarantee that. All in all, I'd say that using this is more secure than using any other SaaS as a password manager because everything is under your control. The most secure system is not a software itself, but it's how and where you use it.

## Installation

Download latest binary from [releases](https://github.com/jarmo/secrets/releases) extract it and add it to somewhere in your **PATH**. That's it.

*Of course, you can compile your own version of binary to be 100% sure that it has not been tampered with.*

## Usage

Here's an output from `secrets --help` command.

```
$ secrets COMMAND [OPTIONS]

Usage:
  secrets list [FILTER] [--path=VAULT_PATH]
  secrets add NAME [--path=VAULT_PATH]
  secrets edit ID [--path=VAULT_PATH]
  secrets delete ID [--path=VAULT_PATH]
  secrets change-password [--path=VAULT_PATH]
  secrets initialize --path=VAULT_PATH

Options:
  --path VAULT_PATH      Optional vault path. Defaults to the path in configuration.
  -h --help              Show this screen.
  -v --version           Show version.
```

### Initializing Vault

Vault needs to be initialized if there is going to be a default vault. Otherwise specifying `--path` with any command is supported. Initializing vault just stores location to your vault into a configuration file:

```
$ secrets initialize --path /home/user/.secrets.json
Vault successfully configured at /home/user/.secrets.conf.json and is ready to store your secrets!
```

### Adding a New Secret

Add your first secret:

```
$ secrets add "my secret"
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
$ secrets list
Enter vault password: [your secure passphrase]

[299ed462-b171-4d67-ba21-264b221d9913]
my secret
my secret value
```

### Listing Specific Secrets

```
$ secrets list "secret"
Enter vault password: [your secure passphrase]

[299ed462-b171-4d67-ba21-264b221d9913]
my secret
my secret value
```

### Editing a Secret

```
$ secrets edit 299ed462-b171-4d67-ba21-264b221d9913
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
$ secrets delete 299ed462-b171-4d67-ba21-264b221d9913
Enter vault password: 
Deleted: 
[299ed462-b171-4d67-ba21-264b221d9913]
different secret name
different secret value
yet another secret value line
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

I've used [LastPass](https://www.lastpass.com/) and [mitro](http://www.mitro.co/) in the past to store my secrets, but didn't feel too secure with either of them due to security vulnerabilities and/or one of them being shut down. I've got enough of switching between different managers and decided to write my own. I did write a version of **secrets** in Ruby a few years ago, but decided to give Go a try due to its portability features and here's the result. I've also decided to use a cryptographic library called libsodium. I've done my best, but there's no guarantees that it's secure.
