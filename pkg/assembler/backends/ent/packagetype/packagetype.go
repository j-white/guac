// Code generated by ent, DO NOT EDIT.

package packagetype

import (
	"entgo.io/ent/dialect/gremlin/graph/dsl"
)

const (
	// Label holds the string label denoting the packagetype type in the database.
	Label = "package_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeNamespaces holds the string denoting the namespaces edge name in mutations.
	EdgeNamespaces = "namespaces"
	// NamespacesLabel holds the string label denoting the namespaces edge type in the database.
	NamespacesLabel = "package_type_namespaces"
)

var (
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
)

// OrderOption defines the ordering options for the PackageType queries.
type OrderOption func(*dsl.Traversal)
