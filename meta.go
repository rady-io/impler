package main

import (
	"go/ast"
	"go/types"
	"strings"
)

const (
	LF = "\n"
)

const (
	HttpService Annotation = "@HttpService"
	Monitor                = "@Monitor"
	Data                   = "@Data"
	Sync                   = "@Sync"
	Proxy                  = "@Proxy"
)

const (
	TypeDecl   DeclType = "TypeDecl"
	ImportDecl          = "ImportDecl"
	ConstDecl           = "ConstDecl"
	VarDecl             = "VarDecl"
	FuncDecl            = "FuncDecl"
)

type (
	Annotation string
	DeclType   string
	DeclMap    map[Annotation]*Decls
)

type (
	Decls struct {
		Types   []*Type
		Imports []*Import
		Consts  []*Const
		Vars    []*Var
		Funcs   []*Func
	}

	// GenDecl
	Type struct {
		*ast.TypeSpec
		CombinedComments string
		Object           types.Object
		UnderlineType    types.Type
		FieldComments    map[string]string
	}

	Import struct {
		*ast.ImportSpec
		CombinedComments string
	}

	Const struct {
		*ast.ValueSpec
		CombinedComments string
	}

	Var struct {
		*ast.ValueSpec
		CombinedComments string
	}

	// FuncDecl
	Func struct {
		*ast.FuncDecl
		Comments string
	}

	// Field
	Field struct {
		*ast.Field
		Comments string
	}
)

type (
	// TypeSepc
	Struct struct {
		*Type
	}

	Interface struct {
		*Type
	}

	Map struct {
		*Type
	}

	Slice struct {
		*Type
	}

	Other struct {
		*Type
	}
)

func MakeDeclMap() DeclMap {
	return DeclMap{
		HttpService: NewDecls(TypeDecl),
		Monitor:     NewDecls(TypeDecl),
		Data:        NewDecls(TypeDecl),
		Sync:        NewDecls(TypeDecl),
		Proxy:       NewDecls(FuncDecl),
	}
}

func NewDecls(declTypes ...DeclType) (decls *Decls) {
	decls = new(Decls)
	for _, declType := range declTypes {
		switch declType {
		case TypeDecl:
			decls.Types = make([]*Type, 0)
		case ImportDecl:
			decls.Imports = make([]*Import, 0)
		case ConstDecl:
			decls.Consts = make([]*Const, 0)
		case VarDecl:
			decls.Vars = make([]*Var, 0)
		case FuncDecl:
			decls.Funcs = make([]*Func, 0)
		}
	}
	return
}

func (declMap DeclMap) resolveComments(commentMap ast.CommentMap, pkg *types.Package) (err error) {

	return
}

func combineComments(typeSpecComments, interfaceComments string) (comments string) {
	typeSpecComments = strings.Trim(typeSpecComments, LF)
	interfaceComments = strings.Trim(interfaceComments, LF)
	comments = typeSpecComments + LF + interfaceComments
	return
}
