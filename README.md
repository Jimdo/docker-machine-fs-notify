# Docker Machine FS notify

Forward file system events to a docker machine VM

[![Build Status](https://circleci.com/gh/Jimdo/docker-machine-fs-notify/tree/master.svg?style=shield)](https://circleci.com/gh/Jimdo/docker-machine-fs-notify)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg "MIT License")](https://github.com/twbs/no-carrier/blob/master/LICENSE.txt)

## Dependencies

* Docker
* make

## Building the project

```
make build
```

## Creating a new release for this project

Make sure that you are on the latest `master` branch and that you have a clean Git working directory.

```
VERSION=<version> GITHUB_TOKEN=<github-api-token> make release
```

## Usage

```
./docker-machine-fs-notify <directory> <docker-machine-name>
```

## Updating dependencies:

Start a `Go` Docker Container:
```
docker run -it -v $PWD:/go/src/github.com/Jimdo/docker-machine-fs-notify -w /go/src/github.com/Jimdo/docker-machine-fs-notify golang:1.5.1 /bin/bash
```

Install packages to `Godeps`:
```
go get github.com/tools/godep
godep restore
# ...
# Install your packages
# ...
godep save
```

## Links

* https://github.com/codekitchen/fsevents_to_vm
