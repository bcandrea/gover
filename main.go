package main

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// The Version of gover.
const Version = "0.1.0"

// GetVersion returns the value of a constant or variable named "Version" defined in
// the package with a given name.
func GetVersion(pkg string) (string, error) {
	pkgDir := filepath.Join(os.Getenv("GOPATH"), "src", pkg)
	pkgName := filepath.Base(pkg)

	if _, err := os.Stat(pkgDir); os.IsNotExist(err) {
		return "", err
	}

	fset := token.NewFileSet()
	packages, err := parser.ParseDir(fset, pkgDir, nil, 0)
	if err != nil {
		return "", nil
	}

	var pkgAst *ast.Package
	if p, found := packages["main"]; found {
		pkgAst = p
	} else {
		if p, found := packages[pkgName]; found {
			pkgAst = p
		}
	}
	if pkgAst == nil {
		return "", fmt.Errorf("cannot find package main or %s in %s", pkgName, pkgDir)
	}

	tree := ast.MergePackageFiles(pkgAst, 0)

	for _, decl := range tree.Decls {
		switch decl.(type) {
		case *ast.GenDecl:
		default:
			continue
		}
		for _, spec := range decl.(*ast.GenDecl).Specs {
			switch spec.(type) {
			case *ast.ValueSpec:
			default:
				continue
			}
			for i, name := range spec.(*ast.ValueSpec).Names {
				if name.Name != "Version" {
					continue
				}
				v := spec.(*ast.ValueSpec).Values[i]
				switch v.(type) {
				case *ast.BasicLit:
				default:
					return "", errors.New("the Version object should be a basic literal")
				}
				ver := strings.Trim(v.(*ast.BasicLit).Value, "\"")
				return ver, nil
			}
		}
	}
	return "", fmt.Errorf("no Version object in package %s", pkg)
}

func main() {
	showVersion := flag.Bool("v", false, "display version")
	flag.Parse()

	if *showVersion {
		fmt.Println("gover version", Version)
		os.Exit(0)
	}

	if len(flag.Args()) < 1 {
		fmt.Fprintln(os.Stderr, "usage: gover <package>")
		os.Exit(1)
	}
	ver, err := GetVersion(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, "gover:", err)
		os.Exit(2)
	}
	fmt.Println(ver)
}
