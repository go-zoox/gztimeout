# Timeout CLI - A timeout command

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/timeout-cli)](https://pkg.go.dev/github.com/go-zoox/timeout-cli)
[![Build Status](https://github.com/go-zoox/timeout-cli/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/timeout-cli/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/timeout-cli)](https://goreportcard.com/report/github.com/go-zoox/timeout-cli)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/timeout-cli/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/timeout-cli?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/timeout-cli.svg)](https://github.com/go-zoox/timeout-cli/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/timeout-cli.svg?label=Release)](https://github.com/go-zoox/timeout-cli/tags)

## Installation
To install the package, run:
```bash
go install github.com/go-zoox/timeout-cli@latest
```

## Getting Started

```bash
timeout-cli -a 3 -e "sleep 10" -m "build timeout 3s"
```

## License
GoZoox is released under the [MIT License](./LICENSE).
