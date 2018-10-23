package main

import "go/ast"

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
