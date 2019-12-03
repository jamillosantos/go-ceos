// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: tpl/store.gohtml

package tpl

import (
	. "github.com/jamillosantos/go-ceous/generator/helpers"
	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Store generates tpl/store.gohtml
func Store(env *models.Environment, store *models.Store) string {
	var _b strings.Builder
	RenderStore(&_b, env, store)
	return _b.String()
}

// RenderStore render tpl/store.gohtml
func RenderStore(_buffer io.StringWriter, env *models.Environment, store *models.Store) {
	_buffer.WriteString("\n\n// ")
	_buffer.WriteString(gorazor.HTMLEscape(store.FullName))
	_buffer.WriteString(" is the query for the store `")
	_buffer.WriteString(gorazor.HTMLEscape(store.Name))
	_buffer.WriteString("`.\ntype ")
	_buffer.WriteString(gorazor.HTMLEscape(store.FullName))
	_buffer.WriteString(" struct {\n\t*ceous.BaseStore\n}\n\n// New")
	_buffer.WriteString(gorazor.HTMLEscape(store.FullName))
	_buffer.WriteString(" creates a new query for model `")
	_buffer.WriteString(gorazor.HTMLEscape(store.Name))
	_buffer.WriteString("`.\nfunc New")
	_buffer.WriteString(gorazor.HTMLEscape(store.FullName))
	_buffer.WriteString("(options ...ceous.CeousOption) ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(store.FullName))
	_buffer.WriteString(" {\n\treturn ")
	_buffer.WriteString(("&"))
	_buffer.WriteString(gorazor.HTMLEscape(store.FullName))
	_buffer.WriteString("{\n\t\tBaseStore: ceous.NewStore(")
	_buffer.WriteString(gorazor.HTMLEscape(env.InputPkgCtx.Ref(env.OutputPkg, "Schema."+store.Name)))
	_buffer.WriteString(", options...),\n\t}\n}\n\nfunc (store ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(store.FullName))
	_buffer.WriteString(") Insert(record ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(env.InputPkgCtx.Ref(env.OutputPkg, store.Name)))
	_buffer.WriteString(", fields ...ceous.SchemaField) error {\n\treturn store.BaseStore.Insert(record, fields...)\n}\n\nfunc (store ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(store.FullName))
	_buffer.WriteString(") Update(record ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(env.InputPkgCtx.Ref(env.OutputPkg, store.Name)))
	_buffer.WriteString(", fields ...ceous.SchemaField) (int64, error) {\n\treturn store.BaseStore.Update(record, fields...)\n}")

}
