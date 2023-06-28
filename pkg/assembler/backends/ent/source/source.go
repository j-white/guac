// Code generated by ent, DO NOT EDIT.

package source

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the source type in the database.
	Label = "source"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeNamespaces holds the string denoting the namespaces edge name in mutations.
	EdgeNamespaces = "namespaces"
	// Table holds the table name of the source in the database.
	Table = "sources"
	// NamespacesTable is the table that holds the namespaces relation/edge.
	NamespacesTable = "source_namespaces"
	// NamespacesInverseTable is the table name for the SourceNamespace entity.
	// It exists in this package in order to avoid circular dependency with the "sourcenamespace" package.
	NamespacesInverseTable = "source_namespaces"
	// NamespacesColumn is the table column denoting the namespaces relation/edge.
	NamespacesColumn = "source_id"
)

// Columns holds all SQL columns for source fields.
var Columns = []string{
	FieldID,
	FieldType,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Source queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByNamespacesCount orders the results by namespaces count.
func ByNamespacesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNamespacesStep(), opts...)
	}
}

// ByNamespaces orders the results by namespaces terms.
func ByNamespaces(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNamespacesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newNamespacesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NamespacesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, NamespacesTable, NamespacesColumn),
	)
}
