package gremlin

import (
	"context"
	"fmt"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"time"
)

const (
	HasMetadata Label = "hasMetadata"
	Metadata    Label = "metadata"
)

func getMetadataFromEdge(result *gremlinQueryResult) (*model.HasMetadata, error) {
	hasMetadata := &model.HasMetadata{
		ID:            result.inId,
		Key:           result.in[key].(string),
		Value:         result.in[value].(string),
		Timestamp:     result.in[timestamp].(time.Time),
		Justification: result.in[justification].(string),
		Origin:        result.in[collector].(string),
		Collector:     result.in[origin].(string),
	}
	if result.outLabel == Package {
		hasMetadata.Subject = getPackageObject(result.outId, result.out)
	} else if result.outLabel == Source {
		hasMetadata.Subject = getSourceObject(result.outId, result.out)
	} else if result.outLabel == Artifact {
		hasMetadata.Subject = getArtifactObject(result.outId, result.out)
	} else {
		return nil, fmt.Errorf("unsupported label: %v", result.outLabel)
	}
	return hasMetadata, nil
}

func createUpsertForMetadataVertex(hasMetadata *model.HasMetadataInputSpec) *gremlinQueryBuilder[*model.HasMetadata] {
	return createUpsertForVertex[*model.HasMetadata](Metadata).
		withPropString(key, &hasMetadata.Key).
		withPropString(value, &hasMetadata.Value).
		withPropTime(timestamp, &hasMetadata.Timestamp).
		withPropString(justification, &hasMetadata.Justification).
		withPropString(origin, &hasMetadata.Origin).
		withPropString(collector, &hasMetadata.Collector)
}

func createUpsertForMetadata(subject *model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, hasMetadata *model.HasMetadataInputSpec) *gremlinQueryBuilder[*model.HasMetadata] {
	return createUpsertForEdge[*model.HasMetadata](HasMetadata).
		withInVertex(createUpsertForMetadataVertex(hasMetadata)).
		withOutVertex(createQueryToMatchPackageSourceOrArtifactInput[*model.HasMetadata](subject, pkgMatchType)).
		withMapper(getMetadataFromEdge)
}

func (c *gremlinClient) IngestHasMetadata(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, hasMetadata model.HasMetadataInputSpec) (*model.HasMetadata, error) {
	return createUpsertForMetadata(&subject, pkgMatchType, &hasMetadata).upsert(c)
}

func (c *gremlinClient) HasMetadata(ctx context.Context, hasMetadataSpec *model.HasMetadataSpec) ([]*model.HasMetadata, error) {
	q := createQueryForEdge[*model.HasMetadata](HasMetadata).
		withInVertex(createQueryForVertex[*model.HasMetadata](Metadata).
			withId(hasMetadataSpec.ID).
			withPropString(key, hasMetadataSpec.Key).
			withPropString(value, hasMetadataSpec.Value).
			withPropTimeGreaterOrEqual(timestamp, hasMetadataSpec.Since).
			withPropString(justification, hasMetadataSpec.Justification).
			withPropString(origin, hasMetadataSpec.Origin).
			withPropString(collector, hasMetadataSpec.Collector)).
		withMapper(getMetadataFromEdge)
	if hasMetadataSpec.Subject != nil {
		q = q.withOutVertex(createQueryToMatchPackageSourceOrArtifactSpec[*model.HasMetadata](hasMetadataSpec.Subject))
	}
	return q.findAllEdges(c)
}
