// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: tpl/relation.gohtml

package tpl

import (
	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Relation generates tpl/relation.gohtml
func Relation(relation *models.ModelRelation) string {
	var _b strings.Builder
	RenderRelation(&_b, relation)
	return _b.String()
}

// RenderRelation render tpl/relation.gohtml
func RenderRelation(_buffer io.StringWriter, relation *models.ModelRelation) {
	_buffer.WriteString("\n\ntype ")
	_buffer.WriteString(gorazor.HTMLEscape(relation.RelationName()))
	_buffer.WriteString(" struct {\n\tkeys []interface{}\n\trecords map[")
	_buffer.WriteString(gorazor.HTMLEscape(relation.PkType()))
	_buffer.WriteString("][]")
	_buffer.WriteString(("*"))
	_buffer.WriteString(gorazor.HTMLEscape(relation.FromModel.Name))
	_buffer.WriteString("\n}\n\nfunc New")
	_buffer.WriteString(gorazor.HTMLEscape(relation.RelationName()))
	_buffer.WriteString("() ")
	_buffer.WriteString(("*"))
	_buffer.WriteString(gorazor.HTMLEscape(relation.RelationName()))
	_buffer.WriteString(" {\n\treturn ")
	_buffer.WriteString(("&"))
	_buffer.WriteString(gorazor.HTMLEscape(relation.RelationName()))
	_buffer.WriteString("{\n\t\tkeys:    make([]interface{}, 0),\n\t\trecords: make(map[")
	_buffer.WriteString(gorazor.HTMLEscape(relation.PkType()))
	_buffer.WriteString("][]")
	_buffer.WriteString(("*"))
	_buffer.WriteString(gorazor.HTMLEscape(relation.FromModel.Name))
	_buffer.WriteString("),\n\t}\n}\n\nfunc (relation ")
	_buffer.WriteString(("*"))
	_buffer.WriteString(gorazor.HTMLEscape(relation.RelationName()))
	_buffer.WriteString(") Aggregate(record ceous.Record) error {\n\tugRecord, ok := record.(")
	_buffer.WriteString(("*"))
	_buffer.WriteString(gorazor.HTMLEscape(relation.FromModel.Name))
	_buffer.WriteString(")\n\tif !ok {\n\t\treturn ceous.ErrInvalidRecordType\n\t}\n\tif rs, ok := relation.records[ugRecord.")
	_buffer.WriteString(gorazor.HTMLEscape(relation.Column().FullField))
	_buffer.WriteString("]; ok {\n\t\trelation.records[ugRecord.")
	_buffer.WriteString(gorazor.HTMLEscape(relation.Column().FullField))
	_buffer.WriteString("] = append(rs, ugRecord)\n\t\t// No need to add the key here, since its will be already in the `keys`.\n\t} else {\n\t\trelation.records[ugRecord.")
	_buffer.WriteString(gorazor.HTMLEscape(relation.Column().FullField))
	_buffer.WriteString("] = append(rs, ugRecord)\n\t\trelation.keys = append(relation.keys, ugRecord.")
	_buffer.WriteString(gorazor.HTMLEscape(relation.Column().FullField))
	_buffer.WriteString(")\n\t}\n\treturn nil\n}\n\nfunc (relation ")
	_buffer.WriteString(("*"))
	_buffer.WriteString(gorazor.HTMLEscape(relation.RelationName()))
	_buffer.WriteString(") Realize() error {\n\trecords, err := NewUserQuery(ceous.WithDB(DB)).Where(ceous.Eq(Schema.")
	_buffer.WriteString(gorazor.HTMLEscape(relation.ToModel.Name))
	_buffer.WriteString(".ID, relation.keys)).All()\n\tif err != nil {\n\t\treturn err // TODO(jota): Shall this be wrapped into a custom error?\n\t}\n\tfor _, record := range records {\n\t\tmasterRecords, ok := relation.records[record.ID]\n\t\tif !ok {\n\t\t\treturn ceous.ErrInconsistentRelationResult\n\t\t}\n\t\tfor _, r := range masterRecords {\n\t\t\tr.")
	_buffer.WriteString(gorazor.HTMLEscape(relation.FromField))
	_buffer.WriteString(" = record\n\t\t}\n\t}\n\treturn nil\n}\n\nfunc (q ")
	_buffer.WriteString(("*"))
	_buffer.WriteString(gorazor.HTMLEscape(relation.FromModel.QueryName()))
	_buffer.WriteString(") With")
	_buffer.WriteString(gorazor.HTMLEscape(relation.FromField))
	_buffer.WriteString("() ")
	_buffer.WriteString(("*"))
	_buffer.WriteString(gorazor.HTMLEscape(relation.FromModel.QueryName()))
	_buffer.WriteString(" {\n\tq.BaseQuery.Relations = append(q.BaseQuery.Relations, New")
	_buffer.WriteString(gorazor.HTMLEscape(relation.RelationName()))
	_buffer.WriteString("())\n\treturn q\n}")

}
