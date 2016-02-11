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

// Version constant for gover.
const Version = "0.1.1"

// packageDir returns the actual directory on the file system where a package is located.
// It accepts a standard package name (e.g. github.com/bcandrea/gover) or a
// relative path (e.g. ./mypackage) as input.
func packageDir(pkg string) (string, error) {
	d := filepath.Join(os.Getenv("GOPATH"), "src", pkg)
	if strings.HasPrefix(pkg, ".") {
		var err error
		if d, err = filepath.Abs(pkg); err != nil {
			return d, err
		}
	}
	// this check covers non-existing directories
	if _, err := os.Stat(d); err != nil {
		return d, err
	}
	return d, nil
}

// syntaxTree retrieves the AST for the given package by merging all its files
// and constructing a global syntax tree.
func syntaxTree(pkgDir string) (*ast.File, error) {
	fset := token.NewFileSet()
	packages, err := parser.ParseDir(fset, pkgDir, nil, 0)
	if err != nil {
		return nil, err
	}

	var pkgAst *ast.Package
	pkgName := filepath.Base(pkgDir)
	if p, found := packages["main"]; found {
		pkgAst = p
	} else {
		if p, found := packages[pkgName]; found {
			pkgAst = p
		}
	}
	if pkgAst == nil {
		return nil, fmt.Errorf("cannot find package main or %s in %s", pkgName, pkgDir)
	}

	return ast.MergePackageFiles(pkgAst, 0), nil
}

// versionFromAST retrieves the version from a constant or variable defined in
// the abstract syntax tree.
func versionFromAST(tree *ast.File) (string, error) {
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
	return "", errors.New("no Version declaration found")
}

// GetVersion returns the value of a constant or variable named "Version" defined in
// the package with a given name. It accepts a standard package name
// (e.g. github.com/bcandrea/gover) or a relative path (e.g. ./mypackage) as input.
func GetVersion(pkg string) (string, error) {
	pkgDir, err := packageDir(pkg)
	if err != nil {
		return "", err
	}

	tree, err := syntaxTree(pkgDir)
	if err != nil {
		return "", err
	}

	return versionFromAST(tree)
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
