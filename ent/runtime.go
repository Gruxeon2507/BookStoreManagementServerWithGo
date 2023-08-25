// Code generated by ent, DO NOT EDIT.

package ent

import (
	"bookstore/ent/book"
	"bookstore/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	bookFields := schema.Book{}.Fields()
	_ = bookFields
	// bookDescIsApproved is the schema descriptor for isApproved field.
	bookDescIsApproved := bookFields[2].Descriptor()
	// book.DefaultIsApproved holds the default value on creation for the isApproved field.
	book.DefaultIsApproved = bookDescIsApproved.Default.(bool)
	// bookDescNoSale is the schema descriptor for noSale field.
	bookDescNoSale := bookFields[7].Descriptor()
	// book.DefaultNoSale holds the default value on creation for the noSale field.
	book.DefaultNoSale = bookDescNoSale.Default.(int)
	// bookDescNoView is the schema descriptor for noView field.
	bookDescNoView := bookFields[8].Descriptor()
	// book.DefaultNoView holds the default value on creation for the noView field.
	book.DefaultNoView = bookDescNoView.Default.(int)
}
