// Code generated by ent, DO NOT EDIT.

package occurrence

import (
	"entgo.io/ent/dialect/gremlin/graph/dsl"
)

const (
	// Label holds the string label denoting the occurrence type in the database.
	Label = "occurrence"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldArtifactID holds the string denoting the artifact_id field in the database.
	FieldArtifactID = "artifact_id"
	// FieldJustification holds the string denoting the justification field in the database.
	FieldJustification = "justification"
	// FieldOrigin holds the string denoting the origin field in the database.
	FieldOrigin = "origin"
	// FieldCollector holds the string denoting the collector field in the database.
	FieldCollector = "collector"
	// FieldSourceID holds the string denoting the source_id field in the database.
	FieldSourceID = "source_id"
	// FieldPackageID holds the string denoting the package_id field in the database.
	FieldPackageID = "package_id"
	// EdgeArtifact holds the string denoting the artifact edge name in mutations.
	EdgeArtifact = "artifact"
	// EdgePackage holds the string denoting the package edge name in mutations.
	EdgePackage = "package"
	// EdgeSource holds the string denoting the source edge name in mutations.
	EdgeSource = "source"
	// ArtifactLabel holds the string label denoting the artifact edge type in the database.
	ArtifactLabel = "occurrence_artifact"
	// PackageLabel holds the string label denoting the package edge type in the database.
	PackageLabel = "occurrence_package"
	// SourceLabel holds the string label denoting the source edge type in the database.
	SourceLabel = "occurrence_source"
)

// OrderOption defines the ordering options for the Occurrence queries.
type OrderOption func(*dsl.Traversal)
