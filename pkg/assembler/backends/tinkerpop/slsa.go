package tinkerpop

import (
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

import (
	"context"
)

func (c *tinkerpopClient) IngestSLSA(ctx context.Context, subject model.ArtifactInputSpec,
	builtFrom []*model.ArtifactInputSpec, builtBy model.BuilderInputSpec,
	slsa model.SLSAInputSpec) (*model.HasSlsa, error) {
	if len(builtFrom) < 1 {
		return nil, gqlerror.Errorf("IngestSLSA :: Must have at least 1 builtFrom")
	}

	artifact := &model.Artifact{
		ID:        "1",
		Algorithm: subject.Algorithm,
		Digest:    subject.Digest,
	}
	//
	//predicate := &model.SLSAPredicate{
	//
	//}
	return &model.HasSlsa{
		ID:      "1",
		Subject: artifact,
		Slsa: &model.Slsa{
			BuiltFrom:     nil,
			BuiltBy:       &model.Builder{},
			BuildType:     slsa.BuildType,
			SlsaPredicate: nil,
			SlsaVersion:   slsa.SlsaVersion,
			StartedOn:     slsa.StartedOn,
			FinishedOn:    slsa.FinishedOn,
			Origin:        slsa.Origin,
			Collector:     slsa.Collector,
		},
	}, nil
}

func (c *tinkerpopClient) IngestSLSAs(ctx context.Context, subjects []*model.ArtifactInputSpec, builtFromList [][]*model.ArtifactInputSpec, builtByList []*model.BuilderInputSpec, slsaList []*model.SLSAInputSpec) ([]*model.HasSlsa, error) {
	var hasSlsaList []*model.HasSlsa
	for k := range subjects {
		hasSlsa, _ := c.IngestSLSA(ctx, *subjects[k], builtFromList[k], *builtByList[k], *slsaList[k])
		hasSlsaList = append(hasSlsaList, hasSlsa)
	}
	return hasSlsaList, nil
}
