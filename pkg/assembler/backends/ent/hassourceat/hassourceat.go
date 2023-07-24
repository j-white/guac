// Code generated by ent, DO NOT EDIT.

package hassourceat

import (
	"entgo.io/ent/dialect/gremlin/graph/dsl"
)

const (
	// Label holds the string label denoting the hassourceat type in the database.
	Label = "has_source_at"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPackageVersionID holds the string denoting the package_version_id field in the database.
	FieldPackageVersionID = "package_version_id"
	// FieldPackageNameID holds the string denoting the package_name_id field in the database.
	FieldPackageNameID = "package_name_id"
	// FieldSourceID holds the string denoting the source_id field in the database.
	FieldSourceID = "source_id"
	// FieldKnownSince holds the string denoting the known_since field in the database.
	FieldKnownSince = "known_since"
	// FieldJustification holds the string denoting the justification field in the database.
	FieldJustification = "justification"
	// FieldOrigin holds the string denoting the origin field in the database.
	FieldOrigin = "origin"
	// FieldCollector holds the string denoting the collector field in the database.
	FieldCollector = "collector"
	// EdgePackageVersion holds the string denoting the package_version edge name in mutations.
	EdgePackageVersion = "package_version"
	// EdgeAllVersions holds the string denoting the all_versions edge name in mutations.
	EdgeAllVersions = "all_versions"
	// EdgeSource holds the string denoting the source edge name in mutations.
	EdgeSource = "source"
	// PackageVersionLabel holds the string label denoting the package_version edge type in the database.
	PackageVersionLabel = "has_source_at_package_version"
	// AllVersionsLabel holds the string label denoting the all_versions edge type in the database.
	AllVersionsLabel = "has_source_at_all_versions"
	// SourceLabel holds the string label denoting the source edge type in the database.
	SourceLabel = "has_source_at_source"
)

// OrderOption defines the ordering options for the HasSourceAt queries.
type OrderOption func(*dsl.Traversal)
