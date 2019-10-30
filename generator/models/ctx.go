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

	Ctx struct {
		Pkg          *myasthurts.Package
		Imports      map[string]*CtxPkg
		importsAlias map[string]string
		count        int
	}
)

func NewCtx(pkgs ...*myasthurts.Package) *Ctx {
	ctx := &Ctx{
		Pkg:          pkgs[0],
		Imports:      make(map[string]*CtxPkg),
		importsAlias: make(map[string]string),
	}
	for _, pkg := range pkgs {
		ctx.addImportPkg(pkg).Alias = "-"
	}
	return ctx
}

func (ctx *Ctx) addImportPkg(pkg *myasthurts.Package) *CtxPkg {
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

func (ctx *Ctx) AddRefType(refType myasthurts.RefType) *CtxPkg {
	pkg := refType.Pkg()
	ctxPkg, ok := ctx.Imports[pkg.ImportPath]
	if !ok {
		return ctx.addImportPkg(pkg)
	}
	return ctxPkg
}
