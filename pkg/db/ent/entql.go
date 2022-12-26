// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/basal-manager/pkg/db/ent/api"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 1)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   api.Table,
			Columns: api.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: api.FieldID,
			},
		},
		Type: "Api",
		Fields: map[string]*sqlgraph.FieldSpec{
			api.FieldCreatedAt:   {Type: field.TypeUint32, Column: api.FieldCreatedAt},
			api.FieldUpdatedAt:   {Type: field.TypeUint32, Column: api.FieldUpdatedAt},
			api.FieldDeletedAt:   {Type: field.TypeUint32, Column: api.FieldDeletedAt},
			api.FieldProtocol:    {Type: field.TypeString, Column: api.FieldProtocol},
			api.FieldServiceName: {Type: field.TypeString, Column: api.FieldServiceName},
			api.FieldMethod:      {Type: field.TypeString, Column: api.FieldMethod},
			api.FieldMethodName:  {Type: field.TypeString, Column: api.FieldMethodName},
			api.FieldPath:        {Type: field.TypeString, Column: api.FieldPath},
			api.FieldExported:    {Type: field.TypeBool, Column: api.FieldExported},
			api.FieldPathPrefix:  {Type: field.TypeString, Column: api.FieldPathPrefix},
			api.FieldDomains:     {Type: field.TypeJSON, Column: api.FieldDomains},
			api.FieldDepracated:  {Type: field.TypeBool, Column: api.FieldDepracated},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (aq *APIQuery) addPredicate(pred func(s *sql.Selector)) {
	aq.predicates = append(aq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the APIQuery builder.
func (aq *APIQuery) Filter() *APIFilter {
	return &APIFilter{config: aq.config, predicateAdder: aq}
}

// addPredicate implements the predicateAdder interface.
func (m *APIMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the APIMutation builder.
func (m *APIMutation) Filter() *APIFilter {
	return &APIFilter{config: m.config, predicateAdder: m}
}

// APIFilter provides a generic filtering capability at runtime for APIQuery.
type APIFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *APIFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *APIFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(api.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *APIFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(api.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *APIFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(api.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *APIFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(api.FieldDeletedAt))
}

// WhereProtocol applies the entql string predicate on the protocol field.
func (f *APIFilter) WhereProtocol(p entql.StringP) {
	f.Where(p.Field(api.FieldProtocol))
}

// WhereServiceName applies the entql string predicate on the service_name field.
func (f *APIFilter) WhereServiceName(p entql.StringP) {
	f.Where(p.Field(api.FieldServiceName))
}

// WhereMethod applies the entql string predicate on the method field.
func (f *APIFilter) WhereMethod(p entql.StringP) {
	f.Where(p.Field(api.FieldMethod))
}

// WhereMethodName applies the entql string predicate on the method_name field.
func (f *APIFilter) WhereMethodName(p entql.StringP) {
	f.Where(p.Field(api.FieldMethodName))
}

// WherePath applies the entql string predicate on the path field.
func (f *APIFilter) WherePath(p entql.StringP) {
	f.Where(p.Field(api.FieldPath))
}

// WhereExported applies the entql bool predicate on the exported field.
func (f *APIFilter) WhereExported(p entql.BoolP) {
	f.Where(p.Field(api.FieldExported))
}

// WherePathPrefix applies the entql string predicate on the path_prefix field.
func (f *APIFilter) WherePathPrefix(p entql.StringP) {
	f.Where(p.Field(api.FieldPathPrefix))
}

// WhereDomains applies the entql json.RawMessage predicate on the domains field.
func (f *APIFilter) WhereDomains(p entql.BytesP) {
	f.Where(p.Field(api.FieldDomains))
}

// WhereDepracated applies the entql bool predicate on the depracated field.
func (f *APIFilter) WhereDepracated(p entql.BoolP) {
	f.Where(p.Field(api.FieldDepracated))
}
