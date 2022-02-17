# baldeweg/mission

A CLI to create markdown files for a wiki.

## Requirements

- [Go](https://go.dev/)
- Basic knowledge about the command line

## Getting Started

First, you need to install [Go](https://go.dev/).

Download the project archive from the [git repository](https://github.com/abaldeweg/wiki).

Inside the project directory, you can build the app with the `go build` command. If you have [GoReleaser](https://goreleaser.com/) installed, instead run `goreleaser build --snapshot --rm-dist`.

Run the command `wiki`. Depending on the OS you need to add a file extension.

The articles will be created in your working directory.

## Commands

- wiki new - Adds a new wiki
- wiki help - Shows the help

## Flags

- --path - Specify the directory where the data should be stored.
- --file - Specify the filename.
