// Code generated by ent, DO NOT EDIT.

package certifyscorecard

import (
	"entgo.io/ent/dialect/gremlin/graph/dsl"
)

const (
	// Label holds the string label denoting the certifyscorecard type in the database.
	Label = "certify_scorecard"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSourceID holds the string denoting the source_id field in the database.
	FieldSourceID = "source_id"
	// FieldScorecardID holds the string denoting the scorecard_id field in the database.
	FieldScorecardID = "scorecard_id"
	// EdgeScorecard holds the string denoting the scorecard edge name in mutations.
	EdgeScorecard = "scorecard"
	// EdgeSource holds the string denoting the source edge name in mutations.
	EdgeSource = "source"
	// ScorecardInverseLabel holds the string label denoting the scorecard inverse edge type in the database.
	ScorecardInverseLabel = "scorecard_certifications"
	// SourceLabel holds the string label denoting the source edge type in the database.
	SourceLabel = "certify_scorecard_source"
)

// OrderOption defines the ordering options for the CertifyScorecard queries.
type OrderOption func(*dsl.Traversal)
