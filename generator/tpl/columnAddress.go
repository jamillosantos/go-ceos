// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: tpl/columnAddress.gohtml

package tpl

import (
	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// ColumnAddress generates tpl/columnAddress.gohtml
func ColumnAddress(baseSchema *models.BaseSchema) string {
	var _b strings.Builder
	RenderColumnAddress(&_b, baseSchema)
	return _b.String()
}

// RenderColumnAddress render tpl/columnAddress.gohtml
func RenderColumnAddress(_buffer io.StringWriter, baseSchema *models.BaseSchema) {
	_buffer.WriteString("\n\n// ColumnAddress returns the pointer address of a field given its column name.\nfunc (model ")
	_buffer.WriteString(("*"))
	_buffer.WriteString(gorazor.HTMLEscape(baseSchema.Name))
	_buffer.WriteString(") ColumnAddress(name string) (interface{}, error) {\n\tswitch name {")
	for _, field := range baseSchema.Fields {
		_buffer.WriteString("\n\tcase \"")
		_buffer.WriteString(gorazor.HTMLEscape(field.ColumnName))
		_buffer.WriteString("\":\n\t\treturn &model.")
		_buffer.WriteString(gorazor.HTMLEscape(field.Name))
		_buffer.WriteString(", nil")
	}
	_buffer.WriteString("\n\tdefault:\n\t\treturn nil, errors.Wrapf(ceous.ErrFieldNotFound, \"field %s not found\", name)\n\t}\n}")

}
