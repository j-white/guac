package gremlin

import (
	"context"
	"fmt"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

const (
	CertifyBad          Label = "certifyBad"
	SubjectToCertifyBad Label = "subject-to-certifyBad"
)

func createUpsertForCertifyBadVertex(certifyBad *model.CertifyBadInputSpec) *gremlinQueryBuilder[*model.CertifyBad] {
	return createUpsertForVertex[*model.CertifyBad](CertifyBad).
		withPropString(justification, &certifyBad.Justification).
		withPropString(origin, &certifyBad.Origin).
		withPropString(collector, &certifyBad.Collector)
}

func getCertifyBadFromEdge(result *gremlinQueryResult) (*model.CertifyBad, error) {
	certifyBad := &model.CertifyBad{
		ID:            result.inId,
		Justification: result.in[justification].(string),
		Origin:        result.in[collector].(string),
		Collector:     result.in[origin].(string),
	}
	if result.outLabel == Package {
		certifyBad.Subject = getPackageObject(result.outId, result.out)
	} else if result.outLabel == Source {
		certifyBad.Subject = getSourceObject(result.outId, result.out)
	} else if result.outLabel == Artifact {
		certifyBad.Subject = getArtifactObject(result.outId, result.out)
	} else {
		return nil, fmt.Errorf("unsupported label: %v", result.outLabel)
	}
	return certifyBad, nil
}

func createUpsertForCertifyBad(subject *model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, certifyBad *model.CertifyBadInputSpec) *gremlinQueryBuilder[*model.CertifyBad] {
	return createUpsertForEdge[*model.CertifyBad](SubjectToCertifyBad).
		withInVertex(createUpsertForCertifyBadVertex(certifyBad)).
		withOutVertex(createQueryToMatchPackageSourceOrArtifactInput[*model.CertifyBad](subject, pkgMatchType)).
		withMapper(getCertifyBadFromEdge)
}

func (c *gremlinClient) IngestCertifyBad(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, certifyBad model.CertifyBadInputSpec) (*model.CertifyBad, error) {
	return createUpsertForCertifyBad(&subject, pkgMatchType, &certifyBad).upsert(c)
}

func (c *gremlinClient) CertifyBad(ctx context.Context, certifyBadSpec *model.CertifyBadSpec) ([]*model.CertifyBad, error) {
	q := createQueryForEdge[*model.CertifyBad](SubjectToCertifyBad).
		withInVertex(createQueryForVertex[*model.CertifyBad](CertifyBad).
			withId(certifyBadSpec.ID).
			withPropString(justification, certifyBadSpec.Justification).
			withPropString(origin, certifyBadSpec.Origin).
			withPropString(collector, certifyBadSpec.Collector)).
		withMapper(getCertifyBadFromEdge)
	if certifyBadSpec.Subject != nil {
		q = q.withOutVertex(createQueryToMatchPackageSourceOrArtifactSpec[*model.CertifyBad](certifyBadSpec.Subject))
	}
	return q.findAllEdges(c)
}
