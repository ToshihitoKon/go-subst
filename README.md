# go-subst

## Features

Support `env`, `must_env`, and `json_escape`. 

## Installation

Download binary from GitHub releases.

or

```bash
go install github.com/ToshihitoKon/go-subst
```

## Usage

```bash
export SOME_ENV=env-value
cat template | go-subst > output-file
```

## Examples

## Lisence

This project is licensed under the MIT License, see the LICENSE.txt and CREDIT files for details.

## Motivation

Inspired by and using [kayac/go-config](https://github.com/kayac/go-config) as a package.
