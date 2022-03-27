# GOST KeePassCLI
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/gostkp/blob/master/LICENSE.md) 
[![GitHub downloads](https://img.shields.io/github/downloads/pedroalbanese/gostkp/total.svg?logo=github&logoColor=white)](https://github.com/pedroalbanese/gostkp/releases)
[![GoDoc](https://godoc.org/github.com/pedroalbanese/gostkp?status.png)](http://godoc.org/github.com/pedroalbanese/gostkp)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/gostkp)](https://goreportcard.com/report/github.com/pedroalbanese/gostkp)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pedroalbanese/gostkp)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/gostkp)](https://github.com/pedroalbanese/gostkp/releases)

This project is a reimplementation of [kpcli](http://kpcli.sourceforge.net/) with a few additional features thrown in. It provides a shell-like interface for navigating a KeePass-like database and manipulating entries. 

Features, such as `search` have been added.

Roll of Algorithms:
- GOST R 34.11-2012 Стрибог (Streebog) hash function (RFC 6986)
- GOST R 34.12-2015 128-bit block cipher Кузнечик (Kuznechik) (RFC 7801)
- GOST R 34.12-2015 64-bit block cipher Магма (Magma)

## Usage

### Command-line
```
> ./kp -help
Usage of ./kp:
  -db string
        the db to open
  -key string
        a key file to use to unlock the db
  -n string
        execute a given command and exit
  -version
        print version and exit
```

### Program Shell
```
/ > help

Commands:
  attach       attach <get|show|delete> <entry> <filesystem location>
  cd           cd <path>
  clear        clear the screen
  edit         edit <entry>
  exit         exit the program
  help         display help
  ls           ls [path]
  mkdir        mkdir <group name>
  mv           mv <source> <destination>
  new          new <path>
  pwd          pwd
  rm           rm <entry>
  save         save
  saveas       saveas <file.kdb>
  search       search <term>
  show         show [-f] <entry>
  version      version
  xk           xk <entry>
  xp           xp <entry>
  xu           xu
  xw           xw
  xx           xx
```
Running a command with the argument `help` will display a more detailed usage message
```
/ > attach help

manages the attachment for a given entry

Commands:
  create       attach create <entry> <name> <filesystem location>
  details      attach details <entry>
  get          attach get <entry> <filesystem location> // ('-' for stdout)
```

## Overview
There are two main components, the shell and the libraries that interact with the database directly.  The shell interfaces with the database through those abstractions.  The shell works by having individual files for each command which are strung together in `main.go`.

## License

This project is licensed under the ISC License.
#### Copyright (c) 2020-2022 Pedro F. Albanese - ALBANESE Research Lab.
