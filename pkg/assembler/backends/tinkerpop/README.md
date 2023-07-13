
# Notes

* Optional values are stored as empty strings

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
ingestArtifact(artifact:{
algorithm: "andy",
digest: "hope is a good thing",
}) {
id,
algorithm,
digest
}
}
```

## Query artifact

```graphql
query {
artifacts(artifactSpec:{id:"2"}){
id,
digest
}
}
```

## CertifyScorecard
```graphql
mutation{
  certifyScorecard(scorecard:{
    checks:[{
    check:"a",
      score: 1
    }],
    aggregateScore: 0.1,
    timeScanned: "2023-04-03T16:28:44.835711634Z",
  	scorecardVersion: "1",
    scorecardCommit:"asdf",
    origin:"asdf",
    collector:"sadf"
  },
  source:{
    type:"t",
    namespace:"",
    name:"sd",
    commit:"yes that is correct"
  }) {
    id
  }
}
```

# Bucket list

* Review connection pooling & lifecycle
* Add metrics
