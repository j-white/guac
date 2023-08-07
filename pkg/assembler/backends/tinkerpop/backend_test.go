package tinkerpop

import (
	"fmt"
	"testing"
)

func Test_connectToNepture(t *testing.T) {
	config := &TinkerPopConfig{
		Url: "ws://guac-neptune-db-1-1865211709.us-east-1.elb.amazonaws.com:8182/gremlin",
	}
	backend, err := GetBackend(config)
	if err != nil {
		t.Errorf("creating backend failed error = %v", err)
		return
	}

	c := backend.(*tinkerpopClient)
	schema, err := printSchema(c.remote)
	if err != nil {
		t.Errorf("printing the schema failed error = %v", err)
		return
	}
	fmt.Println(schema)
}
