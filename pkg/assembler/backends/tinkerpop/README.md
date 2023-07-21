
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

# JanusGraph notes

```
kubectl exec -it deploy/janusgraph ./bin/gremlin.sh
gremlin> :remote connect tinkerpop.server conf/remote.yaml
gremlin> :> g.V().values()
```

> Use `:>` as prefix to be able to access `g` that's tied to the deployment

Delete everything
```
gremlin> :> g.V().drop()
```

Debugging
```
gremlin> :>  g.V().hasLabel("scorecard").as("scorecard").out().hasLabel("source").as("source").select("scorecard","source")
==>{scorecard=v[65640], source=v[37048]}
==>{scorecard=v[147584], source=v[41080]}
```

# Demo

* Get local environment up with Tilt
* Launch into GraphQL playground
* Enter Gremlin console, show empty graph: `:> g.V()`
* GraphQL playground: upsert and search
* Add `Dedup()` to query, in `scorecard.go:215`, wait for reload, query, error
