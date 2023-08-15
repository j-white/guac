package gremlin

import (
	"context"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"time"
)

const (
	HasSourceAt Label = "hasSourceAt"
)

func createUpsertForHasSourceAt(pkg *model.PkgInputSpec, pkgMatchType *model.MatchFlags, source *model.SourceInputSpec, hasSourceAt *model.HasSourceAtInputSpec) *gremlinQueryBuilder[*model.HasSourceAt] {
	return createUpsertForEdge[*model.HasSourceAt](HasSourceAt).
		withPropTime(knownSince, &hasSourceAt.KnownSince).
		withPropString(justification, &hasSourceAt.Justification).
		withPropString(origin, &hasSourceAt.Origin).
		withPropString(collector, &hasSourceAt.Collector).
		withOutVertex(createQueryToMatchPackageInputWithMatchType[*model.HasSourceAt](pkg, pkgMatchType)).
		withInVertex(createQueryToMatchSourceInput[*model.HasSourceAt](source)).
		withMapper(getHasSourceAtObjectFromEdge)
}

func getHasSourceAtObjectFromEdge(result *gremlinQueryResult) (*model.HasSourceAt, error) {
	hasSourceAt := &model.HasSourceAt{
		ID:            result.id,
		KnownSince:    result.edge[knownSince].(time.Time),
		Package:       getPackageObject(result.outId, result.out),
		Source:        getSourceObject(result.inId, result.in),
		Justification: result.edge[justification].(string),
		Origin:        result.edge[collector].(string),
		Collector:     result.edge[origin].(string),
	}
	return hasSourceAt, nil
}

// IngestHasSourceAt
//
//	pkg -> hasSourceAt -> src
func (c *gremlinClient) IngestHasSourceAt(ctx context.Context, pkg model.PkgInputSpec, pkgMatchType model.MatchFlags, source model.SourceInputSpec, hasSourceAt model.HasSourceAtInputSpec) (*model.HasSourceAt, error) {
	return createUpsertForHasSourceAt(&pkg, &pkgMatchType, &source, &hasSourceAt).upsert(c)
}

func (c *gremlinClient) HasSourceAt(ctx context.Context, hasSourceAtSpec *model.HasSourceAtSpec) ([]*model.HasSourceAt, error) {
	q := createQueryForEdge[*model.HasSourceAt](HasSourceAt).
		withId(hasSourceAtSpec.ID).
		withPropTime(knownSince, hasSourceAtSpec.KnownSince).
		withPropString(justification, hasSourceAtSpec.Justification).
		withPropString(origin, hasSourceAtSpec.Origin).
		withPropString(collector, hasSourceAtSpec.Collector).
		withMapper(getHasSourceAtObjectFromEdge)
	if hasSourceAtSpec.Package != nil {
		q = q.withOutVertex(createQueryToMatchPackage[*model.HasSourceAt](hasSourceAtSpec.Package))
	}
	if hasSourceAtSpec.Source != nil {
		q = q.withInVertex(createQueryToMatchSource[*model.HasSourceAt](hasSourceAtSpec.Source))
	}
	return q.findAll(c)
}
