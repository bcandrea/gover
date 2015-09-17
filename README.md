# gover - a version detector for Go packages

A simple tool that outputs the value of the constant `Version` in a given Go
package. If no version is defined, nothing is printed and the tool exit with
a nonzero status code.

Example:

    $ go get github.com/hashicorp/consul
    $ gover github.com/hashicorp/consul
    0.5.2

The tool is particularly useful during development, e.g. when the version number
needs to be included in the compilation output:

    $ go build -o my-tool_$(gover github.com/me/my-tool)
