package models

import myasthurts "github.com/lab259/go-my-ast-hurts"

// Environment represents all set of information that should be used to generate
// the models complement, queries, stores, schemas and connections.
type Environment struct {
	InputPkg       *myasthurts.Package
	InputPkgCtx    *CtxPkg
	OutputPkg      *myasthurts.Package
	OutputPkgCtx   *CtxPkg
	Queries        []*Query
	BaseSchemas    []*BaseSchema
	Schemas        []*Schema
	Stores         []*Store
	Connections    []*Connection
	connectionsMap map[string]*Connection
	Imports        *CtxImports
	ModelsImports  *CtxImports
}

// NewEnvironment returns a new instance of an `Environment`.
func NewEnvironment(inputPkg, outputPkg *myasthurts.Package, imports *CtxImports, modelsImports *CtxImports) *Environment {
	return &Environment{
		InputPkg: inputPkg,
		InputPkgCtx: &CtxPkg{
			Pkg:   inputPkg,
			Alias: inputPkg.Name,
		},
		OutputPkg: outputPkg,
		OutputPkgCtx: &CtxPkg{
			Pkg:   outputPkg,
			Alias: outputPkg.Name,
		},
		Queries:        make([]*Query, 0),
		BaseSchemas:    make([]*BaseSchema, 0),
		Schemas:        make([]*Schema, 0),
		Stores:         make([]*Store, 0),
		Connections:    make([]*Connection, 0),
		connectionsMap: make(map[string]*Connection, 0),
		Imports:        imports,
		ModelsImports:  modelsImports,
	}
}

// AddSchema adds a `Schema` to the environment schema list.
func (env *Environment) AddSchema(schema *Schema) *Schema {
	env.Schemas = append(env.Schemas, schema)
	return schema
}

// AddBaseSchema adds a `BaseSchema` to the environment base schema list.
func (env *Environment) AddBaseSchema(baseSchema *BaseSchema) *BaseSchema {
	env.BaseSchemas = append(env.BaseSchemas, baseSchema)
	return baseSchema
}

// AddQuery adds a `Query` to the environment query list.
func (env *Environment) AddQuery(query *Query) *Query {
	env.Queries = append(env.Queries, query)
	return query
}

// AddStore adds a `Store` to the environment store list.
func (env *Environment) AddStore(store *Store) *Store {
	env.Stores = append(env.Stores, store)
	return store
}

// EnsureConnection tries to find a Conection in the list, if it does not exists,
// the method will create one and return its reference.
func (env *Environment) EnsureConnection(name string) (*Connection, bool) {
	if conn, ok := env.connectionsMap[name]; ok {
		return conn, true
	}
	conn := NewConnection(name)
	env.Connections = append(env.Connections, conn)
	env.connectionsMap[name] = conn
	return conn, false
}
