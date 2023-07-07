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
	"github.com/guacsec/guac/pkg/assembler/backends"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

type TinkerPopConfig struct {
	SettingsFile string
}

func GetBackend(args backends.BackendArgs) (backends.Backend, error) {
	config := args.(*TinkerPopConfig)
	client := &tinkerpopClient{*config}
	return client, nil
}

type tinkerpopClient struct {
	driver TinkerPopConfig
}

func (c *tinkerpopClient) HasMetadata(ctx context.Context, hasMetadataSpec *model.HasMetadataSpec) ([]*model.HasMetadata, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestHasMetadata(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, hasMetadata model.HasMetadataInputSpec) (*model.HasMetadata, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) Artifacts(ctx context.Context, artifactSpec *model.ArtifactSpec) ([]*model.Artifact, error) {
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

func (c *tinkerpopClient) Sources(ctx context.Context, sourceSpec *model.SourceSpec) ([]*model.Source, error) {
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

func (c *tinkerpopClient) Scorecards(ctx context.Context, certifyScorecardSpec *model.CertifyScorecardSpec) ([]*model.CertifyScorecard, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestArtifact(ctx context.Context, artifact *model.ArtifactInputSpec) (*model.Artifact, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestArtifacts(ctx context.Context, artifacts []*model.ArtifactInputSpec) ([]*model.Artifact, error) {
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

func (c *tinkerpopClient) IngestPackage(ctx context.Context, pkg model.PkgInputSpec) (*model.Package, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestPackages(ctx context.Context, pkgs []*model.PkgInputSpec) ([]*model.Package, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestSource(ctx context.Context, source model.SourceInputSpec) (*model.Source, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) CertifyScorecard(ctx context.Context, source model.SourceInputSpec, scorecard model.ScorecardInputSpec) (*model.CertifyScorecard, error) {
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

func (c *tinkerpopClient) IngestSLSA(ctx context.Context, subject model.ArtifactInputSpec, builtFrom []*model.ArtifactInputSpec, builtBy model.BuilderInputSpec, slsa model.SLSAInputSpec) (*model.HasSlsa, error) {
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
