// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldDisplayName holds the string denoting the displayname field in the database.
	FieldDisplayName = "display_name"
	// FieldDob holds the string denoting the dob field in the database.
	FieldDob = "dob"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldCreatedDate holds the string denoting the createddate field in the database.
	FieldCreatedDate = "created_date"
	// FieldAvatarPath holds the string denoting the avatarpath field in the database.
	FieldAvatarPath = "avatar_path"
	// EdgeCreatedBy holds the string denoting the createdby edge name in mutations.
	EdgeCreatedBy = "createdBy"
	// EdgeRole holds the string denoting the role edge name in mutations.
	EdgeRole = "role"
	// Table holds the table name of the user in the database.
	Table = "users"
	// CreatedByTable is the table that holds the createdBy relation/edge.
	CreatedByTable = "books"
	// CreatedByInverseTable is the table name for the Book entity.
	// It exists in this package in order to avoid circular dependency with the "book" package.
	CreatedByInverseTable = "books"
	// CreatedByColumn is the table column denoting the createdBy relation/edge.
	CreatedByColumn = "created_by"
	// RoleTable is the table that holds the role relation/edge. The primary key declared below.
	RoleTable = "role_user"
	// RoleInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RoleInverseTable = "roles"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldPassword,
	FieldDisplayName,
	FieldDob,
	FieldEmail,
	FieldCreatedDate,
	FieldAvatarPath,
}

var (
	// RolePrimaryKey and RoleColumn2 are the table columns denoting the
	// primary key for the role relation (M2M).
	RolePrimaryKey = []string{"role_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByDisplayName orders the results by the displayName field.
func ByDisplayName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisplayName, opts...).ToFunc()
}

// ByDob orders the results by the dob field.
func ByDob(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDob, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByCreatedDate orders the results by the createdDate field.
func ByCreatedDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedDate, opts...).ToFunc()
}

// ByAvatarPath orders the results by the avatarPath field.
func ByAvatarPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatarPath, opts...).ToFunc()
}

// ByCreatedByCount orders the results by createdBy count.
func ByCreatedByCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCreatedByStep(), opts...)
	}
}

// ByCreatedBy orders the results by createdBy terms.
func ByCreatedBy(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatedByStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRoleCount orders the results by role count.
func ByRoleCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRoleStep(), opts...)
	}
}

// ByRole orders the results by role terms.
func ByRole(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoleStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCreatedByStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatedByInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CreatedByTable, CreatedByColumn),
	)
}
func newRoleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, RoleTable, RolePrimaryKey...),
	)
}
