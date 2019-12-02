package models

import (
	"strconv"

	myasthurts "github.com/lab259/go-my-ast-hurts"
)

type (
	CtxPkg struct {
		Pkg   *myasthurts.Package
		Alias string
	}

	CtxImports struct {
		Pkg          *myasthurts.Package
		Imports      map[string]*CtxPkg
		importsAlias map[string]string
	}
)

func NewCtxImports(pkg *myasthurts.Package) *CtxImports {
	return &CtxImports{
		Pkg:          pkg,
		Imports:      make(map[string]*CtxPkg),
		importsAlias: make(map[string]string),
	}
}

func (ctx *CtxImports) AddImportPkg(pkg *myasthurts.Package) *CtxPkg {
	if pkg.RealPath == ctx.Pkg.RealPath {
		return &CtxPkg{
			Pkg:   pkg,
			Alias: pkg.Name,
		}
	}
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
		return ctx.AddImportPkg(pkg)
	}
	return ctxPkg
}

func (ctx *CtxImports) Ref(refType myasthurts.RefType) string {
	pkg := refType.Pkg()
	if pkg.RealPath == ctx.Pkg.RealPath || pkg.Name == "builtin" {
		return refType.Name()
	}
	ctxPkg, ok := ctx.Imports[pkg.ImportPath]
	if !ok {
		return pkg.Name + "." + refType.Name()
	}
	if ctxPkg.Alias == "." || ctxPkg.Alias == "-" {
		return refType.Name()
	}
	return ctxPkg.Alias + "." + refType.Name()
}

func (ctxPkg *CtxPkg) Ref(pkg *myasthurts.Package, typeName string) string {
	if ctxPkg.Alias == "." || (pkg != nil && pkg.RealPath == ctxPkg.Pkg.RealPath) {
		return typeName
	}
	return ctxPkg.Alias + "." + typeName
}
