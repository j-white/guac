package gremlin

import (
	"context"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

func (c *gremlinClient) HasSBOM(ctx context.Context, hasSBOMSpec *model.HasSBOMSpec) ([]*model.HasSbom, error) {
	return nil, nil
}

func (c *gremlinClient) IngestHasSbom(ctx context.Context, subject model.PackageOrArtifactInput, hasSbom model.HasSBOMInputSpec) (*model.HasSbom, error) {
	hasSBOM := &model.HasSbom{
		ID:               "",
		Subject:          &model.Package{},
		URI:              "",
		Algorithm:        "",
		Digest:           "",
		DownloadLocation: "",
		Origin:           "",
		Collector:        "",
	}
	return hasSBOM, nil
}

func (c *gremlinClient) IngestHasSBOMs(ctx context.Context, subjects model.PackageOrArtifactInputs, hasSBOMs []*model.HasSBOMInputSpec) ([]*model.HasSbom, error) {
	return nil, nil
}
