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

	GenContext struct {
		InputPkg    *myasthurts.Package
		InputPkgCtx *CtxPkg

		OutputPkg    *myasthurts.Package
		OutputPkgCtx *CtxPkg

		Reporter reporters.Reporter

		Models    []*Model
		ModelsMap map[string]*Model

		BaseSchemas []*BaseSchema
		Schemas     []*Schema

		Connections []*Connection

		Imports       *CtxImports
		ModelsImports *CtxImports
		count         int
	}

	ModelContext struct {
		Gen      *GenContext
		Reporter reporters.Reporter
		Model    *Model
	}

	SchemaContext struct {
		Gen            *GenContext
		Reporter       reporters.Reporter
		Model          *Model
		BaseSchema     *BaseSchema
		Schema         *Schema
		ColumnPrefix   []string
		FieldPath      []string
		FieldModifiers []ModelFieldModifier
	}

	SchemaFieldContext struct {
		Gen            *GenContext
		Model          *Model
		BaseSchema     *BaseSchema
		Schema         *Schema
		Reporter       reporters.Reporter
		FieldPath      []string
		ColumnPrefix   []string
		FieldModifiers []ModelFieldModifier
	}

	EmbeddedContext struct {
		Gen          *GenContext
		Model        *Model
		BaseSchema   *BaseSchema
		FieldPath    []string
		ColumnPrefix []string
		Modifiers    []ModelFieldModifier
	}
)

func NewGenContext(reporter reporters.Reporter, inputPackage, outputPackage *myasthurts.Package, pkgs ...*myasthurts.Package) *GenContext {
	ctx := &GenContext{
		Reporter:  reporter,
		InputPkg:  inputPackage,
		OutputPkg: outputPackage,

		Imports:       NewCtxImports(outputPackage),
		ModelsImports: NewCtxImports(inputPackage),

		BaseSchemas: make([]*BaseSchema, 0),

		Connections: make([]*Connection, 0),
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

func NewCtxImports(pkg *myasthurts.Package) *CtxImports {
	return &CtxImports{
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

// structKey returns a string that represents the struct in the context
// modelsMap.
func (ctx *GenContext) structKey(s *myasthurts.Struct) string {
	return s.Package().ImportPath + "." + s.Name()
}

// HasModel receives a `*myasthurts.Struct` and checks if it is defined in the
// models map.
func (ctx *GenContext) HasModel(s *myasthurts.Struct) (*Model, bool) {
	if ctx.Models == nil {
		return nil, false
	}
	model, ok := ctx.ModelsMap[ctx.structKey(s)]
	return model, ok
}

// AddModel adds a model to the Models list.
//
// If given `s` was already parsed, it returns the same instance and true.
// Otherwise, it parses the given `s` returns the new `*Model` and false.
func (ctx *GenContext) AddModel(s *myasthurts.Struct) (*Model, bool) {
	if model, ok := ctx.HasModel(s); ok {
		return model, true
	}
	if ctx.Models == nil {
		ctx.Models = make([]*Model, 0)
	}
	if ctx.ModelsMap == nil {
		ctx.ModelsMap = make(map[string]*Model, 0)
	}
	model := NewModel(s)
	ctx.Models = append(ctx.Models, model)
	ctx.ModelsMap[ctx.structKey(s)] = model
	return model, false
}

// AddBaseSchema creates a new schema, registers it in the context and returns the
// new instance.
func (ctx *GenContext) AddBaseSchema(schema *BaseSchema) *BaseSchema {
	if ctx.BaseSchemas == nil {
		ctx.BaseSchemas = make([]*BaseSchema, 0)
	}
	ctx.BaseSchemas = append(ctx.BaseSchemas, schema)
	return schema
}

// AddSchema creates a new schema, registers it in the context and returns the
// new instance.
func (ctx *GenContext) AddSchema(schema *Schema) *Schema {
	if ctx.Schemas == nil {
		ctx.Schemas = make([]*Schema, 0)
	}
	ctx.Schemas = append(ctx.Schemas, schema)
	return schema
}
