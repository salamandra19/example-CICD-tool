# example-CICD-tool
[![CircleCI](https://circleci.com/gh/salamandra19/example-CICD-tool.svg?style=svg)](https://circleci.com/gh/salamandra19/example-CICD-tool) [![Coverage Status](https://coveralls.io/repos/github/salamandra19/example-CICD-tool/badge.svg?branch=master)](https://coveralls.io/github/salamandra19/example-CICD-tool?branch=master)

This is example command-line tool with configured CI/CD.


## Installation

Download binary for your OS manually from
[releases](https://github.com/salamandra19/example-CICD-tool/releases) or run
this to install the latest version:

```sh
curl -sfL $(curl -s https://api.github.com/repos/salamandra19/example-CICD-tool/releases/latest | grep -i /example-CICD-tool-$(uname -s)-$(uname -m)\" | cut -d\" -f4) | sudo install /dev/stdin /usr/local/bin/example-CICD-tool
```


## Usage

Just run it and it'll greet whole world.

```
$ example-CICD-tool
hello world
$
```
