package tinkerpop

import (
	"testing"
)

func Test_createIndices(t *testing.T) {
	config := &TinkerPopConfig{
		Url: "ws://localhost:8182/gremlin",
	}
	backend, err := GetBackend(config)
	if err != nil {
		t.Errorf("creating backend failed error = %v", err)
		return
	}

	c := backend.(*tinkerpopClient)
	err = printSchema(c.remote)
	if err != nil {
		t.Errorf("printing the schema failed error = %v", err)
		return
	}

	err = createIndexForVertexProperty(c.remote, "namespace")
	if err != nil {
		t.Errorf("creating indices failed error = %v", err)
		return
	}

	err = printSchema(c.remote)
	if err != nil {
		t.Errorf("printing the schema failed error = %v", err)
		return
	}
}
