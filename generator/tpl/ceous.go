// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: tpl/ceous.gohtml

package tpl

import (
	generatorModels "github.com/jamillosantos/go-ceous/generator/models"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Ceous generates tpl/ceous.gohtml
func Ceous(ctxPkg *generatorModels.Ctx, models []*generatorModels.Model, embeddeds []*generatorModels.Model, connections []*generatorModels.Connection) string {
	var _b strings.Builder
	RenderCeous(&_b, ctxPkg, models, embeddeds, connections)
	return _b.String()
}

// RenderCeous render tpl/ceous.gohtml
func RenderCeous(_buffer io.StringWriter, ctxPkg *generatorModels.Ctx, models []*generatorModels.Model, embeddeds []*generatorModels.Model, connections []*generatorModels.Connection) {
	_buffer.WriteString("package ")
	_buffer.WriteString(gorazor.HTMLEscape(ctxPkg.Pkg.Name))
	_buffer.WriteString("\n\nimport (\n\t\"github.com/jamillosantos/go-ceous\"\n\t\"github.com/pkg/errors\"\n\t\"context\"\n\t\"database/sql\"")
	for _, pkg := range ctxPkg.Imports.Imports {
		if pkg.Alias == "-" {
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
	for _, m := range embeddeds {

		_buffer.WriteString(("\n\n"))

		RenderEmbedded(_buffer, m)
	}
	if len(connections) > 0 {
		RenderConnections(_buffer, connections, models)
	}
	RenderTransaction(_buffer, models)
	RenderSchema(_buffer, ctxPkg, models)
	for _, m := range models {
		_buffer.WriteString("\n\t")
		RenderResultset(_buffer, m)
		_buffer.WriteString("\n\t")
		RenderQuery(_buffer, ctxPkg.Pkg, m)
		_buffer.WriteString("\n\t")
		RenderStore(_buffer, m)
	}

}