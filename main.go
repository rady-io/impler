package main

import (
	"fmt"
	. "github.com/rady-io/impler/log"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

//go:generate go run main.go

var (
	commentMaps = make([]ast.CommentMap, 0)
)

func main() {
	var err error
	cfg := &packages.Config{Mode: packages.LoadSyntax}
	pkgs, err := packages.Load(cfg, ".")
	if err != nil {
		Log.Fatal(err.Error())
	}
	if len(pkgs) == 0 {
		Log.Fatal("there is no package")
	}
	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			commentMaps = append(commentMaps, ast.NewCommentMap(pkg.Fset, file, file.Comments))
		}
		fmt.Println(commentMaps)
	}
}
