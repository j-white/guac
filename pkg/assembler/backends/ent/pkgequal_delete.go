// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/pkgequal"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// PkgEqualDelete is the builder for deleting a PkgEqual entity.
type PkgEqualDelete struct {
	config
	hooks    []Hook
	mutation *PkgEqualMutation
}

// Where appends a list predicates to the PkgEqualDelete builder.
func (ped *PkgEqualDelete) Where(ps ...predicate.PkgEqual) *PkgEqualDelete {
	ped.mutation.Where(ps...)
	return ped
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ped *PkgEqualDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ped.gremlinExec, ped.mutation, ped.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ped *PkgEqualDelete) ExecX(ctx context.Context) int {
	n, err := ped.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ped *PkgEqualDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := ped.gremlin().Query()
	if err := ped.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	ped.mutation.done = true
	return res.ReadInt()
}

func (ped *PkgEqualDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(pkgequal.Label)
	for _, p := range ped.mutation.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// PkgEqualDeleteOne is the builder for deleting a single PkgEqual entity.
type PkgEqualDeleteOne struct {
	ped *PkgEqualDelete
}

// Where appends a list predicates to the PkgEqualDelete builder.
func (pedo *PkgEqualDeleteOne) Where(ps ...predicate.PkgEqual) *PkgEqualDeleteOne {
	pedo.ped.mutation.Where(ps...)
	return pedo
}

// Exec executes the deletion query.
func (pedo *PkgEqualDeleteOne) Exec(ctx context.Context) error {
	n, err := pedo.ped.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{pkgequal.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pedo *PkgEqualDeleteOne) ExecX(ctx context.Context) {
	if err := pedo.Exec(ctx); err != nil {
		panic(err)
	}
}
