// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: tpl/model.gohtml

package tpl

import (
	. "github.com/jamillosantos/go-ceous/generator/helpers"
	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Model generates tpl/model.gohtml
func Model(env *models.Environment, model *models.Model) string {
	var _b strings.Builder
	RenderModel(&_b, env, model)
	return _b.String()
}

// RenderModel render tpl/model.gohtml
func RenderModel(_buffer io.StringWriter, env *models.Environment, model *models.Model) {
	if model.PK != nil {
		_buffer.WriteString("\n\n// GetID returns the primary key for model `")
		_buffer.WriteString(gorazor.HTMLEscape(model.Name))
		_buffer.WriteString("`.\nfunc (model ")
		_buffer.WriteString(gorazor.HTMLEscape(Pointer))
		_buffer.WriteString(gorazor.HTMLEscape(model.Name))
		_buffer.WriteString(") GetID() interface{} {\n\treturn model.")
		_buffer.WriteString(gorazor.HTMLEscape(model.PK.Name))
		_buffer.WriteString("\n}")
	}
	for _, relation := range model.Relations {
		_buffer.WriteString("\n// ")
		_buffer.WriteString(gorazor.HTMLEscape(PascalCase(relation.Name)))
		_buffer.WriteString(" returns the ")
		_buffer.WriteString(gorazor.HTMLEscape(relation.Name))
		_buffer.WriteString(" from ")
		_buffer.WriteString(gorazor.HTMLEscape(model.Name))
		_buffer.WriteString(".\nfunc (model ")
		_buffer.WriteString(gorazor.HTMLEscape(Pointer))
		_buffer.WriteString(gorazor.HTMLEscape(model.Name))
		_buffer.WriteString(") ")
		_buffer.WriteString(gorazor.HTMLEscape(PascalCase(relation.LocalField)))
		_buffer.WriteString("() ")
		_buffer.WriteString(gorazor.HTMLEscape(Pointer))
		_buffer.WriteString(gorazor.HTMLEscape(relation.ForeignModelType))
		_buffer.WriteString(" {\n\treturn model.")
		_buffer.WriteString(gorazor.HTMLEscape(relation.LocalField))
		_buffer.WriteString("\n}\n\n// Set")
		_buffer.WriteString(gorazor.HTMLEscape(PascalCase(relation.LocalField)))
		_buffer.WriteString(" updates the value for the ")
		_buffer.WriteString(gorazor.HTMLEscape(relation.LocalField))
		_buffer.WriteString(",\n// updating the ")
		_buffer.WriteString(gorazor.HTMLEscape(relation.LocalField))
		_buffer.WriteString(".\nfunc (model ")
		_buffer.WriteString(gorazor.HTMLEscape(Pointer))
		_buffer.WriteString(gorazor.HTMLEscape(model.Name))
		_buffer.WriteString(") Set")
		_buffer.WriteString(gorazor.HTMLEscape(PascalCase(relation.LocalField)))
		_buffer.WriteString("(value ")
		_buffer.WriteString(gorazor.HTMLEscape(Pointer))
		_buffer.WriteString(gorazor.HTMLEscape(relation.ForeignModelType))
		_buffer.WriteString(") error {\n\tlocalPkPtr, err := model.ColumnAddress(\"")
		_buffer.WriteString(gorazor.HTMLEscape(relation.ForeignColumn))
		_buffer.WriteString("\")\n\tif err != nil {\n\t\treturn err\n\t}\n\n\tlocalFKTypedPtr, ok := localPkPtr.(")
		_buffer.WriteString(gorazor.HTMLEscape(Pointer))
		_buffer.WriteString(gorazor.HTMLEscape(relation.ForeignFieldType))
		_buffer.WriteString(")\n\tif !ok {\n\t\treturn errors.New(\"invalid key type\") // TODO(jota): To formalize this error.\n\t}\n\t*localFKTypedPtr = value.")
		_buffer.WriteString(gorazor.HTMLEscape(relation.ForeignField))
		_buffer.WriteString("\n\tmodel.")
		_buffer.WriteString(gorazor.HTMLEscape(relation.LocalField))
		_buffer.WriteString(" = value\n\treturn nil\n}")
	}

}
