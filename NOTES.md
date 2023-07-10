# IngestArtifact

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

# Query artifact

query {
  artifacts(artifactSpec:{id:"2"}){
    id,
    digest
  }
}

