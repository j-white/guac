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
	"strings"
)

const (
	algorithm     string = "algorithm"
	digest        string = "digest"
	typeStr       string = "type"
	uri           string = "uri"
	year          string = "year"
	cveId         string = "cveId"
	osvId         string = "osvId"
	ghsaId        string = "ghsaId"
	guacEmpty     string = "guac-empty-@@"
	justification string = "justification"
)

type TinkerPopConfig struct {
	Url          string
	SettingsFile string
	MaxLimit     uint32
}

type tinkerpopClient struct {
	config TinkerPopConfig
	remote *gremlingo.DriverRemoteConnection
}

func (c *tinkerpopClient) IngestCertifyBads(ctx context.Context, subjects model.PackageSourceOrArtifactInputs, pkgMatchType *model.MatchFlags, certifyBads []*model.CertifyBadInputSpec) ([]*model.CertifyBad, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestCertifyGoods(ctx context.Context, subjects model.PackageSourceOrArtifactInputs, pkgMatchType *model.MatchFlags, certifyGoods []*model.CertifyGoodInputSpec) ([]*model.CertifyGood, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestHashEquals(ctx context.Context, artifacts []*model.ArtifactInputSpec, otherArtifacts []*model.ArtifactInputSpec, hashEquals []*model.HashEqualInputSpec) ([]*model.HashEqual, error) {
	//TODO implement me
	panic("implement me")
}

func GetBackend(args backends.BackendArgs) (backends.Backend, error) {
	ctx := logging.WithLogger(context.Background())
	logger := logging.FromContext(ctx)

	registerCustomTypeReaders()

	config := args.(*TinkerPopConfig)
	// FIXME: Make this configurable
	config.MaxLimit = 1000
	if strings.TrimSpace(config.Url) == "" {
		config.Url = "ws://janusgraph:8182/gremlin"
	}

	// FIXME: Is there no clean shutdown of the backend?
	remote, err := gremlingo.NewDriverRemoteConnection(config.Url, func(settings *gremlingo.DriverRemoteConnectionSettings) {
		//// enable compression
		//settings.EnableCompression = true
		settings.NewConnectionThreshold = 1
		//settings.InitialConcurrentConnections = 2
		//settings.MaximumConcurrentConnections = 8
	})
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

	// Gremlin is meant to be generic, and some DBs will have specific APIs we need to use to initialize the schema & indices
	// This bit is JanusGraph specific
	err = createIndicesForVertexProperties(remote,
		// packages
		"namespace",
		// scorecards
		"scorecardCommit",
		// artifacts
		"digest",
		// builder
		"uri",
		// cve,
		"cveId",
		// osv,
		"osvId",
		// ghsa,
		"ghsaId",
	)
	if err != nil {
		return nil, err
	}
	// (pkg) -- (dep) --> (pkg)
	err = createIndexForEdge(remote, string(IsDependency), dependencyType)
	if err != nil {
		return nil, err
	}

	schema, err := printSchema(remote)
	logger.Info("Current JanusGraph schema: %s", schema)
	// end of JanusGraph specifics

	client := &tinkerpopClient{*config, remote}
	return client, nil
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

func (c *tinkerpopClient) IsVulnerability(ctx context.Context, isVulnerabilitySpec *model.IsVulnerabilitySpec) ([]*model.IsVulnerability, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) PkgEqual(ctx context.Context, pkgEqualSpec *model.PkgEqualSpec) ([]*model.PkgEqual, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestMaterials(ctx context.Context, materials []*model.ArtifactInputSpec) ([]*model.Artifact, error) {
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

func (c *tinkerpopClient) PointOfContact(ctx context.Context, pointOfContactSpec *model.PointOfContactSpec) ([]*model.PointOfContact, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestPointOfContact(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, pointOfContact model.PointOfContactInputSpec) (*model.PointOfContact, error) {
	//TODO implement me
	panic("implement me")
}
