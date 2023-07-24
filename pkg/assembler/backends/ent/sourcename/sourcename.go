// Code generated by ent, DO NOT EDIT.

package sourcename

import (
	"entgo.io/ent/dialect/gremlin/graph/dsl"
)

const (
	// Label holds the string label denoting the sourcename type in the database.
	Label = "source_name"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCommit holds the string denoting the commit field in the database.
	FieldCommit = "commit"
	// FieldTag holds the string denoting the tag field in the database.
	FieldTag = "tag"
	// FieldNamespaceID holds the string denoting the namespace_id field in the database.
	FieldNamespaceID = "namespace_id"
	// EdgeNamespace holds the string denoting the namespace edge name in mutations.
	EdgeNamespace = "namespace"
	// EdgeOccurrences holds the string denoting the occurrences edge name in mutations.
	EdgeOccurrences = "occurrences"
	// NamespaceLabel holds the string label denoting the namespace edge type in the database.
	NamespaceLabel = "source_name_namespace"
	// OccurrencesInverseLabel holds the string label denoting the occurrences inverse edge type in the database.
	OccurrencesInverseLabel = "occurrence_source"
)

// OrderOption defines the ordering options for the SourceName queries.
type OrderOption func(*dsl.Traversal)
