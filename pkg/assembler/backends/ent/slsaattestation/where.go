// Code generated by ent, DO NOT EDIT.

package slsaattestation

import (
	"time"

	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.HasID(id)
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.HasID(p.EQ(id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.HasID(p.NEQ(id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Within(v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Without(v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.HasID(p.GT(id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.HasID(p.GTE(id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.HasID(p.LT(id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.HasID(p.LTE(id))
	})
}

// BuildType applies equality check predicate on the "build_type" field. It's identical to BuildTypeEQ.
func BuildType(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuildType, p.EQ(v))
	})
}

// BuiltByID applies equality check predicate on the "built_by_id" field. It's identical to BuiltByIDEQ.
func BuiltByID(v int) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuiltByID, p.EQ(v))
	})
}

// SubjectID applies equality check predicate on the "subject_id" field. It's identical to SubjectIDEQ.
func SubjectID(v int) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldSubjectID, p.EQ(v))
	})
}

// SlsaVersion applies equality check predicate on the "slsa_version" field. It's identical to SlsaVersionEQ.
func SlsaVersion(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldSlsaVersion, p.EQ(v))
	})
}

// StartedOn applies equality check predicate on the "started_on" field. It's identical to StartedOnEQ.
func StartedOn(v time.Time) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldStartedOn, p.EQ(v))
	})
}

// FinishedOn applies equality check predicate on the "finished_on" field. It's identical to FinishedOnEQ.
func FinishedOn(v time.Time) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldFinishedOn, p.EQ(v))
	})
}

// Origin applies equality check predicate on the "origin" field. It's identical to OriginEQ.
func Origin(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldOrigin, p.EQ(v))
	})
}

// Collector applies equality check predicate on the "collector" field. It's identical to CollectorEQ.
func Collector(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldCollector, p.EQ(v))
	})
}

// BuiltFromHash applies equality check predicate on the "built_from_hash" field. It's identical to BuiltFromHashEQ.
func BuiltFromHash(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuiltFromHash, p.EQ(v))
	})
}

// BuildTypeEQ applies the EQ predicate on the "build_type" field.
func BuildTypeEQ(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuildType, p.EQ(v))
	})
}

// BuildTypeNEQ applies the NEQ predicate on the "build_type" field.
func BuildTypeNEQ(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuildType, p.NEQ(v))
	})
}

// BuildTypeIn applies the In predicate on the "build_type" field.
func BuildTypeIn(vs ...string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuildType, p.Within(vs...))
	})
}

// BuildTypeNotIn applies the NotIn predicate on the "build_type" field.
func BuildTypeNotIn(vs ...string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuildType, p.Without(vs...))
	})
}

// BuildTypeGT applies the GT predicate on the "build_type" field.
func BuildTypeGT(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuildType, p.GT(v))
	})
}

// BuildTypeGTE applies the GTE predicate on the "build_type" field.
func BuildTypeGTE(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuildType, p.GTE(v))
	})
}

// BuildTypeLT applies the LT predicate on the "build_type" field.
func BuildTypeLT(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuildType, p.LT(v))
	})
}

// BuildTypeLTE applies the LTE predicate on the "build_type" field.
func BuildTypeLTE(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuildType, p.LTE(v))
	})
}

// BuildTypeContains applies the Contains predicate on the "build_type" field.
func BuildTypeContains(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuildType, p.Containing(v))
	})
}

// BuildTypeHasPrefix applies the HasPrefix predicate on the "build_type" field.
func BuildTypeHasPrefix(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuildType, p.StartingWith(v))
	})
}

// BuildTypeHasSuffix applies the HasSuffix predicate on the "build_type" field.
func BuildTypeHasSuffix(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuildType, p.EndingWith(v))
	})
}

// BuiltByIDEQ applies the EQ predicate on the "built_by_id" field.
func BuiltByIDEQ(v int) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuiltByID, p.EQ(v))
	})
}

// BuiltByIDNEQ applies the NEQ predicate on the "built_by_id" field.
func BuiltByIDNEQ(v int) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuiltByID, p.NEQ(v))
	})
}

// BuiltByIDIn applies the In predicate on the "built_by_id" field.
func BuiltByIDIn(vs ...int) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuiltByID, p.Within(vs...))
	})
}

// BuiltByIDNotIn applies the NotIn predicate on the "built_by_id" field.
func BuiltByIDNotIn(vs ...int) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuiltByID, p.Without(vs...))
	})
}

// SubjectIDEQ applies the EQ predicate on the "subject_id" field.
func SubjectIDEQ(v int) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldSubjectID, p.EQ(v))
	})
}

// SubjectIDNEQ applies the NEQ predicate on the "subject_id" field.
func SubjectIDNEQ(v int) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldSubjectID, p.NEQ(v))
	})
}

// SubjectIDIn applies the In predicate on the "subject_id" field.
func SubjectIDIn(vs ...int) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldSubjectID, p.Within(vs...))
	})
}

// SubjectIDNotIn applies the NotIn predicate on the "subject_id" field.
func SubjectIDNotIn(vs ...int) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldSubjectID, p.Without(vs...))
	})
}

// SlsaPredicateIsNil applies the IsNil predicate on the "slsa_predicate" field.
func SlsaPredicateIsNil() predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldSlsaPredicate)
	})
}

// SlsaPredicateNotNil applies the NotNil predicate on the "slsa_predicate" field.
func SlsaPredicateNotNil() predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldSlsaPredicate)
	})
}

// SlsaVersionEQ applies the EQ predicate on the "slsa_version" field.
func SlsaVersionEQ(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldSlsaVersion, p.EQ(v))
	})
}

// SlsaVersionNEQ applies the NEQ predicate on the "slsa_version" field.
func SlsaVersionNEQ(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldSlsaVersion, p.NEQ(v))
	})
}

// SlsaVersionIn applies the In predicate on the "slsa_version" field.
func SlsaVersionIn(vs ...string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldSlsaVersion, p.Within(vs...))
	})
}

// SlsaVersionNotIn applies the NotIn predicate on the "slsa_version" field.
func SlsaVersionNotIn(vs ...string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldSlsaVersion, p.Without(vs...))
	})
}

// SlsaVersionGT applies the GT predicate on the "slsa_version" field.
func SlsaVersionGT(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldSlsaVersion, p.GT(v))
	})
}

// SlsaVersionGTE applies the GTE predicate on the "slsa_version" field.
func SlsaVersionGTE(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldSlsaVersion, p.GTE(v))
	})
}

// SlsaVersionLT applies the LT predicate on the "slsa_version" field.
func SlsaVersionLT(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldSlsaVersion, p.LT(v))
	})
}

// SlsaVersionLTE applies the LTE predicate on the "slsa_version" field.
func SlsaVersionLTE(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldSlsaVersion, p.LTE(v))
	})
}

// SlsaVersionContains applies the Contains predicate on the "slsa_version" field.
func SlsaVersionContains(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldSlsaVersion, p.Containing(v))
	})
}

// SlsaVersionHasPrefix applies the HasPrefix predicate on the "slsa_version" field.
func SlsaVersionHasPrefix(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldSlsaVersion, p.StartingWith(v))
	})
}

// SlsaVersionHasSuffix applies the HasSuffix predicate on the "slsa_version" field.
func SlsaVersionHasSuffix(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldSlsaVersion, p.EndingWith(v))
	})
}

// StartedOnEQ applies the EQ predicate on the "started_on" field.
func StartedOnEQ(v time.Time) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldStartedOn, p.EQ(v))
	})
}

// StartedOnNEQ applies the NEQ predicate on the "started_on" field.
func StartedOnNEQ(v time.Time) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldStartedOn, p.NEQ(v))
	})
}

// StartedOnIn applies the In predicate on the "started_on" field.
func StartedOnIn(vs ...time.Time) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldStartedOn, p.Within(vs...))
	})
}

// StartedOnNotIn applies the NotIn predicate on the "started_on" field.
func StartedOnNotIn(vs ...time.Time) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldStartedOn, p.Without(vs...))
	})
}

// StartedOnGT applies the GT predicate on the "started_on" field.
func StartedOnGT(v time.Time) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldStartedOn, p.GT(v))
	})
}

// StartedOnGTE applies the GTE predicate on the "started_on" field.
func StartedOnGTE(v time.Time) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldStartedOn, p.GTE(v))
	})
}

// StartedOnLT applies the LT predicate on the "started_on" field.
func StartedOnLT(v time.Time) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldStartedOn, p.LT(v))
	})
}

// StartedOnLTE applies the LTE predicate on the "started_on" field.
func StartedOnLTE(v time.Time) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldStartedOn, p.LTE(v))
	})
}

// StartedOnIsNil applies the IsNil predicate on the "started_on" field.
func StartedOnIsNil() predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldStartedOn)
	})
}

// StartedOnNotNil applies the NotNil predicate on the "started_on" field.
func StartedOnNotNil() predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldStartedOn)
	})
}

// FinishedOnEQ applies the EQ predicate on the "finished_on" field.
func FinishedOnEQ(v time.Time) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldFinishedOn, p.EQ(v))
	})
}

// FinishedOnNEQ applies the NEQ predicate on the "finished_on" field.
func FinishedOnNEQ(v time.Time) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldFinishedOn, p.NEQ(v))
	})
}

// FinishedOnIn applies the In predicate on the "finished_on" field.
func FinishedOnIn(vs ...time.Time) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldFinishedOn, p.Within(vs...))
	})
}

// FinishedOnNotIn applies the NotIn predicate on the "finished_on" field.
func FinishedOnNotIn(vs ...time.Time) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldFinishedOn, p.Without(vs...))
	})
}

// FinishedOnGT applies the GT predicate on the "finished_on" field.
func FinishedOnGT(v time.Time) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldFinishedOn, p.GT(v))
	})
}

// FinishedOnGTE applies the GTE predicate on the "finished_on" field.
func FinishedOnGTE(v time.Time) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldFinishedOn, p.GTE(v))
	})
}

// FinishedOnLT applies the LT predicate on the "finished_on" field.
func FinishedOnLT(v time.Time) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldFinishedOn, p.LT(v))
	})
}

// FinishedOnLTE applies the LTE predicate on the "finished_on" field.
func FinishedOnLTE(v time.Time) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldFinishedOn, p.LTE(v))
	})
}

// FinishedOnIsNil applies the IsNil predicate on the "finished_on" field.
func FinishedOnIsNil() predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldFinishedOn)
	})
}

// FinishedOnNotNil applies the NotNil predicate on the "finished_on" field.
func FinishedOnNotNil() predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldFinishedOn)
	})
}

// OriginEQ applies the EQ predicate on the "origin" field.
func OriginEQ(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldOrigin, p.EQ(v))
	})
}

// OriginNEQ applies the NEQ predicate on the "origin" field.
func OriginNEQ(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldOrigin, p.NEQ(v))
	})
}

// OriginIn applies the In predicate on the "origin" field.
func OriginIn(vs ...string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldOrigin, p.Within(vs...))
	})
}

// OriginNotIn applies the NotIn predicate on the "origin" field.
func OriginNotIn(vs ...string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldOrigin, p.Without(vs...))
	})
}

// OriginGT applies the GT predicate on the "origin" field.
func OriginGT(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldOrigin, p.GT(v))
	})
}

// OriginGTE applies the GTE predicate on the "origin" field.
func OriginGTE(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldOrigin, p.GTE(v))
	})
}

// OriginLT applies the LT predicate on the "origin" field.
func OriginLT(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldOrigin, p.LT(v))
	})
}

// OriginLTE applies the LTE predicate on the "origin" field.
func OriginLTE(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldOrigin, p.LTE(v))
	})
}

// OriginContains applies the Contains predicate on the "origin" field.
func OriginContains(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldOrigin, p.Containing(v))
	})
}

// OriginHasPrefix applies the HasPrefix predicate on the "origin" field.
func OriginHasPrefix(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldOrigin, p.StartingWith(v))
	})
}

// OriginHasSuffix applies the HasSuffix predicate on the "origin" field.
func OriginHasSuffix(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldOrigin, p.EndingWith(v))
	})
}

// CollectorEQ applies the EQ predicate on the "collector" field.
func CollectorEQ(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldCollector, p.EQ(v))
	})
}

// CollectorNEQ applies the NEQ predicate on the "collector" field.
func CollectorNEQ(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldCollector, p.NEQ(v))
	})
}

// CollectorIn applies the In predicate on the "collector" field.
func CollectorIn(vs ...string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldCollector, p.Within(vs...))
	})
}

// CollectorNotIn applies the NotIn predicate on the "collector" field.
func CollectorNotIn(vs ...string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldCollector, p.Without(vs...))
	})
}

// CollectorGT applies the GT predicate on the "collector" field.
func CollectorGT(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldCollector, p.GT(v))
	})
}

// CollectorGTE applies the GTE predicate on the "collector" field.
func CollectorGTE(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldCollector, p.GTE(v))
	})
}

// CollectorLT applies the LT predicate on the "collector" field.
func CollectorLT(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldCollector, p.LT(v))
	})
}

// CollectorLTE applies the LTE predicate on the "collector" field.
func CollectorLTE(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldCollector, p.LTE(v))
	})
}

// CollectorContains applies the Contains predicate on the "collector" field.
func CollectorContains(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldCollector, p.Containing(v))
	})
}

// CollectorHasPrefix applies the HasPrefix predicate on the "collector" field.
func CollectorHasPrefix(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldCollector, p.StartingWith(v))
	})
}

// CollectorHasSuffix applies the HasSuffix predicate on the "collector" field.
func CollectorHasSuffix(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldCollector, p.EndingWith(v))
	})
}

// BuiltFromHashEQ applies the EQ predicate on the "built_from_hash" field.
func BuiltFromHashEQ(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuiltFromHash, p.EQ(v))
	})
}

// BuiltFromHashNEQ applies the NEQ predicate on the "built_from_hash" field.
func BuiltFromHashNEQ(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuiltFromHash, p.NEQ(v))
	})
}

// BuiltFromHashIn applies the In predicate on the "built_from_hash" field.
func BuiltFromHashIn(vs ...string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuiltFromHash, p.Within(vs...))
	})
}

// BuiltFromHashNotIn applies the NotIn predicate on the "built_from_hash" field.
func BuiltFromHashNotIn(vs ...string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuiltFromHash, p.Without(vs...))
	})
}

// BuiltFromHashGT applies the GT predicate on the "built_from_hash" field.
func BuiltFromHashGT(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuiltFromHash, p.GT(v))
	})
}

// BuiltFromHashGTE applies the GTE predicate on the "built_from_hash" field.
func BuiltFromHashGTE(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuiltFromHash, p.GTE(v))
	})
}

// BuiltFromHashLT applies the LT predicate on the "built_from_hash" field.
func BuiltFromHashLT(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuiltFromHash, p.LT(v))
	})
}

// BuiltFromHashLTE applies the LTE predicate on the "built_from_hash" field.
func BuiltFromHashLTE(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuiltFromHash, p.LTE(v))
	})
}

// BuiltFromHashContains applies the Contains predicate on the "built_from_hash" field.
func BuiltFromHashContains(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuiltFromHash, p.Containing(v))
	})
}

// BuiltFromHashHasPrefix applies the HasPrefix predicate on the "built_from_hash" field.
func BuiltFromHashHasPrefix(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuiltFromHash, p.StartingWith(v))
	})
}

// BuiltFromHashHasSuffix applies the HasSuffix predicate on the "built_from_hash" field.
func BuiltFromHashHasSuffix(v string) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.Has(Label, FieldBuiltFromHash, p.EndingWith(v))
	})
}

// HasBuiltFrom applies the HasEdge predicate on the "built_from" edge.
func HasBuiltFrom() predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.OutE(BuiltFromLabel).OutV()
	})
}

// HasBuiltFromWith applies the HasEdge predicate on the "built_from" edge with a given conditions (other predicates).
func HasBuiltFromWith(preds ...predicate.Artifact) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		tr := __.InV()
		for _, p := range preds {
			p(tr)
		}
		t.OutE(BuiltFromLabel).Where(tr).OutV()
	})
}

// HasBuiltBy applies the HasEdge predicate on the "built_by" edge.
func HasBuiltBy() predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.OutE(BuiltByLabel).OutV()
	})
}

// HasBuiltByWith applies the HasEdge predicate on the "built_by" edge with a given conditions (other predicates).
func HasBuiltByWith(preds ...predicate.Builder) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		tr := __.InV()
		for _, p := range preds {
			p(tr)
		}
		t.OutE(BuiltByLabel).Where(tr).OutV()
	})
}

// HasSubject applies the HasEdge predicate on the "subject" edge.
func HasSubject() predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		t.OutE(SubjectLabel).OutV()
	})
}

// HasSubjectWith applies the HasEdge predicate on the "subject" edge with a given conditions (other predicates).
func HasSubjectWith(preds ...predicate.Artifact) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(t *dsl.Traversal) {
		tr := __.InV()
		for _, p := range preds {
			p(tr)
		}
		t.OutE(SubjectLabel).Where(tr).OutV()
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SLSAAttestation) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(tr *dsl.Traversal) {
		trs := make([]any, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.And(trs...))
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SLSAAttestation) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(tr *dsl.Traversal) {
		trs := make([]any, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.Or(trs...))
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SLSAAttestation) predicate.SLSAAttestation {
	return predicate.SLSAAttestation(func(tr *dsl.Traversal) {
		t := __.New()
		p(t)
		tr.Where(__.Not(t))
	})
}
