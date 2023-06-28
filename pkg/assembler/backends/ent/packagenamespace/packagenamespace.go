// Code generated by ent, DO NOT EDIT.

package packagenamespace

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the packagenamespace type in the database.
	Label = "package_namespace"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPackageID holds the string denoting the package_id field in the database.
	FieldPackageID = "package_id"
	// FieldNamespace holds the string denoting the namespace field in the database.
	FieldNamespace = "namespace"
	// EdgePackage holds the string denoting the package edge name in mutations.
	EdgePackage = "package"
	// EdgeNames holds the string denoting the names edge name in mutations.
	EdgeNames = "names"
	// Table holds the table name of the packagenamespace in the database.
	Table = "package_namespaces"
	// PackageTable is the table that holds the package relation/edge.
	PackageTable = "package_namespaces"
	// PackageInverseTable is the table name for the PackageType entity.
	// It exists in this package in order to avoid circular dependency with the "packagetype" package.
	PackageInverseTable = "package_types"
	// PackageColumn is the table column denoting the package relation/edge.
	PackageColumn = "package_id"
	// NamesTable is the table that holds the names relation/edge.
	NamesTable = "package_names"
	// NamesInverseTable is the table name for the PackageName entity.
	// It exists in this package in order to avoid circular dependency with the "packagename" package.
	NamesInverseTable = "package_names"
	// NamesColumn is the table column denoting the names relation/edge.
	NamesColumn = "namespace_id"
)

// Columns holds all SQL columns for packagenamespace fields.
var Columns = []string{
	FieldID,
	FieldPackageID,
	FieldNamespace,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the PackageNamespace queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPackageID orders the results by the package_id field.
func ByPackageID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPackageID, opts...).ToFunc()
}

// ByNamespace orders the results by the namespace field.
func ByNamespace(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNamespace, opts...).ToFunc()
}

// ByPackageField orders the results by package field.
func ByPackageField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPackageStep(), sql.OrderByField(field, opts...))
	}
}

// ByNamesCount orders the results by names count.
func ByNamesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNamesStep(), opts...)
	}
}

// ByNames orders the results by names terms.
func ByNames(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNamesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPackageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PackageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PackageTable, PackageColumn),
	)
}
func newNamesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NamesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NamesTable, NamesColumn),
	)
}
