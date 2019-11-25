// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: tpl/query.gohtml

package tpl

import (
	. "github.com/jamillosantos/go-ceous/generator/helpers"
	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Query generates tpl/query.gohtml
func Query(ctx *models.Ctx, model *models.Model) string {
	var _b strings.Builder
	RenderQuery(&_b, ctx, model)
	return _b.String()
}

// RenderQuery render tpl/query.gohtml
func RenderQuery(_buffer io.StringWriter, ctx *models.Ctx, model *models.Model) {
	_buffer.WriteString("\n\n// ")
	_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
	_buffer.WriteString(" is the query for the model `")
	_buffer.WriteString(gorazor.HTMLEscape(model.Name))
	_buffer.WriteString("`.\ntype ")
	_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
	_buffer.WriteString(" struct {\n\t*ceous.BaseQuery")
	for _, field := range model.Fields {
		_buffer.WriteString("\n\t")
		_buffer.WriteString(gorazor.HTMLEscape(field.Name))
		_buffer.WriteString("\tceous.SchemaField")
	}
	_buffer.WriteString("\n}\n\n// New")
	_buffer.WriteString(gorazor.HTMLEscape(model.Name))
	_buffer.WriteString("Query creates a new query for model `")
	_buffer.WriteString(gorazor.HTMLEscape(model.Name))
	_buffer.WriteString("`.\nfunc New")
	_buffer.WriteString(gorazor.HTMLEscape(model.Name))
	_buffer.WriteString("Query(options ...ceous.CeousOption) ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
	_buffer.WriteString(" {\n\tbq := ceous.NewBaseQuery(options...)\n\tif bq.Schema == nil {\n\t\tbq.Schema = ")
	_buffer.WriteString(gorazor.HTMLEscape(ctx.InputPkgCtx.Ref(ctx.OutputPkg, "Schema")))
	_buffer.WriteString(".")
	_buffer.WriteString(gorazor.HTMLEscape(model.Name))
	_buffer.WriteString(".BaseSchema\n\t}\n\treturn ")
	_buffer.WriteString(("&"))
	_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
	_buffer.WriteString("{\n\t\tBaseQuery: bq,\n\t}\n}")
	for _, field := range model.Conditions {
		_buffer.WriteString("\n// By")
		_buffer.WriteString(gorazor.HTMLEscape(field.NameForMethod))
		_buffer.WriteString(" add a filter by `")
		_buffer.WriteString(gorazor.HTMLEscape(field.Field))
		_buffer.WriteString("`.\nfunc (q ")
		_buffer.WriteString(gorazor.HTMLEscape(Pointer))
		_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
		_buffer.WriteString(") By")
		_buffer.WriteString(gorazor.HTMLEscape(field.NameForMethod))
		_buffer.WriteString("(value ")
		_buffer.WriteString(gorazor.HTMLEscape(ctx.Imports.Ref(field.Type.RefType)))
		_buffer.WriteString(") ")
		_buffer.WriteString(gorazor.HTMLEscape(Pointer))
		_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
		_buffer.WriteString(" {\n\t")
		for _, condition := range field.Conditions {
			_buffer.WriteString("\n\tq.BaseQuery.Where(ceous.Eq(")
			_buffer.WriteString(gorazor.HTMLEscape(ctx.InputPkgCtx.Ref(ctx.OutputPkg, "Schema")))
			_buffer.WriteString(".")
			_buffer.WriteString(gorazor.HTMLEscape(condition.SchemaField))
			_buffer.WriteString(", value")
			_buffer.WriteString(gorazor.HTMLEscape(condition.StringField()))
			_buffer.WriteString("))\n\t")
		}
		_buffer.WriteString("\n\treturn q\n}")
	}
	_buffer.WriteString("\n\n// Select defines what fields should be selected from the database.\nfunc (q ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
	_buffer.WriteString(") Select(fields ...ceous.SchemaField) ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
	_buffer.WriteString(" {\n\tq.BaseQuery.Select(fields...)\n\treturn q\n}\n\n// ExcludeFields defines what fields should not be selected from the database.\nfunc (q ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
	_buffer.WriteString(") ExcludeFields(fields ...ceous.SchemaField) ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
	_buffer.WriteString(" {\n\tq.BaseQuery.ExcludeFields(fields...)\n\treturn q\n}\n\n// Where defines the conditions for \nfunc (q ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
	_buffer.WriteString(") Where(pred interface{}, args ...interface{}) ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
	_buffer.WriteString(" {\n\tq.BaseQuery.Where(pred, args...)\n\treturn q\n}\n\nfunc (q ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
	_buffer.WriteString(") Limit(limit uint64) ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
	_buffer.WriteString(" {\n\tq.BaseQuery.Limit(limit)\n\treturn q\n}\n\nfunc (q ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
	_buffer.WriteString(") Offset(offset uint64) ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
	_buffer.WriteString(" {\n\tq.BaseQuery.Offset(offset)\n\treturn q\n}\n\n// One results only one record matching the query.\nfunc (q ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
	_buffer.WriteString(") One() (m ")
	_buffer.WriteString(gorazor.HTMLEscape(ctx.InputPkgCtx.Ref(ctx.OutputPkg, model.Name)))
	_buffer.WriteString(", err error) {\n\tq.Limit(1).Offset(0)\n\n\tquery, err := q.RawQuery()\n\tif err != nil {\n\t\treturn\n\t}\n\n\trs := New")
	_buffer.WriteString(gorazor.HTMLEscape(model.Name))
	_buffer.WriteString("ResultSet(query, nil)\n\tdefer rs.Close()\n\n\tif rs.Next() {\n\t\terr = rs.ToModel(&m)\n\t\tif err != nil {\n\t\t\treturn\n\t\t}\n\n\t\tfor _, rel := range q.BaseQuery.Relations {\n\t\t\terr = rel.Aggregate(&m)\n\t\t\tif err != nil {\n\t\t\t\treturn ")
	_buffer.WriteString(gorazor.HTMLEscape(ctx.InputPkgCtx.Ref(ctx.OutputPkg, model.Name)))
	_buffer.WriteString("{}, err // TODO(jota): Shall this error be wrapped? At first, yes.\n\t\t\t}\n\t\t}\n\t} else {\n\t\terr = ceous.ErrNotFound\n\t}\n\n\tif err == nil {\n\t\tfor _, rel := range q.BaseQuery.Relations {\n\t\t\terr = rel.Realize()\n\t\t\tif err != nil {\n\t\t\t\treturn ")
	_buffer.WriteString(gorazor.HTMLEscape(ctx.InputPkgCtx.Ref(ctx.OutputPkg, model.Name)))
	_buffer.WriteString("{}, err // TODO(jota): Shall this error be wrapped? At first, yes.\n\t\t\t}\n\t\t}\n\t}\n\n\treturn\n}\n\n// All return all records that match the query.\nfunc (q ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
	_buffer.WriteString(") All() ([]")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(ctx.InputPkgCtx.Ref(ctx.OutputPkg, model.Name)))
	_buffer.WriteString(", error) {\n\tquery, err := q.RawQuery()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\n\trs := New")
	_buffer.WriteString(gorazor.HTMLEscape(model.Name))
	_buffer.WriteString("ResultSet(query, nil)\n\tdefer rs.Close()\n\n\tms := make([]")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(ctx.InputPkgCtx.Ref(ctx.OutputPkg, model.Name)))
	_buffer.WriteString(", 0)\n\tfor rs.Next() {\n\t\tm := &")
	_buffer.WriteString(gorazor.HTMLEscape(ctx.InputPkgCtx.Ref(ctx.OutputPkg, model.Name)))
	_buffer.WriteString("{}\n\t\terr = rs.ToModel(m)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\n\t\tfor _, rel := range q.BaseQuery.Relations {\n\t\t\terr = rel.Aggregate(m)\n\t\t\tif err != nil {\n\t\t\t\treturn nil, err // TODO(jota): Shall this error be wrapped? At first, yes.\n\t\t\t}\n\t\t}\n\t\tms = append(ms, m)\n\t}\n\n\tfor _, rel := range q.BaseQuery.Relations {\n\t\terr = rel.Realize()\n\t\tif err != nil {\n\t\t\treturn nil, err // TODO(jota): Shall this error be wrapped? At first, yes.\n\t\t}\n\t}\n\n\treturn ms, nil\n}\n\nfunc (q ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
	_buffer.WriteString(") OrderBy(fields ...interface{}) ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
	_buffer.WriteString(" {\n\tq.BaseQuery.OrderBy(fields...)\n\treturn q\n}")
	for _, relation := range model.Relations {
		RenderRelation(_buffer, ctx, relation)
	}

}
