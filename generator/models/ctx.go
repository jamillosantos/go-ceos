package models

import (
	"strconv"

	"github.com/jamillosantos/go-ceous/generator/reporters"
	myasthurts "github.com/lab259/go-my-ast-hurts"
)

type (
	CtxPkg struct {
		Pkg   *myasthurts.Package
		Alias string
	}

	CtxImports struct {
		Imports      map[string]*CtxPkg
		importsAlias map[string]string
	}

	Ctx struct {
		Pkg           *myasthurts.Package
		Reporter      reporters.Reporter
		Models        map[string]*Model
		Imports       CtxImports
		ModelsImports CtxImports
		count         int
	}
)

func newCtxImports() CtxImports {
	return CtxImports{
		Imports:      make(map[string]*CtxPkg),
		importsAlias: make(map[string]string),
	}
}

func (ctx *CtxImports) addImportPkg(pkg *myasthurts.Package) *CtxPkg {
	i := 0
	pkgName := pkg.Name
	for {
		if i > 0 {
			pkgName = pkg.Name + strconv.Itoa(i)
		}
		_, ok := ctx.importsAlias[pkgName]
		if ok {
			i++
			continue
		}
		ctx.importsAlias[pkgName] = pkg.ImportPath
		break
	}

	ctxPkg := &CtxPkg{
		Pkg:   pkg,
		Alias: pkgName,
	}

	ctx.Imports[pkg.ImportPath] = ctxPkg
	return ctxPkg
}

func (ctx *CtxImports) AddRefType(refType myasthurts.RefType) *CtxPkg {
	pkg := refType.Pkg()
	ctxPkg, ok := ctx.Imports[pkg.ImportPath]
	if !ok {
		return ctx.addImportPkg(pkg)
	}
	return ctxPkg
}

func NewCtx(reporter reporters.Reporter, pkgs ...*myasthurts.Package) *Ctx {
	ctx := &Ctx{
		Reporter:      reporter,
		Pkg:           pkgs[0],
		Models:        make(map[string]*Model, 0),
		Imports:       newCtxImports(),
		ModelsImports: newCtxImports(),
	}
	for _, pkg := range pkgs {
		ctx.Imports.addImportPkg(pkg).Alias = "-"
	}
	return ctx
}

func (ctx *Ctx) EnsureModel(name string) *Model {
	m, ok := ctx.Models[name]
	if ok {
		return m
	}
	m = NewModel(name)
	ctx.Models[name] = m
	return m
}
