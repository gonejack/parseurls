# parseurls
Parse urls from plain text or stdin

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gonejack/parseurls)
![Build](https://github.com/gonejack/parseurls/actions/workflows/go.yml/badge.svg)
[![GitHub license](https://img.shields.io/github/license/gonejack/parseurls.svg?color=blue)](LICENSE)

### Install
```shell
> go get github.com/gonejack/parseurls
```

### Usage
```shell
> echo "text..." | parseurls
```
```shell
> parseurls *.txt
```
```
Usage:
  parseurls [*] [flags]

Flags:
  -v, --verbose   verbose
  -h, --help      help for parseurls
```
