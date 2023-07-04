// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/securityadvisory"
)

// SecurityAdvisory is the model entity for the SecurityAdvisory schema.
type SecurityAdvisory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// GhsaID holds the value of the "ghsa_id" field.
	GhsaID *string `json:"ghsa_id,omitempty"`
	// CveID holds the value of the "cve_id" field.
	CveID *string `json:"cve_id,omitempty"`
	// CveYear holds the value of the "cve_year" field.
	CveYear      *int `json:"cve_year,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SecurityAdvisory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case securityadvisory.FieldID, securityadvisory.FieldCveYear:
			values[i] = new(sql.NullInt64)
		case securityadvisory.FieldGhsaID, securityadvisory.FieldCveID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SecurityAdvisory fields.
func (sa *SecurityAdvisory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case securityadvisory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sa.ID = int(value.Int64)
		case securityadvisory.FieldGhsaID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ghsa_id", values[i])
			} else if value.Valid {
				sa.GhsaID = new(string)
				*sa.GhsaID = value.String
			}
		case securityadvisory.FieldCveID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cve_id", values[i])
			} else if value.Valid {
				sa.CveID = new(string)
				*sa.CveID = value.String
			}
		case securityadvisory.FieldCveYear:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cve_year", values[i])
			} else if value.Valid {
				sa.CveYear = new(int)
				*sa.CveYear = int(value.Int64)
			}
		default:
			sa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SecurityAdvisory.
// This includes values selected through modifiers, order, etc.
func (sa *SecurityAdvisory) Value(name string) (ent.Value, error) {
	return sa.selectValues.Get(name)
}

// Update returns a builder for updating this SecurityAdvisory.
// Note that you need to call SecurityAdvisory.Unwrap() before calling this method if this SecurityAdvisory
// was returned from a transaction, and the transaction was committed or rolled back.
func (sa *SecurityAdvisory) Update() *SecurityAdvisoryUpdateOne {
	return NewSecurityAdvisoryClient(sa.config).UpdateOne(sa)
}

// Unwrap unwraps the SecurityAdvisory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sa *SecurityAdvisory) Unwrap() *SecurityAdvisory {
	_tx, ok := sa.config.driver.(*txDriver)
	if !ok {
		panic("ent: SecurityAdvisory is not a transactional entity")
	}
	sa.config.driver = _tx.drv
	return sa
}

// String implements the fmt.Stringer.
func (sa *SecurityAdvisory) String() string {
	var builder strings.Builder
	builder.WriteString("SecurityAdvisory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sa.ID))
	if v := sa.GhsaID; v != nil {
		builder.WriteString("ghsa_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sa.CveID; v != nil {
		builder.WriteString("cve_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sa.CveYear; v != nil {
		builder.WriteString("cve_year=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// SecurityAdvisories is a parsable slice of SecurityAdvisory.
type SecurityAdvisories []*SecurityAdvisory
