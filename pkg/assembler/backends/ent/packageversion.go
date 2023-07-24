// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/gremlin"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packagename"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// PackageVersion is the model entity for the PackageVersion schema.
type PackageVersion struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// NameID holds the value of the "name_id" field.
	NameID int `json:"name_id,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// Subpath holds the value of the "subpath" field.
	Subpath string `json:"subpath,omitempty"`
	// Qualifiers holds the value of the "qualifiers" field.
	Qualifiers []model.PackageQualifier `json:"qualifiers,omitempty"`
	// A SHA1 of the qualifiers, subpath, version fields after sorting keys, used to ensure uniqueness of version records.
	Hash string `json:"hash,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PackageVersionQuery when eager-loading is set.
	Edges PackageVersionEdges `json:"edges"`
}

// PackageVersionEdges holds the relations/edges for other nodes in the graph.
type PackageVersionEdges struct {
	// Name holds the value of the name edge.
	Name *PackageName `json:"name,omitempty"`
	// Occurrences holds the value of the occurrences edge.
	Occurrences []*Occurrence `json:"occurrences,omitempty"`
	// Sbom holds the value of the sbom edge.
	Sbom []*BillOfMaterials `json:"sbom,omitempty"`
	// EqualPackages holds the value of the equal_packages edge.
	EqualPackages []*PkgEqual `json:"equal_packages,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// NameOrErr returns the Name value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PackageVersionEdges) NameOrErr() (*PackageName, error) {
	if e.loadedTypes[0] {
		if e.Name == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: packagename.Label}
		}
		return e.Name, nil
	}
	return nil, &NotLoadedError{edge: "name"}
}

// OccurrencesOrErr returns the Occurrences value or an error if the edge
// was not loaded in eager-loading.
func (e PackageVersionEdges) OccurrencesOrErr() ([]*Occurrence, error) {
	if e.loadedTypes[1] {
		return e.Occurrences, nil
	}
	return nil, &NotLoadedError{edge: "occurrences"}
}

// SbomOrErr returns the Sbom value or an error if the edge
// was not loaded in eager-loading.
func (e PackageVersionEdges) SbomOrErr() ([]*BillOfMaterials, error) {
	if e.loadedTypes[2] {
		return e.Sbom, nil
	}
	return nil, &NotLoadedError{edge: "sbom"}
}

// EqualPackagesOrErr returns the EqualPackages value or an error if the edge
// was not loaded in eager-loading.
func (e PackageVersionEdges) EqualPackagesOrErr() ([]*PkgEqual, error) {
	if e.loadedTypes[3] {
		return e.EqualPackages, nil
	}
	return nil, &NotLoadedError{edge: "equal_packages"}
}

// FromResponse scans the gremlin response data into PackageVersion.
func (pv *PackageVersion) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanpv struct {
		ID         int                      `json:"id,omitempty"`
		NameID     int                      `json:"name_id,omitempty"`
		Version    string                   `json:"version,omitempty"`
		Subpath    string                   `json:"subpath,omitempty"`
		Qualifiers []model.PackageQualifier `json:"qualifiers,omitempty"`
		Hash       string                   `json:"hash,omitempty"`
	}
	if err := vmap.Decode(&scanpv); err != nil {
		return err
	}
	pv.ID = scanpv.ID
	pv.NameID = scanpv.NameID
	pv.Version = scanpv.Version
	pv.Subpath = scanpv.Subpath
	pv.Qualifiers = scanpv.Qualifiers
	pv.Hash = scanpv.Hash
	return nil
}

// QueryName queries the "name" edge of the PackageVersion entity.
func (pv *PackageVersion) QueryName() *PackageNameQuery {
	return NewPackageVersionClient(pv.config).QueryName(pv)
}

// QueryOccurrences queries the "occurrences" edge of the PackageVersion entity.
func (pv *PackageVersion) QueryOccurrences() *OccurrenceQuery {
	return NewPackageVersionClient(pv.config).QueryOccurrences(pv)
}

// QuerySbom queries the "sbom" edge of the PackageVersion entity.
func (pv *PackageVersion) QuerySbom() *BillOfMaterialsQuery {
	return NewPackageVersionClient(pv.config).QuerySbom(pv)
}

// QueryEqualPackages queries the "equal_packages" edge of the PackageVersion entity.
func (pv *PackageVersion) QueryEqualPackages() *PkgEqualQuery {
	return NewPackageVersionClient(pv.config).QueryEqualPackages(pv)
}

// Update returns a builder for updating this PackageVersion.
// Note that you need to call PackageVersion.Unwrap() before calling this method if this PackageVersion
// was returned from a transaction, and the transaction was committed or rolled back.
func (pv *PackageVersion) Update() *PackageVersionUpdateOne {
	return NewPackageVersionClient(pv.config).UpdateOne(pv)
}

// Unwrap unwraps the PackageVersion entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pv *PackageVersion) Unwrap() *PackageVersion {
	_tx, ok := pv.config.driver.(*txDriver)
	if !ok {
		panic("ent: PackageVersion is not a transactional entity")
	}
	pv.config.driver = _tx.drv
	return pv
}

// String implements the fmt.Stringer.
func (pv *PackageVersion) String() string {
	var builder strings.Builder
	builder.WriteString("PackageVersion(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pv.ID))
	builder.WriteString("name_id=")
	builder.WriteString(fmt.Sprintf("%v", pv.NameID))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(pv.Version)
	builder.WriteString(", ")
	builder.WriteString("subpath=")
	builder.WriteString(pv.Subpath)
	builder.WriteString(", ")
	builder.WriteString("qualifiers=")
	builder.WriteString(fmt.Sprintf("%v", pv.Qualifiers))
	builder.WriteString(", ")
	builder.WriteString("hash=")
	builder.WriteString(pv.Hash)
	builder.WriteByte(')')
	return builder.String()
}

// PackageVersions is a parsable slice of PackageVersion.
type PackageVersions []*PackageVersion

// FromResponse scans the gremlin response data into PackageVersions.
func (pv *PackageVersions) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanpv []struct {
		ID         int                      `json:"id,omitempty"`
		NameID     int                      `json:"name_id,omitempty"`
		Version    string                   `json:"version,omitempty"`
		Subpath    string                   `json:"subpath,omitempty"`
		Qualifiers []model.PackageQualifier `json:"qualifiers,omitempty"`
		Hash       string                   `json:"hash,omitempty"`
	}
	if err := vmap.Decode(&scanpv); err != nil {
		return err
	}
	for _, v := range scanpv {
		node := &PackageVersion{ID: v.ID}
		node.NameID = v.NameID
		node.Version = v.Version
		node.Subpath = v.Subpath
		node.Qualifiers = v.Qualifiers
		node.Hash = v.Hash
		*pv = append(*pv, node)
	}
	return nil
}
