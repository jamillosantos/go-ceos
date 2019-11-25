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
		Pkg          *myasthurts.Package
		Imports      map[string]*CtxPkg
		importsAlias map[string]string
	}

	Ctx struct {
		InputPkg      *myasthurts.Package
		InputPkgCtx   *CtxPkg
		OutputPkg     *myasthurts.Package
		OutputPkgCtx  *CtxPkg
		Reporter      reporters.Reporter
		Models        map[string]*Model
		Imports       CtxImports
		ModelsImports CtxImports
		count         int
	}
)

func newCtxImports(pkg *myasthurts.Package) CtxImports {
	return CtxImports{
		Pkg:          pkg,
		Imports:      make(map[string]*CtxPkg),
		importsAlias: make(map[string]string),
	}
}

func (ctx *CtxImports) addImportPkg(pkg *myasthurts.Package) *CtxPkg {
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
		return ctx.addImportPkg(pkg)
	}
	return ctxPkg
}

func (ctx *CtxImports) Ref(refType myasthurts.RefType) string {
	pkg := refType.Pkg()
	if pkg.RealPath == ctx.Pkg.RealPath {
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

func NewCtx(reporter reporters.Reporter, inputPackage, outputPackage *myasthurts.Package, pkgs ...*myasthurts.Package) *Ctx {
	ctx := &Ctx{
		Reporter:      reporter,
		InputPkg:      inputPackage,
		OutputPkg:     outputPackage,
		Models:        make(map[string]*Model, 0),
		Imports:       newCtxImports(outputPackage),
		ModelsImports: newCtxImports(inputPackage),
	}
	inputPkg := ctx.ModelsImports.addImportPkg(inputPackage)
	ctx.InputPkgCtx = &CtxPkg{
		Pkg:   inputPackage,
		Alias: inputPackage.Name,
	}
	inputPkg.Alias = "-"
	for _, pkg := range pkgs {
		ctx.ModelsImports.addImportPkg(pkg).Alias = "-"
		ctx.Imports.addImportPkg(pkg).Alias = "-"
	}
	ceousPkg := &myasthurts.Package{
		Name:       "ceous",
		ImportPath: "github.com/jamillosantos/go-ceous",
	}
	ctx.ModelsImports.addImportPkg(ceousPkg)
	ctx.Imports.addImportPkg(ceousPkg)
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
