# secrets
[![CircleCI](https://circleci.com/gh/jarmo/secrets.svg?style=svg)](https://circleci.com/gh/jarmo/secrets)

**Secure** and simple passwords manager written in [Go](https://golang.org/). It aims to be *NYAPM* (Not Yet Another Password Manager), but tries to be different from others by following UNIX philosophy of doing only one thing and doing it well.

## Features

* stores your secrets encrypted at rest;
* secrets can be anything from passwords, 2FA backup codes, diary entries to private keys;
* does not leak count nor nature of your secrets;
* uses an alternative easy to use secure cryptography provided by [libsodium](https://download.libsodium.org/doc/) and [Argon2id](https://www.cryptolux.org/images/0/0d/Argon2.pdf);
* supports multiple vaults with different passwords;
* has [CLI](https://en.wikipedia.org/wiki/Command-line_interface) interface pre-built binaries for macOS, Linux and Windows, but can be compiled for many other platforms too due to usage of underlying Go language;
* may be used as an independent Go library.

### Anti-Features

* does not sync your secrets to any cloud - you have [complete control](https://palant.de/2019/03/18/should-you-be-concerned-about-lastpass-uploading-your-passwords-to-its-server/) over them;
* does not allow to recover any passwords when vault password has been forgotten - there's no built-in backdoor;
* does not generate any passwords - use [proper tools](https://linux.die.net/man/1/pwgen) for that, but avoid [improper ones](http://seclists.org/oss-sec/2018/q1/11);
* does not auto-fill any passwords anywhere (you [don't want](https://freedom-to-tinker.com/2017/12/27/no-boundaries-for-user-identities-web-trackers-exploit-browser-login-managers/) that anyway) - it's up to you how you will fill your passwords;
* does not have a running application when no password is needed for retrieval - less chance of your [secrets to be leaked](https://www.securityevaluators.com/casestudies/password-manager-hacking/);
* does not have any mobile apps nor browser plugins - less possible attack vectors;
* does not [remove already existing features](https://discussions.agilebits.com/discussion/105305/standalone-local-vault-option-gone) - always possibility to create your own fork since it is an open-source software and will be like that.

## Is it secure?

**Yes**, as long as its underlying cryptography is not broken. However, there are no 100% secure systems and there's no way to guarantee that. All in all, I'd say that using this is more secure than using any SaaS as a password manager because everything is under your control. The most secure system is not a software itself, but it's how and where you use it.

## Installation

Download latest binary from [releases](https://github.com/jarmo/secrets/releases), extract it and add it to somewhere in your **PATH**. That's it.

*Of course, you're free to compile your own version of binary to be 100% sure that it has not been tampered with, since this is an open-source project after all.*

## Usage

Here's an output from `secrets --help` command.

```
$ secrets COMMAND [OPTIONS]

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
  -v --version           Show version.
```

### Initializing Vault

Vault needs to be initialized if there is going to be a default vault. Otherwise specifying `--path` or `--alias` with any command is supported. Initializing vault just stores location and alias to your vault into a configuration file (supporting [XDG Base Directory standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html)):

```
$ secrets initialize --path /home/user/.secrets.json --alias main
Vault successfully configured at /home/user/.config/secrets/config.json and is ready to store your secrets!
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

## Using multiple vaults

Just append `--alias` after any command to operate against selected vault.
When `--alias` is not specified a first vault existing in configuration file will be used.

## But how do I sync vault between different devices?!

One way to sync would be to use any already existing syncing platforms like Dropbox, Microsoft OneDrive or Google Drive.
Since you can specify vault storage location then it is up to you how (or if even) you sync.

## Development

Retrieve dependencies, build and install binaries to `$GOPATH/bin/`

```
go get -u github.com/jarmo/secrets
cd $GOPATH/src/github.com/jarmo/secrets
make
make install
```


## Background Story

I've used [LastPass](https://www.lastpass.com/) and [mitro](http://www.mitro.co/) in the past to store my secrets, but didn't feel too secure with either of them due to security vulnerabilities and/or one of them being shut down. I've got enough of switching between different managers and decided to write my own. I did write a version of **secrets** in Ruby a few years ago, but decided to give Go a try due to its portability features and here's the result. I've also decided to use a cryptographic library called libsodium, which is secure and has an easy API for avoiding making stupid mistakes.
