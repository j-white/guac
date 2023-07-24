// Code generated by ent, DO NOT EDIT.

package sourcetype

import (
	"entgo.io/ent/dialect/gremlin/graph/dsl"
)

const (
	// Label holds the string label denoting the sourcetype type in the database.
	Label = "source_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeNamespaces holds the string denoting the namespaces edge name in mutations.
	EdgeNamespaces = "namespaces"
	// NamespacesInverseLabel holds the string label denoting the namespaces inverse edge type in the database.
	NamespacesInverseLabel = "source_namespace_source_type"
)

// OrderOption defines the ordering options for the SourceType queries.
type OrderOption func(*dsl.Traversal)
