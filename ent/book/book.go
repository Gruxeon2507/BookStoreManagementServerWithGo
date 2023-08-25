// Code generated by ent, DO NOT EDIT.

package book

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the book type in the database.
	Label = "book"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldAuthorName holds the string denoting the authorname field in the database.
	FieldAuthorName = "author_name"
	// FieldIsApproved holds the string denoting the isapproved field in the database.
	FieldIsApproved = "is_approved"
	// FieldCoverPath holds the string denoting the coverpath field in the database.
	FieldCoverPath = "cover_path"
	// FieldPdfPath holds the string denoting the pdfpath field in the database.
	FieldPdfPath = "pdf_path"
	// FieldCreatedBy holds the string denoting the createdby field in the database.
	FieldCreatedBy = "created_by"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldNoSale holds the string denoting the nosale field in the database.
	FieldNoSale = "no_sale"
	// FieldNoView holds the string denoting the noview field in the database.
	FieldNoView = "no_view"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "category"
	// Table holds the table name of the book in the database.
	Table = "books"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "books"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "created_by"
	// CategoryTable is the table that holds the category relation/edge. The primary key declared below.
	CategoryTable = "book_category"
	// CategoryInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoryInverseTable = "categories"
)

// Columns holds all SQL columns for book fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldAuthorName,
	FieldIsApproved,
	FieldCoverPath,
	FieldPdfPath,
	FieldCreatedBy,
	FieldPrice,
	FieldNoSale,
	FieldNoView,
	FieldDescription,
}

var (
	// CategoryPrimaryKey and CategoryColumn2 are the table columns denoting the
	// primary key for the category relation (M2M).
	CategoryPrimaryKey = []string{"book_id", "category_id"}
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

var (
	// DefaultIsApproved holds the default value on creation for the "isApproved" field.
	DefaultIsApproved bool
	// DefaultNoSale holds the default value on creation for the "noSale" field.
	DefaultNoSale int
	// DefaultNoView holds the default value on creation for the "noView" field.
	DefaultNoView int
)

// OrderOption defines the ordering options for the Book queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByAuthorName orders the results by the authorName field.
func ByAuthorName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthorName, opts...).ToFunc()
}

// ByIsApproved orders the results by the isApproved field.
func ByIsApproved(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsApproved, opts...).ToFunc()
}

// ByCoverPath orders the results by the coverPath field.
func ByCoverPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoverPath, opts...).ToFunc()
}

// ByPdfPath orders the results by the pdfPath field.
func ByPdfPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPdfPath, opts...).ToFunc()
}

// ByCreatedBy orders the results by the createdBy field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByNoSale orders the results by the noSale field.
func ByNoSale(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNoSale, opts...).ToFunc()
}

// ByNoView orders the results by the noView field.
func ByNoView(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNoView, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByCategoryCount orders the results by category count.
func ByCategoryCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCategoryStep(), opts...)
	}
}

// ByCategory orders the results by category terms.
func ByCategory(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCategoryStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newCategoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, CategoryTable, CategoryPrimaryKey...),
	)
}