package main

import (
	"go/ast"
	"strings"
)

const (
	LF = "\n"
)

type (
	Package struct {
	}

	// GenDecl
	Type struct {
		*ast.TypeSpec
		CombinedComments string
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

func combineComments(typeSpecComments, interfaceComments string) (comments string) {
	typeSpecComments = strings.Trim(typeSpecComments, LF)
	interfaceComments = strings.Trim(interfaceComments, LF)
	comments = typeSpecComments + LF + interfaceComments
	return
}