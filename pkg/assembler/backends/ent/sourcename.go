// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/sourcename"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/sourcenamespace"
)

// SourceName is the model entity for the SourceName schema.
type SourceName struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Commit holds the value of the "commit" field.
	Commit string `json:"commit,omitempty"`
	// Tag holds the value of the "tag" field.
	Tag string `json:"tag,omitempty"`
	// NamespaceID holds the value of the "namespace_id" field.
	NamespaceID int `json:"namespace_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SourceNameQuery when eager-loading is set.
	Edges        SourceNameEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SourceNameEdges holds the relations/edges for other nodes in the graph.
type SourceNameEdges struct {
	// Namespace holds the value of the namespace edge.
	Namespace *SourceNamespace `json:"namespace,omitempty"`
	// Occurrences holds the value of the occurrences edge.
	Occurrences []*IsOccurrence `json:"occurrences,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// NamespaceOrErr returns the Namespace value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SourceNameEdges) NamespaceOrErr() (*SourceNamespace, error) {
	if e.loadedTypes[0] {
		if e.Namespace == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: sourcenamespace.Label}
		}
		return e.Namespace, nil
	}
	return nil, &NotLoadedError{edge: "namespace"}
}

// OccurrencesOrErr returns the Occurrences value or an error if the edge
// was not loaded in eager-loading.
func (e SourceNameEdges) OccurrencesOrErr() ([]*IsOccurrence, error) {
	if e.loadedTypes[1] {
		return e.Occurrences, nil
	}
	return nil, &NotLoadedError{edge: "occurrences"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SourceName) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sourcename.FieldID, sourcename.FieldNamespaceID:
			values[i] = new(sql.NullInt64)
		case sourcename.FieldName, sourcename.FieldCommit, sourcename.FieldTag:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SourceName fields.
func (sn *SourceName) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sourcename.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sn.ID = int(value.Int64)
		case sourcename.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sn.Name = value.String
			}
		case sourcename.FieldCommit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field commit", values[i])
			} else if value.Valid {
				sn.Commit = value.String
			}
		case sourcename.FieldTag:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tag", values[i])
			} else if value.Valid {
				sn.Tag = value.String
			}
		case sourcename.FieldNamespaceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field namespace_id", values[i])
			} else if value.Valid {
				sn.NamespaceID = int(value.Int64)
			}
		default:
			sn.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SourceName.
// This includes values selected through modifiers, order, etc.
func (sn *SourceName) Value(name string) (ent.Value, error) {
	return sn.selectValues.Get(name)
}

// QueryNamespace queries the "namespace" edge of the SourceName entity.
func (sn *SourceName) QueryNamespace() *SourceNamespaceQuery {
	return NewSourceNameClient(sn.config).QueryNamespace(sn)
}

// QueryOccurrences queries the "occurrences" edge of the SourceName entity.
func (sn *SourceName) QueryOccurrences() *IsOccurrenceQuery {
	return NewSourceNameClient(sn.config).QueryOccurrences(sn)
}

// Update returns a builder for updating this SourceName.
// Note that you need to call SourceName.Unwrap() before calling this method if this SourceName
// was returned from a transaction, and the transaction was committed or rolled back.
func (sn *SourceName) Update() *SourceNameUpdateOne {
	return NewSourceNameClient(sn.config).UpdateOne(sn)
}

// Unwrap unwraps the SourceName entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sn *SourceName) Unwrap() *SourceName {
	_tx, ok := sn.config.driver.(*txDriver)
	if !ok {
		panic("ent: SourceName is not a transactional entity")
	}
	sn.config.driver = _tx.drv
	return sn
}

// String implements the fmt.Stringer.
func (sn *SourceName) String() string {
	var builder strings.Builder
	builder.WriteString("SourceName(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sn.ID))
	builder.WriteString("name=")
	builder.WriteString(sn.Name)
	builder.WriteString(", ")
	builder.WriteString("commit=")
	builder.WriteString(sn.Commit)
	builder.WriteString(", ")
	builder.WriteString("tag=")
	builder.WriteString(sn.Tag)
	builder.WriteString(", ")
	builder.WriteString("namespace_id=")
	builder.WriteString(fmt.Sprintf("%v", sn.NamespaceID))
	builder.WriteByte(')')
	return builder.String()
}

// SourceNames is a parsable slice of SourceName.
type SourceNames []*SourceName
