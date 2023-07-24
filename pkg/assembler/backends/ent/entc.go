//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	_, err := entgql.NewExtension()
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	//if err := entc.Generate("./schema", &gen.Config{Features: []gen.Feature{gen.FeatureUpsert}}); err != nil {
	//	log.Fatalf("running ent codegen: %v", err)
	//}

	storage, err := gen.NewStorage("gremlin")
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	if err := entc.Generate("./schema", &gen.Config{
		//Target:   "gremlin",
		//Package:  "gremlin",
		Storage:  storage,
		Features: []gen.Feature{gen.FeatureUpsert}}); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
