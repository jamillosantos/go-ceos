// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: tpl/ceous.gohtml

package tpl

import (
	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Ceous generates tpl/ceous.gohtml
func Ceous(env *models.Environment) string {
	var _b strings.Builder
	RenderCeous(&_b, env)
	return _b.String()
}

// RenderCeous render tpl/ceous.gohtml
func RenderCeous(_buffer io.StringWriter, env *models.Environment) {
	_buffer.WriteString("// Code generated by https://github.com/jamillosantos/go-ceous DO NOT EDIT\n\npackage ")
	_buffer.WriteString(gorazor.HTMLEscape(env.OutputPkg.Name))
	_buffer.WriteString("\n\nimport (\n\t\"context\"\n\t\"database/sql\"\n\tceous \"github.com/jamillosantos/go-ceous\"")
	for _, pkg := range env.Imports.Imports {
		if pkg.Alias == "-" || pkg.Alias == "ceous" || pkg.Alias == "builtin" {
			continue
		}

		_buffer.WriteString(("\n"))

		_buffer.WriteString("\t")
		_buffer.WriteString(gorazor.HTMLEscape(pkg.Alias))
		_buffer.WriteString(" \"")
		_buffer.WriteString(gorazor.HTMLEscape(pkg.Pkg.ImportPath))
		_buffer.WriteString("\"")
	}
	_buffer.WriteString("\n)")
	RenderConnections(_buffer, env)
	RenderTransaction(_buffer, env)
	for _, query := range env.Queries {
		_buffer.WriteString("\n\t")
		RenderQuery(_buffer, env, query)
	}
	for _, store := range env.Stores {
		RenderStore(_buffer, env, store)
	}
	for _, schema := range env.Schemas {
		RenderResultset(_buffer, schema)
	}

}
