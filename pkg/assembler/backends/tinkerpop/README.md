
# Observations

For every model object:
* ingest: input spec to query / upsert criteria, model object from spec
* query: query spec to query, model object from results

> Could be a good candidate for code generation from existing models.

# Design choices

* Store nested properties as JSON
* Optional values are not set with null value
* Objects are built from model used as input to upsert, only retrieving the id, for speedy ingest
* Global limit set for search queries

# Compatibility matrix

Things that support TinkerPop/Gremlin:
 * Amazon Neptune: https://docs.aws.amazon.com/neptune/latest/userguide/access-graph-gremlin-differences.html
 * Azure Cosmos DB: https://learn.microsoft.com/en-us/azure/cosmos-db/gremlin/support
 * JanusGraph: https://docs.janusgraph.org/ (Elasticsearch + ScyllaDB or Cassandra)
 * DSE Graph: https://docs.datastax.com/en/dse/5.1/dse-dev/datastax_enterprise/graph/using/insertDataGremlin.html
 * ArangoDB: https://github.com/ArangoDB-Community/arangodb-tinkerpop-provider
 * Aerospike: https://docs.aerospike.com/graph/overview

Currently developing against `JanusGraph` with the intent to expand and test others.

# GraphQL queries used in development

## IngestArtifact

```graphql
mutation {
  ingestArtifact(artifact: {algorithm: "andy", digest: "hope is a good thing"}) {
    id
    algorithm
    digest
  }
}
```

## Query artifact

```graphql
{
  artifacts(artifactSpec: {id: "2"}) {
    id
    digest
  }
}
```

## CertifyScorecard
```graphql
mutation {
  certifyScorecard(
    scorecard: {checks: [{check: "Binary_Artifacts", score: 4}], aggregateScore: 2.9, timeScanned: "2023-07-14T01:45:31.29Z", scorecardVersion: "v4.10.2", scorecardCommit: "5e6a521", origin: "Demo ingestion", collector: "sadf"}
    source: {type: "t", namespace: "asd", name: "sd", commit: "yes that is correct"}
  ) {
    id
    scorecard: id
  }
}
```

## Query scorecards

```graphql
{
  scorecards(scorecardSpec:{}){
    id,
    scorecard:id,
    source:id,
  }
}
```

# Bucket list

* Review connection pooling & lifecycle
* Add metrics
