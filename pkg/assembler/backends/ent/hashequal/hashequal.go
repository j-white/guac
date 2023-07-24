// Code generated by ent, DO NOT EDIT.

package hashequal

import (
	"entgo.io/ent/dialect/gremlin/graph/dsl"
)

const (
	// Label holds the string label denoting the hashequal type in the database.
	Label = "hash_equal"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOrigin holds the string denoting the origin field in the database.
	FieldOrigin = "origin"
	// FieldCollector holds the string denoting the collector field in the database.
	FieldCollector = "collector"
	// FieldJustification holds the string denoting the justification field in the database.
	FieldJustification = "justification"
	// EdgeArtifacts holds the string denoting the artifacts edge name in mutations.
	EdgeArtifacts = "artifacts"
	// ArtifactsLabel holds the string label denoting the artifacts edge type in the database.
	ArtifactsLabel = "hash_equal_artifacts"
)

// OrderOption defines the ordering options for the HashEqual queries.
type OrderOption func(*dsl.Traversal)
