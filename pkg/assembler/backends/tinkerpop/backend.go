//
// Copyright 2023 The GUAC Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tinkerpop

import (
	"context"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/guacsec/guac/pkg/assembler/backends"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/guacsec/guac/pkg/logging"
)

type TinkerPopConfig struct {
	SettingsFile string
	MaxLimit     uint32
}

type tinkerpopClient struct {
	config TinkerPopConfig
	remote *gremlingo.DriverRemoteConnection
}

func (c *tinkerpopClient) IngestSLSAs(ctx context.Context, subjects []*model.ArtifactInputSpec, builtFromList [][]*model.ArtifactInputSpec, builtByList []*model.BuilderInputSpec, slsaList []*model.SLSAInputSpec) ([]*model.HasSlsa, error) {
	//TODO implement me
	panic("implement me")
}

func GetBackend(args backends.BackendArgs) (backends.Backend, error) {
	ctx := logging.WithLogger(context.Background())
	logger := logging.FromContext(ctx)

	config := args.(*TinkerPopConfig)
	// FIXME: Make this configurable
	config.MaxLimit = 1000

	// FIXME: Is there no clean shutdown of the backend?
	remote, err := gremlingo.NewDriverRemoteConnection("ws://janusgraph:8182/gremlin")
	if err != nil {
		return nil, err
	}
	logger.Infof("Succesfully connected to Gremlin server.")

	// Verify that transactions are supported by the underlying graph engine
	// FIXME: Is there a cleaner way to check this, something like: graph.features().graph().supportsTransactions()
	g := gremlingo.Traversal_().WithRemote(remote)
	tx := g.Tx()
	gtx, err := tx.Begin()
	if err != nil {
		logger.Errorf("Failed to create transaction: %v", err)
		return nil, err
	}
	promise := gtx.AddV("x").Property("a", "b").Iterate()
	err = <-promise
	if err != nil {
		return nil, err
	}
	err = tx.Rollback()
	if err != nil {
		logger.Errorf("Failed to rollback transaction: %v", err)
		return nil, err
	}

	client := &tinkerpopClient{*config, remote}
	return client, nil
}

func (c *tinkerpopClient) IngestPackage(ctx context.Context, pkg model.PkgInputSpec) (*model.Package, error) {
	g := gremlingo.Traversal_().WithRemote(c.remote)

	g.AddV(pkg.Name).Property("pkgType", pkg.Type)

	return nil, nil
	/*
			// Perform traversal
			result, err := g.V().HasLabel("person").Has("age", __.Is(gt(28))).Order().By("age", order.Desc).Values("name").ToList()

			session := c.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
			defer session.Close()

			values := map[string]any{}
			values["pkgType"] = pkg.Type
			values["name"] = pkg.Name
			if pkg.Namespace != nil {
				values["namespace"] = *pkg.Namespace
			} else {
				values["namespace"] = ""
			}
			if pkg.Version != nil {
				values["version"] = *pkg.Version
			} else {
				values["version"] = ""
			}
			if pkg.Subpath != nil {
				values["subpath"] = *pkg.Subpath
			} else {
				values["subpath"] = ""
			}

			// To ensure consistency, always sort the qualifiers by key
			qualifiersMap := map[string]string{}
			keys := []string{}
			for _, kv := range pkg.Qualifiers {
				qualifiersMap[kv.Key] = kv.Value
				keys = append(keys, kv.Key)
			}
			sort.Strings(keys)
			qualifiers := []string{}
			for _, k := range keys {
				qualifiers = append(qualifiers, k, qualifiersMap[k])
			}
			values["qualifier"] = qualifiers

			result, err := session.WriteTransaction(
				func(tx neo4j.Transaction) (interface{}, error) {
					query := `MERGE (root:Pkg)
		MERGE (root) -[:PkgHasType]-> (type:PkgType{type:$pkgType})
		MERGE (type) -[:PkgHasNamespace]-> (ns:PkgNamespace{namespace:$namespace})
		MERGE (ns) -[:PkgHasName]-> (name:PkgName{name:$name})
		MERGE (name) -[:PkgHasVersion]-> (version:PkgVersion{version:$version,subpath:$subpath,qualifier_list:$qualifier})
		RETURN type.type, ns.namespace, name.name, version.version, version.subpath, version.qualifier_list`
					result, err := tx.Run(query, values)
					if err != nil {
						return nil, err
					}

					// query returns a single record
					record, err := result.Single()
					if err != nil {
						return nil, err
					}

					qualifiersList := record.Values[5]
					subPath := record.Values[4]
					version := record.Values[3]
					nameStr := record.Values[2].(string)
					namespaceStr := record.Values[1].(string)
					pkgType := record.Values[0].(string)

					pkg := generateModelPackage(pkgType, namespaceStr, nameStr, version, subPath, qualifiersList)
					return pkg, nil
				})
			if err != nil {
				return nil, err
			}

			return result.(*model.Package), nil
	*/
}

func (c *tinkerpopClient) HasMetadata(ctx context.Context, hasMetadataSpec *model.HasMetadataSpec) ([]*model.HasMetadata, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestHasMetadata(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, hasMetadata model.HasMetadataInputSpec) (*model.HasMetadata, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) Builders(ctx context.Context, builderSpec *model.BuilderSpec) ([]*model.Builder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) Cve(ctx context.Context, cveSpec *model.CVESpec) ([]*model.Cve, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) Ghsa(ctx context.Context, ghsaSpec *model.GHSASpec) ([]*model.Ghsa, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) Osv(ctx context.Context, osvSpec *model.OSVSpec) ([]*model.Osv, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) Packages(ctx context.Context, pkgSpec *model.PkgSpec) ([]*model.Package, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) CertifyGood(ctx context.Context, certifyGoodSpec *model.CertifyGoodSpec) ([]*model.CertifyGood, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) CertifyVEXStatement(ctx context.Context, certifyVEXStatementSpec *model.CertifyVEXStatementSpec) ([]*model.CertifyVEXStatement, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) CertifyVuln(ctx context.Context, certifyVulnSpec *model.CertifyVulnSpec) ([]*model.CertifyVuln, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) HasSBOM(ctx context.Context, hasSBOMSpec *model.HasSBOMSpec) ([]*model.HasSbom, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) HasSlsa(ctx context.Context, hasSLSASpec *model.HasSLSASpec) ([]*model.HasSlsa, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) HasSourceAt(ctx context.Context, hasSourceAtSpec *model.HasSourceAtSpec) ([]*model.HasSourceAt, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) HashEqual(ctx context.Context, hashEqualSpec *model.HashEqualSpec) ([]*model.HashEqual, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IsDependency(ctx context.Context, isDependencySpec *model.IsDependencySpec) ([]*model.IsDependency, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IsOccurrence(ctx context.Context, isOccurrenceSpec *model.IsOccurrenceSpec) ([]*model.IsOccurrence, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IsVulnerability(ctx context.Context, isVulnerabilitySpec *model.IsVulnerabilitySpec) ([]*model.IsVulnerability, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) PkgEqual(ctx context.Context, pkgEqualSpec *model.PkgEqualSpec) ([]*model.PkgEqual, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestBuilder(ctx context.Context, builder *model.BuilderInputSpec) (*model.Builder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestCve(ctx context.Context, cve *model.CVEInputSpec) (*model.Cve, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestGhsa(ctx context.Context, ghsa *model.GHSAInputSpec) (*model.Ghsa, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestMaterials(ctx context.Context, materials []*model.ArtifactInputSpec) ([]*model.Artifact, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestOsv(ctx context.Context, osv *model.OSVInputSpec) (*model.Osv, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestPackages(ctx context.Context, pkgs []*model.PkgInputSpec) ([]*model.Package, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestCertifyBad(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, certifyBad model.CertifyBadInputSpec) (*model.CertifyBad, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestCertifyGood(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, certifyGood model.CertifyGoodInputSpec) (*model.CertifyGood, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestDependency(ctx context.Context, pkg model.PkgInputSpec, depPkg model.PkgInputSpec, dependency model.IsDependencyInputSpec) (*model.IsDependency, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestDependencies(ctx context.Context, pkgs []*model.PkgInputSpec, depPkgs []*model.PkgInputSpec, dependencies []*model.IsDependencyInputSpec) ([]*model.IsDependency, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestHasSbom(ctx context.Context, subject model.PackageOrArtifactInput, hasSbom model.HasSBOMInputSpec) (*model.HasSbom, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestHasSourceAt(ctx context.Context, pkg model.PkgInputSpec, pkgMatchType model.MatchFlags, source model.SourceInputSpec, hasSourceAt model.HasSourceAtInputSpec) (*model.HasSourceAt, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestHashEqual(ctx context.Context, artifact model.ArtifactInputSpec, equalArtifact model.ArtifactInputSpec, hashEqual model.HashEqualInputSpec) (*model.HashEqual, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestIsVulnerability(ctx context.Context, osv model.OSVInputSpec, vulnerability model.CveOrGhsaInput, isVulnerability model.IsVulnerabilityInputSpec) (*model.IsVulnerability, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestOccurrence(ctx context.Context, subject model.PackageOrSourceInput, artifact model.ArtifactInputSpec, occurrence model.IsOccurrenceInputSpec) (*model.IsOccurrence, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestOccurrences(ctx context.Context, subjects model.PackageOrSourceInputs, artifacts []*model.ArtifactInputSpec, occurrences []*model.IsOccurrenceInputSpec) ([]*model.IsOccurrence, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestPkgEqual(ctx context.Context, pkg model.PkgInputSpec, depPkg model.PkgInputSpec, pkgEqual model.PkgEqualInputSpec) (*model.PkgEqual, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestVEXStatement(ctx context.Context, subject model.PackageOrArtifactInput, vulnerability model.VulnerabilityInput, vexStatement model.VexStatementInputSpec) (*model.CertifyVEXStatement, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestVulnerability(ctx context.Context, pkg model.PkgInputSpec, vulnerability model.VulnerabilityInput, certifyVuln model.VulnerabilityMetaDataInput) (*model.CertifyVuln, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) Neighbors(ctx context.Context, node string, usingOnly []model.Edge) ([]model.Node, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) Node(ctx context.Context, node string) (model.Node, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) Nodes(ctx context.Context, nodes []string) ([]model.Node, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) Path(ctx context.Context, subject string, target string, maxPathLength int, usingOnly []model.Edge) ([]model.Node, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) FindSoftware(ctx context.Context, searchText string) ([]model.PackageSourceOrArtifact, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) CertifyBad(ctx context.Context, certifyBadSpec *model.CertifyBadSpec) ([]*model.CertifyBad, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestCVEs(ctx context.Context, cves []*model.CVEInputSpec) ([]*model.Cve, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestGHSAs(ctx context.Context, ghsas []*model.GHSAInputSpec) ([]*model.Ghsa, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestOSVs(ctx context.Context, osvs []*model.OSVInputSpec) ([]*model.Osv, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) PointOfContact(ctx context.Context, pointOfContactSpec *model.PointOfContactSpec) ([]*model.PointOfContact, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestBuilders(ctx context.Context, builders []*model.BuilderInputSpec) ([]*model.Builder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestPointOfContact(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, pointOfContact model.PointOfContactInputSpec) (*model.PointOfContact, error) {
	//TODO implement me
	panic("implement me")
}
