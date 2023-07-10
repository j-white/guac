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
			//BuiltFrom:     builtFrom.,
			//BuiltBy:       c.convBuilder(bb),
			BuildType: slsa.BuildType,
			//SlsaPredicate: predicate,
			SlsaVersion: slsa.SlsaVersion,
			StartedOn:   slsa.StartedOn,
			FinishedOn:  slsa.FinishedOn,
			Origin:      slsa.Origin,
			Collector:   slsa.Collector,
		},
	}, nil
}
