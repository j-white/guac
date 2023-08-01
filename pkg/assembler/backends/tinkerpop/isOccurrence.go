package tinkerpop

import (
	"context"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

func (c *tinkerpopClient) IsOccurrence(ctx context.Context, isOccurrenceSpec *model.IsOccurrenceSpec) ([]*model.IsOccurrence, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestOccurrence(ctx context.Context, subject model.PackageOrSourceInput, artifact model.ArtifactInputSpec, occurrence model.IsOccurrenceInputSpec) (*model.IsOccurrence, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestOccurrences(ctx context.Context, subjects model.PackageOrSourceInputs, artifacts []*model.ArtifactInputSpec, occurrences []*model.IsOccurrenceInputSpec) ([]*model.IsOccurrence, error) {
	var isOccurrenceList []*model.IsOccurrence
	isOccurrence := &model.IsOccurrence{
		Subject:  model.Source{},
		Artifact: &model.Artifact{},
	}
	isOccurrenceList = append(isOccurrenceList, isOccurrence)
	return isOccurrenceList, nil
}
