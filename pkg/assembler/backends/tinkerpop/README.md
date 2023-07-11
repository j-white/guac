
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

# Bucket list

* Review connection pooling & lifecycle
* Add metrics
