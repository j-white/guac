package gremlin

import (
	"context"
	"fmt"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

const (
	IsOccurrence Label = "isOccurrence"
)

func createQueryToMatchArtifactInput[M any](artifact *model.ArtifactInputSpec) *gremlinQueryBuilder[M] {
	return createQueryForVertex[M](Artifact).
		withPropStringToLower(algorithm, &artifact.Algorithm).
		withPropStringToLower(digest, &artifact.Digest)
}

func createQueryToMatchArtifact[M any](artifact *model.ArtifactSpec) *gremlinQueryBuilder[M] {
	return createQueryForVertex[M](Artifact).
		withId(artifact.ID).
		withPropString(algorithm, artifact.Algorithm).
		withPropString(digest, artifact.Digest)
}

func createQueryToMatchPackageInput[M any](pkg *model.PkgInputSpec) *gremlinQueryBuilder[M] {
	return createQueryForVertex[M](Package).
		withPropString(typeStr, &pkg.Type).
		withPropString(name, &pkg.Name).
		withPropString(namespace, pkg.Namespace).
		withPropString(subpath, pkg.Subpath).
		withPropStringOrEmpty(version, pkg.Version)
	// FIXME: Match on qualifiers too
}

func createQueryToMatchSourceInput[M any](src *model.SourceInputSpec) *gremlinQueryBuilder[M] {
	return createQueryForVertex[M](Source).
		withPropString(typeStr, &src.Type).
		withPropString(name, &src.Name).
		withPropString(namespace, &src.Namespace).
		withPropString(commit, src.Commit).
		withPropString(tag, src.Tag)
}

func createQueryToMatchPackageOrSource[M any](subject *model.PackageOrSourceInput) *gremlinQueryBuilder[M] {
	var q *gremlinQueryBuilder[M]
	if subject.Package != nil {
		q = createQueryToMatchPackageInput[M](subject.Package)
	} else if subject.Source != nil {
		q = createQueryToMatchSourceInput[M](subject.Source)
	} else {
		q = nil
	}
	return q
}

func createQueryToMatchSource[M any](src *model.SourceSpec) *gremlinQueryBuilder[M] {
	query := createGraphQuery(Source)
	if src.ID != nil {
		query.id = *src.ID
	}
	if src.Name != nil {
		query.has[name] = *src.Name
	}
	if src.Type != nil {
		query.has[typeStr] = *src.Type
	}
	if src.Namespace != nil {
		query.has[namespace] = *src.Namespace
	}
	if src.Commit != nil {
		query.has[commit] = *src.Commit
	}
	if src.Tag != nil {
		query.has[tag] = *src.Tag
	}
	return &gremlinQueryBuilder[M]{query: query}
}

func getIsOccurrenceObjectFromEdge(result *gremlinQueryResult) (*model.IsOccurrence, error) {
	isOccurrence := &model.IsOccurrence{
		ID:            result.id,
		Artifact:      getArtifactObject(result.outId, result.out),
		Justification: result.edge[justification].(string),
		Origin:        result.edge[origin].(string),
		Collector:     result.edge[collector].(string),
	}

	if result.inLabel == Source {
		isOccurrence.Subject = getSourceObject(result.inId, result.in)
	} else if result.inLabel == Package {
		isOccurrence.Subject = getPackageObject(result.inId, result.in)
	} else {
		return nil, fmt.Errorf("unsupported result type: %v", result.inLabel)
	}

	return isOccurrence, nil
}

func createUpsertForIsOccurrence(subject *model.PackageOrSourceInput, artifact *model.ArtifactInputSpec, occurrence *model.IsOccurrenceInputSpec) *gremlinQueryBuilder[*model.IsOccurrence] {
	return createUpsertForEdge[*model.IsOccurrence](IsOccurrence).
		withPropString(justification, &occurrence.Justification).
		withPropString(origin, &occurrence.Origin).
		withPropString(collector, &occurrence.Collector).
		withOutVertex(createQueryToMatchArtifactInput[*model.IsOccurrence](artifact)).
		withInVertex(createQueryToMatchPackageOrSource[*model.IsOccurrence](subject)).
		withMapper(getIsOccurrenceObjectFromEdge)
}

func (c *gremlinClient) IngestOccurrence(ctx context.Context, subject model.PackageOrSourceInput, artifact model.ArtifactInputSpec, occurrence model.IsOccurrenceInputSpec) (*model.IsOccurrence, error) {
	return createUpsertForIsOccurrence(&subject, &artifact, &occurrence).upsert(c)
}

func (c *gremlinClient) IngestOccurrences(ctx context.Context, inputs model.PackageOrSourceInputs, artifacts []*model.ArtifactInputSpec, occurrences []*model.IsOccurrenceInputSpec) ([]*model.IsOccurrence, error) {
	// subjects relate to either packages, or sources, not a mix
	var subjects []*model.PackageOrSourceInput
	if len(inputs.Packages) > 0 {
		for _, pkg := range inputs.Packages {
			subjects = append(subjects, &model.PackageOrSourceInput{Package: pkg})
		}
	} else if len(inputs.Sources) > 0 {
		for _, src := range inputs.Sources {
			subjects = append(subjects, &model.PackageOrSourceInput{Source: src})
		}
	}

	// build the queries
	var queries []*gremlinQueryBuilder[*model.IsOccurrence]
	for k := range subjects {
		queries = append(queries, createUpsertForIsOccurrence(subjects[k], artifacts[k], occurrences[k]))
	}

	return createBulkUpsertForEdge[*model.IsOccurrence](IsDependency).
		withQueries(queries).
		upsertBulk(c)
}

func (c *gremlinClient) IsOccurrence(ctx context.Context, isOccurrenceSpec *model.IsOccurrenceSpec) ([]*model.IsOccurrence, error) {
	q := createQueryForEdge[*model.IsOccurrence](IsOccurrence).
		withId(isOccurrenceSpec.ID).
		withPropString(justification, isOccurrenceSpec.Justification).
		withPropString(origin, isOccurrenceSpec.Origin).
		withPropString(collector, isOccurrenceSpec.Collector).
		withMapper(getIsOccurrenceObjectFromEdge).
		withOrderByKey(justification)
	if isOccurrenceSpec.Artifact != nil {
		q = q.withOutVertex(createQueryToMatchArtifact[*model.IsOccurrence](isOccurrenceSpec.Artifact))
	}
	if isOccurrenceSpec.Subject != nil {
		if isOccurrenceSpec.Subject.Package != nil {
			q = q.withInVertex(createQueryToMatchPackage[*model.IsOccurrence](isOccurrenceSpec.Subject.Package))
		} else if isOccurrenceSpec.Subject.Source != nil {
			q = q.withInVertex(createQueryToMatchSource[*model.IsOccurrence](isOccurrenceSpec.Subject.Source))
		}
	}
	return q.findAll(c)
}
