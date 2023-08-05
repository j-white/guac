package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// IngestScorecard is the resolver for the ingestScorecard field.
func (r *mutationResolver) IngestScorecard(ctx context.Context, source model.SourceInputSpec, scorecard model.ScorecardInputSpec) (*model.CertifyScorecard, error) {
	return r.Backend.IngestScorecard(ctx, source, scorecard)
}

// IngestScorecards is the resolver for the ingestScorecards field.
func (r *mutationResolver) IngestScorecards(ctx context.Context, sources []*model.SourceInputSpec, scorecards []*model.ScorecardInputSpec) ([]*model.CertifyScorecard, error) {
	return r.Backend.IngestScorecards(ctx, sources, scorecards)
}

// Scorecards is the resolver for the scorecards field.
func (r *queryResolver) Scorecards(ctx context.Context, scorecardSpec model.CertifyScorecardSpec) ([]*model.CertifyScorecard, error) {
	return r.Backend.Scorecards(ctx, &scorecardSpec)
}
