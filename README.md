# example-CICD-tool

This is example command-line tool with configured CI/CD.


## Installation

Download binary for your OS manually from
[releases](https://github.com/salamandra19/example-CICD-tool/releases) or run
this to install the latest version:

```sh
curl -sfL $(curl -s https://api.github.com/repos/salamandra19/example-CICD-tool/releases/latest | grep -i /gh-make-labels-$(uname -s)-$(uname -m)\" | cut -d\" -f4) | sudo install /dev/stdin /usr/local/bin/example-CICD-tool
```


## Usage

Just run it and it'll greet whole world.

```
$ example-CICD-tool
hello world
$
```
