package gremlin

import (
	"context"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// IngestCertifyBad
//
//	           	   ->certifyBadPkgVersionEdgesStr|certifyBadPkgNameEdgesStr-> pkg
//		certifyBad ->certifyBadSrcEdgesStr-> src
//		           ->certifyBadArtEdgesStr-> artifact
func (c *gremlinClient) IngestCertifyBad(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, certifyBad model.CertifyBadInputSpec) (*model.CertifyBad, error) {
	// subject model.PackageSourceOrArtifactInput
	// pkgMatchType *model.MatchFlags
	// certifyBad model.CertifyBadInputSpec

	// different types of edges based on the subject
	//  pkg
	//     certifyBadPkgVersionEdgesStr
	//     certifyBadPkgNameEdgesStr
	//  src
	//     certifyBadSrcEdgesStr
	//  artifact
	//     certifyBadArtEdgesStr

	return nil, nil
}

func (c *gremlinClient) CertifyBad(ctx context.Context, certifyBadSpec *model.CertifyBadSpec) ([]*model.CertifyBad, error) {
	return nil, nil
}
