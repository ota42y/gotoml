[![Circle CI](https://circleci.com/gh/ota42y/gotoml.svg?style=svg)](https://circleci.com/gh/ota42y/gotoml)
[![Coverage Status](https://coveralls.io/repos/ota42y/gotoml/badge.svg?branch=master&service=github)](https://coveralls.io/github/ota42y/gotoml?branch=master)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/ota42y/gotoml/blob/master/LICENSE)

# gotoml

## Description
Convert Toml to golang struct.  
Inspired by [gojson](https://github.com/ChimeraCoder/gojson)

This tool isn't release, so many bug and not support all data type.

## Usage

```bash
cat $GOPATH/src/github.com/ota42y/gotoml/example/normal.toml | gotoml
```

## Install

To install, use `go get`:

```bash
$ go get github.com/ota42y/gotoml/gotoml
```

## Contribution

1. Fork ([https://github.com/ota42y/gotoml/fork](https://github.com/ota42y/gotoml/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[ota42y](https://github.com/ota42y)

## Licence
MIT: [https://github.com/ota42y/gotoml/blob/master/LICENSE](https://github.com/ota42y/gotoml/blob/master/LICENSE)
