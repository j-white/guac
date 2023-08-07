package tinkerpop

import (
	"fmt"
	"testing"
)

func disabled_Test_connectToNeptune(t *testing.T) {
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

func disabled_Test_connectToCosmos(t *testing.T) {

	config := &TinkerPopConfig{
		Url:      "wss://guac-db-1.gremlin.cosmos.azure.com:443/",
		Username: "/dbs/dbid/colls/graphid",
		Password: "CJuLAERDWWcD4Pds2k3Ccb6rshwqj0PfjgDlv4pUjLAmuEXZ7Z9O2vrrwiU2amVCCwoj257Vk5swACDbHJFYZQ==",
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
