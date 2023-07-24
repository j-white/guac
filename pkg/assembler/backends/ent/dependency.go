// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/gremlin"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/dependency"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packagename"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
)

// Dependency is the model entity for the Dependency schema.
type Dependency struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PackageID holds the value of the "package_id" field.
	PackageID int `json:"package_id,omitempty"`
	// DependentPackageID holds the value of the "dependent_package_id" field.
	DependentPackageID int `json:"dependent_package_id,omitempty"`
	// VersionRange holds the value of the "version_range" field.
	VersionRange string `json:"version_range,omitempty"`
	// DependencyType holds the value of the "dependency_type" field.
	DependencyType dependency.DependencyType `json:"dependency_type,omitempty"`
	// Justification holds the value of the "justification" field.
	Justification string `json:"justification,omitempty"`
	// Origin holds the value of the "origin" field.
	Origin string `json:"origin,omitempty"`
	// Collector holds the value of the "collector" field.
	Collector string `json:"collector,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DependencyQuery when eager-loading is set.
	Edges DependencyEdges `json:"edges"`
}

// DependencyEdges holds the relations/edges for other nodes in the graph.
type DependencyEdges struct {
	// Package holds the value of the package edge.
	Package *PackageVersion `json:"package,omitempty"`
	// DependentPackage holds the value of the dependent_package edge.
	DependentPackage *PackageName `json:"dependent_package,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PackageOrErr returns the Package value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DependencyEdges) PackageOrErr() (*PackageVersion, error) {
	if e.loadedTypes[0] {
		if e.Package == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: packageversion.Label}
		}
		return e.Package, nil
	}
	return nil, &NotLoadedError{edge: "package"}
}

// DependentPackageOrErr returns the DependentPackage value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DependencyEdges) DependentPackageOrErr() (*PackageName, error) {
	if e.loadedTypes[1] {
		if e.DependentPackage == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: packagename.Label}
		}
		return e.DependentPackage, nil
	}
	return nil, &NotLoadedError{edge: "dependent_package"}
}

// FromResponse scans the gremlin response data into Dependency.
func (d *Dependency) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scand struct {
		ID                 int                       `json:"id,omitempty"`
		PackageID          int                       `json:"package_id,omitempty"`
		DependentPackageID int                       `json:"dependent_package_id,omitempty"`
		VersionRange       string                    `json:"version_range,omitempty"`
		DependencyType     dependency.DependencyType `json:"dependency_type,omitempty"`
		Justification      string                    `json:"justification,omitempty"`
		Origin             string                    `json:"origin,omitempty"`
		Collector          string                    `json:"collector,omitempty"`
	}
	if err := vmap.Decode(&scand); err != nil {
		return err
	}
	d.ID = scand.ID
	d.PackageID = scand.PackageID
	d.DependentPackageID = scand.DependentPackageID
	d.VersionRange = scand.VersionRange
	d.DependencyType = scand.DependencyType
	d.Justification = scand.Justification
	d.Origin = scand.Origin
	d.Collector = scand.Collector
	return nil
}

// QueryPackage queries the "package" edge of the Dependency entity.
func (d *Dependency) QueryPackage() *PackageVersionQuery {
	return NewDependencyClient(d.config).QueryPackage(d)
}

// QueryDependentPackage queries the "dependent_package" edge of the Dependency entity.
func (d *Dependency) QueryDependentPackage() *PackageNameQuery {
	return NewDependencyClient(d.config).QueryDependentPackage(d)
}

// Update returns a builder for updating this Dependency.
// Note that you need to call Dependency.Unwrap() before calling this method if this Dependency
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Dependency) Update() *DependencyUpdateOne {
	return NewDependencyClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Dependency entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Dependency) Unwrap() *Dependency {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Dependency is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Dependency) String() string {
	var builder strings.Builder
	builder.WriteString("Dependency(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("package_id=")
	builder.WriteString(fmt.Sprintf("%v", d.PackageID))
	builder.WriteString(", ")
	builder.WriteString("dependent_package_id=")
	builder.WriteString(fmt.Sprintf("%v", d.DependentPackageID))
	builder.WriteString(", ")
	builder.WriteString("version_range=")
	builder.WriteString(d.VersionRange)
	builder.WriteString(", ")
	builder.WriteString("dependency_type=")
	builder.WriteString(fmt.Sprintf("%v", d.DependencyType))
	builder.WriteString(", ")
	builder.WriteString("justification=")
	builder.WriteString(d.Justification)
	builder.WriteString(", ")
	builder.WriteString("origin=")
	builder.WriteString(d.Origin)
	builder.WriteString(", ")
	builder.WriteString("collector=")
	builder.WriteString(d.Collector)
	builder.WriteByte(')')
	return builder.String()
}

// Dependencies is a parsable slice of Dependency.
type Dependencies []*Dependency

// FromResponse scans the gremlin response data into Dependencies.
func (d *Dependencies) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scand []struct {
		ID                 int                       `json:"id,omitempty"`
		PackageID          int                       `json:"package_id,omitempty"`
		DependentPackageID int                       `json:"dependent_package_id,omitempty"`
		VersionRange       string                    `json:"version_range,omitempty"`
		DependencyType     dependency.DependencyType `json:"dependency_type,omitempty"`
		Justification      string                    `json:"justification,omitempty"`
		Origin             string                    `json:"origin,omitempty"`
		Collector          string                    `json:"collector,omitempty"`
	}
	if err := vmap.Decode(&scand); err != nil {
		return err
	}
	for _, v := range scand {
		node := &Dependency{ID: v.ID}
		node.PackageID = v.PackageID
		node.DependentPackageID = v.DependentPackageID
		node.VersionRange = v.VersionRange
		node.DependencyType = v.DependencyType
		node.Justification = v.Justification
		node.Origin = v.Origin
		node.Collector = v.Collector
		*d = append(*d, node)
	}
	return nil
}
