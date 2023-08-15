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
	cveId                string = "cveId"
	osvId                string = "osvId"
	ghsaId               string = "ghsaId"
	guacEmpty            string = "guac-empty-@@"
	justification        string = "justification"
	guacPartitionKey     string = "guac-partition-key"
	guacPkgVersionKey    string = "guac-pkg-version-key"
	guacSecondPkgNameKey string = "guac-second-pkg-name-key"
	knownSince           string = "knownSince"
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

func (c *gremlinClient) IngestCertifyBads(ctx context.Context, subjects model.PackageSourceOrArtifactInputs, pkgMatchType *model.MatchFlags, certifyBads []*model.CertifyBadInputSpec) ([]*model.CertifyBad, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestCertifyGoods(ctx context.Context, subjects model.PackageSourceOrArtifactInputs, pkgMatchType *model.MatchFlags, certifyGoods []*model.CertifyGoodInputSpec) ([]*model.CertifyGood, error) {
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

func (c *gremlinClient) HasMetadata(ctx context.Context, hasMetadataSpec *model.HasMetadataSpec) ([]*model.HasMetadata, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestHasMetadata(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, hasMetadata model.HasMetadataInputSpec) (*model.HasMetadata, error) {
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

func (c *gremlinClient) HasSlsa(ctx context.Context, hasSLSASpec *model.HasSLSASpec) ([]*model.HasSlsa, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IsVulnerability(ctx context.Context, isVulnerabilitySpec *model.IsVulnerabilitySpec) ([]*model.IsVulnerability, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) PkgEqual(ctx context.Context, pkgEqualSpec *model.PkgEqualSpec) ([]*model.PkgEqual, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestMaterials(ctx context.Context, materials []*model.ArtifactInputSpec) ([]*model.Artifact, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestCertifyGood(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, certifyGood model.CertifyGoodInputSpec) (*model.CertifyGood, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestIsVulnerability(ctx context.Context, osv model.OSVInputSpec, vulnerability model.CveOrGhsaInput, isVulnerability model.IsVulnerabilityInputSpec) (*model.IsVulnerability, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestPkgEqual(ctx context.Context, pkg model.PkgInputSpec, depPkg model.PkgInputSpec, pkgEqual model.PkgEqualInputSpec) (*model.PkgEqual, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestVEXStatement(ctx context.Context, subject model.PackageOrArtifactInput, vulnerability model.VulnerabilityInput, vexStatement model.VexStatementInputSpec) (*model.CertifyVEXStatement, error) {
	//TODO implement me
	panic("implement me")
}

func (c *gremlinClient) IngestVulnerability(ctx context.Context, pkg model.PkgInputSpec, vulnerability model.VulnerabilityInput, certifyVuln model.VulnerabilityMetaDataInput) (*model.CertifyVuln, error) {
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
