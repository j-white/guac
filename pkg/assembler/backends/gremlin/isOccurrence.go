package gremlin

import (
	"context"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

const (
	IsOccurrence Label = "isOccurrence"
)

func createQueryToMatchArtifactInput(artifact *model.ArtifactInputSpec) *gremlinQueryBuilder {
	q := createQueryForVertex(Artifact)
	q.withPropString(algorithm, &artifact.Algorithm)
	q.withPropString(digest, &artifact.Digest)
	return q
}

func createQueryToMatchArtifact(artifact *model.ArtifactSpec) *gremlinQueryBuilder {
	q := createQueryForVertex(Artifact)
	q.withId(artifact.ID)
	q.withPropString(algorithm, artifact.Algorithm)
	q.withPropString(digest, artifact.Digest)
	return q
}

func createQueryToMatchPackageInput(pkg *model.PkgInputSpec) *gremlinQueryBuilder {
	q := createQueryForVertex(Package)
	q.withPropString(typeStr, &pkg.Type)
	q.withPropString(name, &pkg.Name)
	q.withPropString(namespace, pkg.Namespace)
	q.withPropString(subpath, pkg.Subpath)
	// FIXMEs
	if pkg.Version != nil {
		// *filter.Version != v.version ||
		//	noMatchInput(filter.Subpath, v.subpath) ||
		//	noMatchQualifiers(filter, v.qualifiers) {
	}
	return q
}

func createQueryToMatchSourceInput(src *model.SourceInputSpec) *gremlinQueryBuilder {
	q := createQueryForVertex(Package)
	q.withPropString(typeStr, &src.Type)
	q.withPropString(name, &src.Name)
	q.withPropString(namespace, &src.Namespace)
	q.withPropString(commit, src.Commit)
	q.withPropString(tag, src.Tag)
	return q
}

func createQueryToMatchPackageOrSource(subject *model.PackageOrSourceInput) *gremlinQueryBuilder {
	var q *gremlinQueryBuilder
	if subject.Package != nil {
		q = createQueryToMatchPackageInput(subject.Package)
	} else if subject.Source != nil {
		q = createQueryToMatchSourceInput(subject.Source)
	} else {
		q = nil
	}
	return q
}

func (q *gremlinQueryBuilder) withOccurrenceQueryValues(occurrence *model.IsOccurrenceInputSpec) *gremlinQueryBuilder {
	return q.withPropString(justification, &occurrence.Justification).
		withPropString(origin, &occurrence.Origin).
		withPropString(collector, &occurrence.Collector)
}

func getIsOccurrenceObjectFromEdge(result *gremlinQueryResult) *model.IsOccurrence {
	isOccurrence := &model.IsOccurrence{
		ID:            result.id,
		Subject:       nil,
		Artifact:      nil,
		Justification: result.edge[origin].(string),
		Origin:        result.edge[origin].(string),
		Collector:     result.edge[collector].(string),
	}
	return isOccurrence
}

func createUpsertForIsDependency(subject *model.PackageOrSourceInput, artifact *model.ArtifactInputSpec, occurrence *model.IsOccurrenceInputSpec) *gremlinQueryBuilder {
	return createUpsertForEdge[*model.IsOccurrence](IsOccurrence).
		withOccurrenceQueryValues(occurrence).
		withOutVertex(createQueryToMatchPackageOrSource(subject)).
		withInVertex(createQueryToMatchArtifactInput(artifact)).
		withMapper(getIsOccurrenceObjectFromEdge)
}

func (c *gremlinClient) IngestOccurrence(ctx context.Context, subject model.PackageOrSourceInput, artifact model.ArtifactInputSpec, occurrence model.IsOccurrenceInputSpec) (*model.IsOccurrence, error) {
	return createUpsertForIsDependency(&subject, &artifact, &occurrence).upsert()
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
	var queries []*gremlinQueryBuilder
	for k := range subjects {
		queries = append(queries, createUpsertForIsDependency(subjects[k], artifacts[k], occurrences[k]))
	}

	return createBulkUpsertForEdge[*model.IsOccurrence](IsDependency).
		withQueries(queries).
		upsert()
}

func (c *gremlinClient) IsOccurrence(ctx context.Context, isOccurrenceSpec *model.IsOccurrenceSpec) ([]*model.IsOccurrence, error) {
	q := createQueryForEdge(IsOccurrence).
		withId(isOccurrenceSpec.ID).
		withPropString(justification, isOccurrenceSpec.Justification).
		withPropString(origin, isOccurrenceSpec.Origin).
		withPropString(collector, isOccurrenceSpec.Collector).
		withMapper(getIsOccurrenceObjectFromEdge)
	if isOccurrenceSpec.Artifact != nil {
		q = q.withOutVertex(createQueryToMatchArtifact(isOccurrenceSpec.Artifact))
	}
	if isOccurrenceSpec.Subject != nil {
		if isOccurrenceSpec.Subject.Package != nil {
			q = q.withInVertex(createQueryToMatchPackage(isOccurrenceSpec.Subject.Package))
		} else if isOccurrenceSpec.Subject.Source != nil {
			q = q.withInVertex(createQueryToMatchSource(isOccurrenceSpec.Subject.Source))
		}
	}
	return q.find()
}
