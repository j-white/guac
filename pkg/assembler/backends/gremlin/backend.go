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

package gremlin

import (
	"context"
	"crypto/tls"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/guacsec/guac/pkg/assembler/backends"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/guacsec/guac/pkg/logging"
	"go.uber.org/zap"
	"strings"
)

const (
	algorithm            string = "algorithm"
	digest               string = "digest"
	typeStr              string = "type"
	uri                  string = "uri"
	year                 string = "year"
	cveId                string = "cve-id"
	osvId                string = "osv-id"
	ghsaId               string = "ghsa-id"
	guacEmpty            string = "guac-empty-@@"
	justification        string = "justification"
	guacPartitionKey     string = "guac-partition-key"
	guacPkgVersionKey    string = "guac-pkg-version-key"
	guacSecondPkgNameKey string = "guac-second-pkg-name-key"
	knownSince           string = "known-since"
	email                string = "email"
	since                string = "since"
	info                 string = "info"
	key                  string = "guak-key"
	value                string = "value"
	timestamp            string = "timestamp"
	vulnerabilityId      string = "vulnerabilityId"
)

type Flavor int64

const (
	JanusGraph Flavor = iota
	Neptune
	CosmosDB
)

type GremlinConfig struct {
	Flavor             Flavor
	Url                string
	MaxResultsPerQuery uint32
	// auth
	Username string
	Password string
	// tls
	InsecureTLSSkipVerify bool
}

type gremlinClient struct {
	config GremlinConfig
	remote *gremlingo.DriverRemoteConnection
}

func (c *gremlinClient) Artifacts(ctx context.Context, artifactSpec *model.ArtifactSpec) ([]*model.Artifact, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) Builders(ctx context.Context, builderSpec *model.BuilderSpec) ([]*model.Builder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) Licenses(ctx context.Context, licenseSpec *model.LicenseSpec) ([]*model.License, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) Packages(ctx context.Context, pkgSpec *model.PkgSpec) ([]*model.Package, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) Sources(ctx context.Context, sourceSpec *model.SourceSpec) ([]*model.Source, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) Vulnerabilities(ctx context.Context, vulnSpec *model.VulnerabilitySpec) ([]*model.Vulnerability, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) CertifyBad(ctx context.Context, certifyBadSpec *model.CertifyBadSpec) ([]*model.CertifyBad, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) CertifyGood(ctx context.Context, certifyGoodSpec *model.CertifyGoodSpec) ([]*model.CertifyGood, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) CertifyVEXStatement(ctx context.Context, certifyVEXStatementSpec *model.CertifyVEXStatementSpec) ([]*model.CertifyVEXStatement, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) CertifyVuln(ctx context.Context, certifyVulnSpec *model.CertifyVulnSpec) ([]*model.CertifyVuln, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) CertifyLegal(ctx context.Context, certifyLegalSpec *model.CertifyLegalSpec) ([]*model.CertifyLegal, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) HasSBOM(ctx context.Context, hasSBOMSpec *model.HasSBOMSpec) ([]*model.HasSbom, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) HasSlsa(ctx context.Context, hasSLSASpec *model.HasSLSASpec) ([]*model.HasSlsa, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) HasSourceAt(ctx context.Context, hasSourceAtSpec *model.HasSourceAtSpec) ([]*model.HasSourceAt, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) HasMetadata(ctx context.Context, hasMetadataSpec *model.HasMetadataSpec) ([]*model.HasMetadata, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) HashEqual(ctx context.Context, hashEqualSpec *model.HashEqualSpec) ([]*model.HashEqual, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IsDependency(ctx context.Context, isDependencySpec *model.IsDependencySpec) ([]*model.IsDependency, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IsOccurrence(ctx context.Context, isOccurrenceSpec *model.IsOccurrenceSpec) ([]*model.IsOccurrence, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) PkgEqual(ctx context.Context, pkgEqualSpec *model.PkgEqualSpec) ([]*model.PkgEqual, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) PointOfContact(ctx context.Context, pointOfContactSpec *model.PointOfContactSpec) ([]*model.PointOfContact, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) Scorecards(ctx context.Context, certifyScorecardSpec *model.CertifyScorecardSpec) ([]*model.CertifyScorecard, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) VulnEqual(ctx context.Context, vulnEqualSpec *model.VulnEqualSpec) ([]*model.VulnEqual, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) VulnerabilityMetadata(ctx context.Context, vulnerabilityMetadataSpec *model.VulnerabilityMetadataSpec) ([]*model.VulnerabilityMetadata, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestArtifact(ctx context.Context, artifact *model.ArtifactInputSpec) (*model.Artifact, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestArtifacts(ctx context.Context, artifacts []*model.ArtifactInputSpec) ([]*model.Artifact, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestBuilder(ctx context.Context, builder *model.BuilderInputSpec) (*model.Builder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestBuilders(ctx context.Context, builders []*model.BuilderInputSpec) ([]*model.Builder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestLicense(ctx context.Context, license *model.LicenseInputSpec) (*model.License, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestLicenses(ctx context.Context, licenses []*model.LicenseInputSpec) ([]*model.License, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestPackage(ctx context.Context, pkg model.PkgInputSpec) (*model.Package, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestPackages(ctx context.Context, pkgs []*model.PkgInputSpec) ([]*model.Package, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestSource(ctx context.Context, source model.SourceInputSpec) (*model.Source, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestSources(ctx context.Context, sources []*model.SourceInputSpec) ([]*model.Source, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestVulnerability(ctx context.Context, vuln model.VulnerabilityInputSpec) (*model.Vulnerability, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestVulnerabilities(ctx context.Context, vulns []*model.VulnerabilityInputSpec) ([]*model.Vulnerability, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestCertifyBad(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, certifyBad model.CertifyBadInputSpec) (*model.CertifyBad, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestCertifyBads(ctx context.Context, subjects model.PackageSourceOrArtifactInputs, pkgMatchType *model.MatchFlags, certifyBads []*model.CertifyBadInputSpec) ([]*model.CertifyBad, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestCertifyGood(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, certifyGood model.CertifyGoodInputSpec) (*model.CertifyGood, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestCertifyGoods(ctx context.Context, subjects model.PackageSourceOrArtifactInputs, pkgMatchType *model.MatchFlags, certifyGoods []*model.CertifyGoodInputSpec) ([]*model.CertifyGood, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestCertifyVuln(ctx context.Context, pkg model.PkgInputSpec, vulnerability model.VulnerabilityInputSpec, certifyVuln model.ScanMetadataInput) (*model.CertifyVuln, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestCertifyVulns(ctx context.Context, pkgs []*model.PkgInputSpec, vulnerabilities []*model.VulnerabilityInputSpec, certifyVulns []*model.ScanMetadataInput) ([]*model.CertifyVuln, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestCertifyLegal(ctx context.Context, subject model.PackageOrSourceInput, declaredLicenses []*model.LicenseInputSpec, discoveredLicenses []*model.LicenseInputSpec, certifyLegal *model.CertifyLegalInputSpec) (*model.CertifyLegal, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestCertifyLegals(ctx context.Context, subjects model.PackageOrSourceInputs, declaredLicensesList [][]*model.LicenseInputSpec, discoveredLicensesList [][]*model.LicenseInputSpec, certifyLegals []*model.CertifyLegalInputSpec) ([]*model.CertifyLegal, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestDependency(ctx context.Context, pkg model.PkgInputSpec, depPkg model.PkgInputSpec, depPkgMatchType model.MatchFlags, dependency model.IsDependencyInputSpec) (*model.IsDependency, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestDependencies(ctx context.Context, pkgs []*model.PkgInputSpec, depPkgs []*model.PkgInputSpec, depPkgMatchType model.MatchFlags, dependencies []*model.IsDependencyInputSpec) ([]*model.IsDependency, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestHasSbom(ctx context.Context, subject model.PackageOrArtifactInput, hasSbom model.HasSBOMInputSpec, includes model.HasSBOMIncludesInputSpec) (*model.HasSbom, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestHasSBOMs(ctx context.Context, subjects model.PackageOrArtifactInputs, hasSBOMs []*model.HasSBOMInputSpec, includes []*model.HasSBOMIncludesInputSpec) ([]*model.HasSbom, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestHasSourceAt(ctx context.Context, pkg model.PkgInputSpec, pkgMatchType model.MatchFlags, source model.SourceInputSpec, hasSourceAt model.HasSourceAtInputSpec) (*model.HasSourceAt, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestHasSourceAts(ctx context.Context, pkgs []*model.PkgInputSpec, pkgMatchType *model.MatchFlags, sources []*model.SourceInputSpec, hasSourceAts []*model.HasSourceAtInputSpec) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestHasMetadata(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, hasMetadata model.HasMetadataInputSpec) (*model.HasMetadata, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestBulkHasMetadata(ctx context.Context, subjects model.PackageSourceOrArtifactInputs, pkgMatchType *model.MatchFlags, hasMetadataList []*model.HasMetadataInputSpec) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestHashEqual(ctx context.Context, artifact model.ArtifactInputSpec, equalArtifact model.ArtifactInputSpec, hashEqual model.HashEqualInputSpec) (*model.HashEqual, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestHashEquals(ctx context.Context, artifacts []*model.ArtifactInputSpec, otherArtifacts []*model.ArtifactInputSpec, hashEquals []*model.HashEqualInputSpec) ([]*model.HashEqual, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestOccurrence(ctx context.Context, subject model.PackageOrSourceInput, artifact model.ArtifactInputSpec, occurrence model.IsOccurrenceInputSpec) (*model.IsOccurrence, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestOccurrences(ctx context.Context, subjects model.PackageOrSourceInputs, artifacts []*model.ArtifactInputSpec, occurrences []*model.IsOccurrenceInputSpec) ([]*model.IsOccurrence, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestPkgEqual(ctx context.Context, pkg model.PkgInputSpec, depPkg model.PkgInputSpec, pkgEqual model.PkgEqualInputSpec) (*model.PkgEqual, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestPkgEquals(ctx context.Context, pkgs []*model.PkgInputSpec, otherPackages []*model.PkgInputSpec, pkgEquals []*model.PkgEqualInputSpec) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestPointOfContact(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, pointOfContact model.PointOfContactInputSpec) (*model.PointOfContact, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestPointOfContacts(ctx context.Context, subjects model.PackageSourceOrArtifactInputs, pkgMatchType *model.MatchFlags, pointOfContacts []*model.PointOfContactInputSpec) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestSLSA(ctx context.Context, subject model.ArtifactInputSpec, builtFrom []*model.ArtifactInputSpec, builtBy model.BuilderInputSpec, slsa model.SLSAInputSpec) (*model.HasSlsa, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestSLSAs(ctx context.Context, subjects []*model.ArtifactInputSpec, builtFromList [][]*model.ArtifactInputSpec, builtByList []*model.BuilderInputSpec, slsaList []*model.SLSAInputSpec) ([]*model.HasSlsa, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestScorecard(ctx context.Context, source model.SourceInputSpec, scorecard model.ScorecardInputSpec) (*model.CertifyScorecard, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestScorecards(ctx context.Context, sources []*model.SourceInputSpec, scorecards []*model.ScorecardInputSpec) ([]*model.CertifyScorecard, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestVEXStatement(ctx context.Context, subject model.PackageOrArtifactInput, vulnerability model.VulnerabilityInputSpec, vexStatement model.VexStatementInputSpec) (*model.CertifyVEXStatement, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestVEXStatements(ctx context.Context, subjects model.PackageOrArtifactInputs, vulnerabilities []*model.VulnerabilityInputSpec, vexStatements []*model.VexStatementInputSpec) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestVulnEqual(ctx context.Context, vulnerability model.VulnerabilityInputSpec, otherVulnerability model.VulnerabilityInputSpec, vulnEqual model.VulnEqualInputSpec) (*model.VulnEqual, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestVulnEquals(ctx context.Context, vulnerabilities []*model.VulnerabilityInputSpec, otherVulnerabilities []*model.VulnerabilityInputSpec, vulnEquals []*model.VulnEqualInputSpec) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestVulnerabilityMetadata(ctx context.Context, vulnerability model.VulnerabilityInputSpec, vulnerabilityMetadata model.VulnerabilityMetadataInputSpec) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestBulkVulnerabilityMetadata(ctx context.Context, vulnerabilities []*model.VulnerabilityInputSpec, vulnerabilityMetadataList []*model.VulnerabilityMetadataInputSpec) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) Neighbors(ctx context.Context, node string, usingOnly []model.Edge) ([]model.Node, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) Node(ctx context.Context, node string) (model.Node, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) Nodes(ctx context.Context, nodes []string) ([]model.Node, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) Path(ctx context.Context, subject string, target string, maxPathLength int, usingOnly []model.Edge) ([]model.Node, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) FindSoftware(ctx context.Context, searchText string) ([]model.PackageSourceOrArtifact, error) {
	//TODO implement me
	panic("implement me")
}

type gremlinLogger struct {
	logger *zap.SugaredLogger
}

func (l *gremlinLogger) Log(_ gremlingo.LogVerbosity, v ...interface{}) {
	l.logger.Info(v...)
}

func (l *gremlinLogger) Logf(_ gremlingo.LogVerbosity, format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}

func (c *gremlinClient) Close() {
	// FIXME: Is there a place to call this in the backends.Backend lifecycle?
	c.remote.Close()
}

func GetBackend(args backends.BackendArgs) (backends.Backend, error) {
	ctx := logging.WithLogger(context.Background())
	logger := logging.FromContext(ctx)

	// sanitize config
	config := args.(*GremlinConfig)
	if config.MaxResultsPerQuery < 1 {
		config.MaxResultsPerQuery = 1000
	}

	remote, err := gremlingo.NewDriverRemoteConnection(config.Url, func(settings *gremlingo.DriverRemoteConnectionSettings) {
		// route everything to our logger
		settings.Logger = &gremlinLogger{logger: logger}
		settings.LogVerbosity = gremlingo.Debug
		if strings.TrimSpace(config.Username) != "" {
			settings.AuthInfo = &gremlingo.AuthInfo{
				Username: config.Username,
				Password: config.Password,
			}
		}
		settings.TlsConfig = &tls.Config{
			InsecureSkipVerify: config.InsecureTLSSkipVerify,
		}
		//// more knobs for performance tuning comms and the connection pool
		// settings.EnableCompression = true
		// settings.NewConnectionThreshold = 1
		// settings.InitialConcurrentConnections = 2
		// settings.MaximumConcurrentConnections = 8
	})
	if err != nil {
		logger.Errorf("Failed to initialise connection to Gremlin server at URL: %s. Error: %v", config.Url, err)
		return nil, err
	}
	logger.Infof("Succesfully connected to Gremlin server at URL: %s", config.Url)

	// Verify that transactions are supported by the underlying graph engine
	transactionsSupported, err := supportsTransactions(remote)
	if err != nil {
		logger.Errorf("Failed to verify if server supports transactions: %v", err)
		return nil, err
	}
	logger.Infof("Gremlin server supports transactions: %v", transactionsSupported)

	if config.Flavor == JanusGraph {
		// edge ids are custom types, add support for deserializing them
		registerCustomTypeReadersForJanusGraph()
		// schema
		err := createIndicesForJanusGraph(ctx, remote)
		if err != nil {
			logger.Errorf("Failed to initialise schema for JanusGraph at URL: %s. Error: %v", config.Url, err)
			return nil, err
		}
	} else {
		logger.Warn("Let's see what happens with: %s", config.Flavor)
	}

	client := &gremlinClient{*config, remote}
	return client, nil
}
