package main

import (
	"github.com/rady-io/annotation-processor"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"sort"
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
	Validate               = "@Validate"
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
	TypeList []*Type

	DeclCtrl struct {
		declMap  DeclMap
		TypeList TypeList // struct or interface
		Imports  []*Import
		Consts   []*Const
		Vars     []*Var
		Funcs    []*Func
	}

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

func NewDeclCtrl() *DeclCtrl {
	return &DeclCtrl{
		declMap:  MakeDeclMap(),
		TypeList: make(TypeList, 0),
		Imports:  make([]*Import, 0),
		Consts:   make([]*Const, 0),
		Vars:     make([]*Var, 0),
		Funcs:    make([]*Func, 0),
	}
}

func MakeDeclMap() DeclMap {
	return DeclMap{
		HttpService: NewDecls(TypeDecl),
		Monitor:     NewDecls(TypeDecl),
		Data:        NewDecls(TypeDecl),
		Sync:        NewDecls(TypeDecl),
		Proxy:       NewDecls(FuncDecl),
		Validate:    NewDecls(TypeDecl),
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

func (ctrl *DeclCtrl) resolveComments(commentMap ast.CommentMap, pkg *packages.Package) (err error) {
	for node := range commentMap {
		switch tok := node.(type) {
		case *ast.GenDecl:
			err = ctrl.resolveGenDecl(tok, commentMap, pkg)
		case *ast.FuncDecl:

		}
		if err != nil {
			break
		}
	}

	if err == nil {
		ctrl.resolveFields(commentMap, pkg)
		err = ctrl.resolveMap(pkg)
		if err == nil {
		}
	}
	return
}

func (ctrl *DeclCtrl) resolveFields(commentMap ast.CommentMap, pkg *packages.Package) {
	if !sort.IsSorted(ctrl.TypeList) {
		sort.Sort(ctrl.TypeList) // sort
	}
	for node := range commentMap {
		if tok, ok := node.(*ast.Field); ok && ctrl.TypeList.Len() > 0 {
			index := sort.Search(ctrl.TypeList.Len(), func(i int) bool {
				return tok.Pos() > ctrl.TypeList[i].Pos() && tok.End() < ctrl.TypeList[i].End()
			})
			if index < ctrl.TypeList.Len() {
				for _, name := range tok.Names {
					ctrl.TypeList[index].FieldComments[name.Name] = tok.Doc.Text()
				}
			}
		}
	}
}

func (ctrl *DeclCtrl) resolveGenDecl(node *ast.GenDecl, commentMap ast.CommentMap, pkg *packages.Package) (err error) {
	switch node.Tok {
	case token.TYPE:
		for _, spec := range node.Specs {
			typ := spec.(*ast.TypeSpec)
			obj := pkg.TypesInfo.Defs[typ.Name]
			if obj != nil {
				newType := &Type{
					TypeSpec:         typ,
					CombinedComments: combineComments(node.Doc.Text(), typ.Doc.Text()),
					Object:           obj,
					UnderlineType:    obj.Type().Underlying(),
				}
				switch newType.UnderlineType.(type) {
				case *types.Interface:
					newType.FieldComments = make(map[string]string)
				case *types.Struct:
					newType.FieldComments = make(map[string]string)
				}
				ctrl.TypeList = append(ctrl.TypeList, newType)
			}

		}
	case token.IMPORT:
	case token.CONST:
	case token.VAR:
	}
	return
}

func (ctrl *DeclCtrl) resolveMap(pkg *packages.Package) (err error) {
	if err = ctrl.resolveTypeList(); err == nil {
	}
	return
}

func (ctrl *DeclCtrl) resolveTypeList() (err error) {
	for _, typ := range ctrl.TypeList {
		err = processor.NewProcessor(typ.CombinedComments).Scan(func(ann, key, value string) (err error) {
			if decls, ok := ctrl.declMap[Annotation(ann)]; ok {
				if decls.Types == nil {
					err = UnsupportedAnnotationForError(ann, token.TYPE.String())
				}
				if err == nil {
					decls.Types = append(decls.Types, typ)
				}
			}
			return
		})
		if err != nil {
			break
		}
	}
	return
}

func combineComments(typeSpecComments, interfaceComments string) (comments string) {
	typeSpecComments = strings.Trim(typeSpecComments, LF)
	interfaceComments = strings.Trim(interfaceComments, LF)
	comments = typeSpecComments + LF + interfaceComments
	return
}

func (list TypeList) Len() int {
	return len(list)
}

func (list TypeList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list TypeList) Less(i, j int) bool {
	return list[i].Pos() < list[j].Pos()
}
