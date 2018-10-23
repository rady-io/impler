package main

import (
	"fmt"
	. "github.com/rady-io/impler/log"
	"go/ast"
	"go/build"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
)

var (
	pkgInfo  *build.Package
	fileSet  = token.NewFileSet()
	fileInfo = &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}
	sources     = make([]*ast.File, 0)
	commentMaps = make([]ast.CommentMap, 0)
)

func main() {
	var err error
	pkgInfo, err = build.ImportDir(".", 0)
	if err != nil {
		Log.Fatal(err)
	}
	for _, source := range pkgInfo.GoFiles {
		Log.Debug(source)
		file, err := parser.ParseFile(fileSet, source, nil, parser.ParseComments)
		if err != nil {
			//Log.Fatal(err.Error())
		}
		sources = append(sources, file)
		commentMaps = append(commentMaps, ast.NewCommentMap(fileSet, file, file.Comments))
	}
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check(pkgInfo.Name, fileSet, sources, fileInfo)
	fmt.Println(pkg.Name())
	fmt.Println(pkg.Scope().Lookup("Service"))
	fmt.Println(pkg.Imports())
	fmt.Println(commentMaps)
}
