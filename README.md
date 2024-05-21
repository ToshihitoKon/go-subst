# go-subst

`go-subst` is a CLI tool that embeds environment variables into text files, used for similar purposes as `envsubst`, but with enhanced safety.

## Features

- Support Go Template Custom Functions: `env` and `must_env`. 
    - ``{{ env `ENVVAR` }}``: Fetches the environment variable `ENVVAR`. If it is not set, the output remains unchanged. 
    - ``{{ must_env `ENVVAR` }}``: If the environment variable is not set, go-subst will return an error.

## Installation

Download binary from GitHub releases.

or

```bash
go install github.com/ToshihitoKon/go-subst/cmd/go-subst@latest
```

## Usage

```bash
export SOME_ENV=env-value
cat template | go-subst > output-file
```

## Examples

Template file `example.service.tmpl`

```service
[Unit]
Description={{ env `SERVICE_DESCRIPTION` }}
After=network.service

[Service]
Type=simple
ExecStart={{ must_env `SERVICE_BIN_PATH`}}
ExecStop=/bin/kill -WINCH ${MAINPID}
Restart=always

[Install]
WantedBy=multi-user.target
```

Generate service file using go-subst.

```shell
 $ export SERVICE_DESCRIPTION="This is example service." SERVICE_BIN_PATH="/usr/local/bin/example"
 $ cat example.service.template | go-subst
[Unit]
Description=This is example service.
After=network.service

[Service]
Type=simple
ExecStart=/usr/local/bin/example
ExecStop=/bin/kill -WINCH ${MAINPID}
Restart=always

[Install]
WantedBy=multi-user.target
```

## Lisence

This project is licensed under the MIT License, see the LICENSE and CREDITS files for details.

## Motivation

Inspired by and using [kayac/go-config](https://github.com/kayac/go-config) as a package.
