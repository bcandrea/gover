# gover - a version detector for Go packages
[![Build Status](https://travis-ci.org/bcandrea/gover.svg?branch=master)](https://travis-ci.org/bcandrea/gover)
[![Coverage Status](https://coveralls.io/repos/github/bcandrea/gover/badge.svg?branch=master)](https://coveralls.io/github/bcandrea/gover?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/bcandrea/gover)](https://goreportcard.com/report/github.com/bcandrea/gover)

A simple tool that outputs the value of the constant (or variable) `Version`
defined in a given Go package.

## Installation

    $ go get github.com/bcandrea/gover

## Usage

Example:

    $ go get github.com/hashicorp/consul
    $ gover github.com/hashicorp/consul
    0.5.2

If supplied with a relative path (i.e. starting with "."), `gover` will treat
the argument as a directory location and not a Go package name.

This is particularly useful during development, e.g. when the version number
needs to be included in the compilation output:

    $ go build -o my-tool_$(gover .)
