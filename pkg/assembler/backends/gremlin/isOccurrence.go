package gremlin

import (
	"context"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

const (
	IsOccurrence Label = "isOccurrence"
)

func createQueryToMatchArtifactInput[M any](artifact *model.ArtifactInputSpec) *gremlinQueryBuilder[M] {
	return createQueryForVertex[M](Artifact).
		withPropString(algorithm, &artifact.Algorithm).
		withPropString(digest, &artifact.Digest)
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
		withPropString(version, pkg.Version)
	// FIXME: Match on qualifiers too
}

func createQueryToMatchSourceInput[M any](src *model.SourceInputSpec) *gremlinQueryBuilder[M] {
	return createQueryForVertex[M](Package).
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

func getIsOccurrenceObjectFromEdge(result *gremlinQueryResult) *model.IsOccurrence {
	isOccurrence := &model.IsOccurrence{
		ID: result.id,
		// FIXME: Render these too
		Subject:       nil,
		Artifact:      nil,
		Justification: result.edge[origin].(string),
		Origin:        result.edge[origin].(string),
		Collector:     result.edge[collector].(string),
	}
	return isOccurrence
}

func createUpsertForIsOccurrence(subject *model.PackageOrSourceInput, artifact *model.ArtifactInputSpec, occurrence *model.IsOccurrenceInputSpec) *gremlinQueryBuilder[*model.IsOccurrence] {
	return createUpsertForEdge[*model.IsOccurrence](IsOccurrence).
		withPropString(justification, &occurrence.Justification).
		withPropString(origin, &occurrence.Origin).
		withPropString(collector, &occurrence.Collector).
		withOutVertex(createQueryToMatchPackageOrSource[*model.IsOccurrence](subject)).
		withInVertex(createQueryToMatchArtifactInput[*model.IsOccurrence](artifact)).
		withMapper(getIsOccurrenceObjectFromEdge)
}

func (c *gremlinClient) IngestOccurrence(ctx context.Context, subject model.PackageOrSourceInput, artifact model.ArtifactInputSpec, occurrence model.IsOccurrenceInputSpec) (*model.IsOccurrence, error) {
	return createUpsertForIsOccurrence(&subject, &artifact, &occurrence).upsert()
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
		upsertBulk()
}

func (c *gremlinClient) IsOccurrence(ctx context.Context, isOccurrenceSpec *model.IsOccurrenceSpec) ([]*model.IsOccurrence, error) {
	q := createQueryForEdge[*model.IsOccurrence](IsOccurrence).
		withId(isOccurrenceSpec.ID).
		withPropString(justification, isOccurrenceSpec.Justification).
		withPropString(origin, isOccurrenceSpec.Origin).
		withPropString(collector, isOccurrenceSpec.Collector).
		withMapper(getIsOccurrenceObjectFromEdge)
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
	return q.findAll()
}
