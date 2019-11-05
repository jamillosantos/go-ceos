// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: tpl/schema.gohtml

package tpl

import (
	generatorModels "github.com/jamillosantos/go-ceous/generator/models"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Schema generates tpl/schema.gohtml
func Schema(ctxPkg *generatorModels.Ctx, models []*generatorModels.Model, embeddeds []*generatorModels.Model) string {
	var _b strings.Builder
	RenderSchema(&_b, ctxPkg, models, embeddeds)
	return _b.String()
}

// RenderSchema render tpl/schema.gohtml
func RenderSchema(_buffer io.StringWriter, ctxPkg *generatorModels.Ctx, models []*generatorModels.Model, embeddeds []*generatorModels.Model) {
	_buffer.WriteString("package ")
	_buffer.WriteString(gorazor.HTMLEscape(ctxPkg.Pkg.Name))
	_buffer.WriteString("\n\nimport (\n\t\"github.com/jamillosantos/go-ceous\"\n\t\"github.com/pkg/errors\"")
	for _, pkg := range ctxPkg.Imports {
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
	for _, m := range models {

		_buffer.WriteString(("\n\n"))

		_buffer.WriteString("/** ")
		_buffer.WriteString(("\n"))

		_buffer.WriteString(" * Declare ")
		_buffer.WriteString(gorazor.HTMLEscape(m.Name))
		_buffer.WriteString(("\n"))

		_buffer.WriteString(" */")

		RenderModel(_buffer, ctxPkg.Pkg, m)
	}
	for _, m := range embeddeds {

		_buffer.WriteString(("\n\n"))

		RenderEmbedded(_buffer, m)
	}
	_buffer.WriteString("\n\ntype schema struct {")
	for _, m := range models {

		_buffer.WriteString(("\n"))

		_buffer.WriteString("\t")
		_buffer.WriteString(gorazor.HTMLEscape(m.Name))
		_buffer.WriteString(" ")
		_buffer.WriteString(("*"))
		_buffer.WriteString(gorazor.HTMLEscape(m.SchemaName()))
	}
	_buffer.WriteString("\n}\n\n// Schema represents the schema of the package \"")
	_buffer.WriteString(gorazor.HTMLEscape(ctxPkg.Pkg.Name))
	_buffer.WriteString("\".\nvar Schema = schema{")
	for _, m := range models {

		_buffer.WriteString(("\n"))

		_buffer.WriteString("\t")
		_buffer.WriteString(gorazor.HTMLEscape(m.Name))
		_buffer.WriteString(": ")
		_buffer.WriteString(("&"))
		_buffer.WriteString(gorazor.HTMLEscape(m.SchemaName()))
		_buffer.WriteString(" {\n\t\tBaseSchema: ")
		_buffer.WriteString(gorazor.HTMLEscape(m.BaseSchemaName()))
		_buffer.WriteString(",\n\t")
		for _, field := range m.Fields {

			_buffer.WriteString("\t")
			_buffer.WriteString(gorazor.HTMLEscape(field.Name))
			_buffer.WriteString(": ")
			if field.SchemaType == "" {
				i := m.ColumnsMap[field.FieldName]

				_buffer.WriteString(gorazor.HTMLEscape(m.BaseSchemaName()))

				_buffer.WriteString(".ColumnsArr[")
				_buffer.WriteString(gorazor.HTMLEscape(i))
				_buffer.WriteString("],")

			} else {

				_buffer.WriteString(gorazor.HTMLEscape(field.SchemaType))

				_buffer.WriteString("Schema,")

			}
			_buffer.WriteString("\n\t")
		}
		_buffer.WriteString("\n\t},")
	}
	_buffer.WriteString("\n}")
	for _, m := range models {

		_buffer.WriteString(("\n\n"))

		RenderResultset(_buffer, m)
	}
	for _, m := range models {

		_buffer.WriteString(("\n\n"))

		RenderQuery(_buffer, ctxPkg.Pkg, m)
	}
	for _, m := range models {

		_buffer.WriteString(("\n\n"))

		RenderStore(_buffer, m)
	}

}