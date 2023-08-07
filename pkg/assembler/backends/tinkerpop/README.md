# This doc

Misc. notes captured while exploring Guac & Gremlin

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

# Connecting to local JanusGraph 

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

MergeE
```
gremlin> :> g.mergeV([(T.label):'Dog',namespace:'Toby']).as('Toby').mergeV([(T.label):'Dog',namespace:'Brandy']).as('Brandy').mergeE([(T.label):'Sibling',created:'2022-02-07',(from):Merge.outV,(to):Merge.inV]).option(Merge.outV, select('Toby')).option(Merge.inV, select('Brandy')).as('edge').id().toList()
```

MergeE
```
gremlin> :> g.mergeV([(T.label):'Dog',namespace:'Toby']).as('Toby').mergeV([(T.label):'Dog',namespace:'Brandy']).as('Brandy').mergeE([(T.label):'Sibling',created:'2022-02-07',(from):Merge.outV,(to):Merge.inV]).option(Merge.outV, select('Toby')).option(Merge.inV, select('Brandy')).as('edge').mergeV([(T.label):'Dog',namespace:'Toby']).as('Toby2').mergeV([(T.label):'Dog',namespace:'Brandy']).as('Brandy2').mergeE([(T.label):'Sibling',created:'2022-02-07',(from):Merge.outV,(to):Merge.inV]).option(Merge.outV, select('Toby2')).option(Merge.inV, select('Brandy2')).as('edge2').select('edge','edge2').unfold().id().toList()
```

## Tuning

Increase frame size for larger bulk inserts

```
- name: gremlinserver.maxContentLength
value: "6553600"
```

# Demo

* Get local environment up with Tilt
* Launch into GraphQL playground
* Enter Gremlin console, show empty graph: `:> g.V()`
* GraphQL playground: upsert and search
* Add `Dedup()` to query, in `scorecard.go:215`, wait for reload, query, error

# Backends

## JanusGraph

* Used in Tilt environment
* Crumbles under concurrent writes

## AWS Neptune

* Neptune has no public endpoints, only available w/ local IP on VPC
* To expose it, create an LB in EC2 to point to port 8182 on the DB and allow the port in a SG
* Differs from JanusGraph in that the IDs returned are strings for vertices and edges, no custom type reader needed
* Notebooks are awesome

## Azure Cosmos DB

Config looks like:
```
config := &TinkerPopConfig{
    Url:      "wss://obfuscated.gremlin.cosmos.azure.com:443/",
    Username: "/dbs/<YOUR_DATABASE>/colls/<YOUR_COLLECTION_OR_GRAPH>",
    Password: "CJuLAERdeadbeef==",
}
```
* Need to set a partition key, must be non-null in inserts and queries
* Gremlin Bytecode is not supported, only JSON, see https://learn.microsoft.com/en-us/azure/cosmos-db/gremlin/support
* Transactions are not supported

> Won't work with gremlin-go
