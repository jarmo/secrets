# secrets

**Secure** and simple passwords manager written in [Go](https://golang.org/). It aims to be *NYAPM* (Not Yet Another Password Manager), but tries to be different from others by following UNIX philosophy of doing only one thing and doing it well.

## Features

* stores your secrets encrypted at rest;
* secrets can be anything from passwords, 2FA backup codes, diary entries to private keys;
* does not leak count nor nature of your secrets;
* uses an alternative easy to use secure cryptography provided by [libsodium](https://download.libsodium.org/doc/) and [Argon2id](https://www.cryptolux.org/images/0/0d/Argon2.pdf);
* supports multiple vaults with different passwords;
* has [CLI](https://github.com/jarmo/secrets-cli) and [Web](https://github.com/jarmo/secrets-web) interface pre-built binaries for macOS, Linux and Windows, but can be compiled for many other platforms too due to usage of underlying Go language;
* may be used as an independent Go library.

### Anti-Features

* does not sync your secrets to any cloud - you have [complete control](https://palant.de/2019/03/18/should-you-be-concerned-about-lastpass-uploading-your-passwords-to-its-server/) over them;
* does not allow to recover any passwords when vault password has been forgotten - there's no built-in backdoor;
* does not generate any passwords - use [proper tools](https://linux.die.net/man/1/pwgen) for that, but avoid [improper ones](http://seclists.org/oss-sec/2018/q1/11);
* does not auto-fill any passwords anywhere (you [don't want](https://freedom-to-tinker.com/2017/12/27/no-boundaries-for-user-identities-web-trackers-exploit-browser-login-managers/) that anyway) - it's up to you how you will fill your passwords;
* does not have any mobile apps nor browser plugins - less possible attack vectors;
* does not [remove already existing features](https://discussions.agilebits.com/discussion/105305/standalone-local-vault-option-gone) - always possibility to create your own fork since it is an open-source software and will be like that.

## Is it secure?

**Yes**, as long as its underlying cryptography is not broken. However, there are no 100% secure systems and there's no way to guarantee that. All in all, I'd say that using this is more secure than using any SaaS as a password manager because everything is under your control. The most secure system is not a software itself, but it's how and where you use it.

## Usage

It is possible to use secrets from [command line](https://github.com/jarmo/secrets-cli), as a [self-hosted web application](https://github.com/jarmo/secrets-web)
or as a library.

## Development

Retrieve dependencies and run tests

```
git clone https://github.com/jarmo/secrets.git
cd secrets
make
```

## Background Story

I've used [LastPass](https://www.lastpass.com/) and [mitro](http://www.mitro.co/) in the past to store my secrets, but didn't feel too secure with either of them due to security vulnerabilities and/or one of them being shut down. I've got enough of switching between different managers and decided to write my own. I did write a version of **secrets** in Ruby a few years ago, but decided to give Go a try due to its portability features and here's the result. I've also decided to use a cryptographic library called libsodium, which is secure and has an easy API for avoiding making stupid mistakes.
