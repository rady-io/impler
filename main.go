package main

import (
	. "github.com/rady-io/impler/log"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

//go:generate ./impler

func main() {
	var err error
	var declMap *DeclCtrl
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
			declMap = NewDeclCtrl()
			err = declMap.resolveComments(ast.NewCommentMap(pkg.Fset, file, file.Comments), pkg)
			if err != nil {
				Log.Fatal(err.Error())
			}
		}
	}
}
