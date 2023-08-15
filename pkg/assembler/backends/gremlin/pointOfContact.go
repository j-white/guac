package gremlin

import (
	"context"
	"fmt"
	"github.com/guacsec/guac/internal/testing/ptrfrom"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"time"
)

const (
	PointOfContact          Label = "pointOfContact"
	SubjectToPointOfContact Label = "subject-to-pointOfContact"
)

func createUpsertForPointOfContactVertex(pointOfContact *model.PointOfContactInputSpec) *gremlinQueryBuilder[*model.PointOfContact] {
	return createUpsertForVertex[*model.PointOfContact](PointOfContact).
		withPropString(email, &pointOfContact.Email).
		withPropString(info, &pointOfContact.Info).
		withPropTime(since, &pointOfContact.Since).
		withPropString(justification, &pointOfContact.Justification).
		withPropString(origin, &pointOfContact.Origin).
		withPropString(collector, &pointOfContact.Collector)
}

func createUpsertForPointOfContact(subject *model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, pointOfContact *model.PointOfContactInputSpec) *gremlinQueryBuilder[*model.PointOfContact] {
	return createUpsertForEdge[*model.PointOfContact](SubjectToPointOfContact).
		withPropString("test", ptrfrom.String("asdf")).
		withInVertex(createQueryToMatchPackageSourceOrArtifactInput[*model.PointOfContact](subject, pkgMatchType)).
		withOutVertex(createUpsertForPointOfContactVertex(pointOfContact)).
		withMapper(getPointOfContactFromEdge)
}

func getPointOfContactFromEdge(result *gremlinQueryResult) (*model.PointOfContact, error) {
	pointOfContact := &model.PointOfContact{
		ID:            result.outId,
		Email:         result.out[email].(string),
		Info:          result.out[info].(string),
		Since:         result.out[since].(time.Time),
		Justification: result.out[justification].(string),
		Origin:        result.out[collector].(string),
		Collector:     result.out[origin].(string),
	}
	if result.inLabel == Package {
		pointOfContact.Subject = getPackageObject(result.inId, result.in)
	} else if result.inLabel == Source {
		pointOfContact.Subject = getSourceObject(result.inId, result.in)
	} else if result.inLabel == Artifact {
		pointOfContact.Subject = getArtifactObject(result.inId, result.in)
	} else {
		return nil, fmt.Errorf("unsupported label: %v", result.inLabel)
	}
	return pointOfContact, nil
}

func createQueryToMatchPackageInputWithMatchType[M any](pkg *model.PkgInputSpec, pkgMatchType *model.MatchFlags) *gremlinQueryBuilder[M] {
	return createQueryForVertex[M](Package).
		withPropString(typeStr, &pkg.Type).
		withPropString(name, &pkg.Name).
		withPropString(namespace, pkg.Namespace).
		withPropString(subpath, pkg.Subpath).
		withPropStringOrEmpty(version, pkg.Version)
}

func createQueryToMatchPackageSourceOrArtifactInput[M any](subject *model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags) *gremlinQueryBuilder[M] {
	if subject.Package != nil {
		return createQueryToMatchPackageInputWithMatchType[M](subject.Package, pkgMatchType)
	} else if subject.Source != nil {
		return createQueryToMatchSourceInput[M](subject.Source)
	} else if subject.Artifact != nil {
		return createQueryToMatchArtifactInput[M](subject.Artifact)
	}
	// FIXME: error handling
	return nil
}

func createQueryToMatchPackageSourceOrArtifactSpec[M any](subject *model.PackageSourceOrArtifactSpec) *gremlinQueryBuilder[M] {
	if subject.Package != nil {
		return createQueryToMatchPackage[M](subject.Package)
	} else if subject.Source != nil {
		return createQueryToMatchSource[M](subject.Source)
	} else if subject.Artifact != nil {
		return createQueryToMatchArtifact[M](subject.Artifact)
	}
	// FIXME: error handling
	return nil
}

// IngestPointOfContact
//
//	pkg,src,artifact ->subject-pointOfContact-> pointOfContact
func (c *gremlinClient) IngestPointOfContact(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, pointOfContact model.PointOfContactInputSpec) (*model.PointOfContact, error) {
	return createUpsertForPointOfContact(&subject, pkgMatchType, &pointOfContact).upsert(c)
}

func (c *gremlinClient) PointOfContact(ctx context.Context, pointOfContactSpec *model.PointOfContactSpec) ([]*model.PointOfContact, error) {
	q := createQueryForEdge[*model.PointOfContact](SubjectToPointOfContact).
		withOutVertex(createQueryForVertex[*model.PointOfContact](PointOfContact).
			withId(pointOfContactSpec.ID).
			withPropString(email, pointOfContactSpec.Email).
			withPropString(info, pointOfContactSpec.Info).
			withPropTimeGreaterOrEqual(since, pointOfContactSpec.Since).
			withPropString(justification, pointOfContactSpec.Justification).
			withPropString(origin, pointOfContactSpec.Origin).
			withPropString(collector, pointOfContactSpec.Collector)).
		withMapper(getPointOfContactFromEdge)
	if pointOfContactSpec.Subject != nil {
		q = q.withInVertex(createQueryToMatchPackageSourceOrArtifactSpec[*model.PointOfContact](pointOfContactSpec.Subject))
	}
	return q.findAll(c)
}
