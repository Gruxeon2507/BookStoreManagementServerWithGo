// Code generated by ent, DO NOT EDIT.

package ent

import (
	"bookstore/ent/book"
	"bookstore/ent/category"
	"bookstore/ent/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BookCreate is the builder for creating a Book entity.
type BookCreate struct {
	config
	mutation *BookMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (bc *BookCreate) SetTitle(s string) *BookCreate {
	bc.mutation.SetTitle(s)
	return bc
}

// SetAuthorName sets the "authorName" field.
func (bc *BookCreate) SetAuthorName(s string) *BookCreate {
	bc.mutation.SetAuthorName(s)
	return bc
}

// SetIsApproved sets the "isApproved" field.
func (bc *BookCreate) SetIsApproved(b bool) *BookCreate {
	bc.mutation.SetIsApproved(b)
	return bc
}

// SetNillableIsApproved sets the "isApproved" field if the given value is not nil.
func (bc *BookCreate) SetNillableIsApproved(b *bool) *BookCreate {
	if b != nil {
		bc.SetIsApproved(*b)
	}
	return bc
}

// SetCoverPath sets the "coverPath" field.
func (bc *BookCreate) SetCoverPath(s string) *BookCreate {
	bc.mutation.SetCoverPath(s)
	return bc
}

// SetPdfPath sets the "pdfPath" field.
func (bc *BookCreate) SetPdfPath(s string) *BookCreate {
	bc.mutation.SetPdfPath(s)
	return bc
}

// SetCreatedBy sets the "createdBy" field.
func (bc *BookCreate) SetCreatedBy(i int) *BookCreate {
	bc.mutation.SetCreatedBy(i)
	return bc
}

// SetNillableCreatedBy sets the "createdBy" field if the given value is not nil.
func (bc *BookCreate) SetNillableCreatedBy(i *int) *BookCreate {
	if i != nil {
		bc.SetCreatedBy(*i)
	}
	return bc
}

// SetPrice sets the "price" field.
func (bc *BookCreate) SetPrice(f float64) *BookCreate {
	bc.mutation.SetPrice(f)
	return bc
}

// SetNoSale sets the "noSale" field.
func (bc *BookCreate) SetNoSale(i int) *BookCreate {
	bc.mutation.SetNoSale(i)
	return bc
}

// SetNillableNoSale sets the "noSale" field if the given value is not nil.
func (bc *BookCreate) SetNillableNoSale(i *int) *BookCreate {
	if i != nil {
		bc.SetNoSale(*i)
	}
	return bc
}

// SetNoView sets the "noView" field.
func (bc *BookCreate) SetNoView(i int) *BookCreate {
	bc.mutation.SetNoView(i)
	return bc
}

// SetNillableNoView sets the "noView" field if the given value is not nil.
func (bc *BookCreate) SetNillableNoView(i *int) *BookCreate {
	if i != nil {
		bc.SetNoView(*i)
	}
	return bc
}

// SetDescription sets the "description" field.
func (bc *BookCreate) SetDescription(s string) *BookCreate {
	bc.mutation.SetDescription(s)
	return bc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (bc *BookCreate) SetUserID(id int) *BookCreate {
	bc.mutation.SetUserID(id)
	return bc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (bc *BookCreate) SetNillableUserID(id *int) *BookCreate {
	if id != nil {
		bc = bc.SetUserID(*id)
	}
	return bc
}

// SetUser sets the "user" edge to the User entity.
func (bc *BookCreate) SetUser(u *User) *BookCreate {
	return bc.SetUserID(u.ID)
}

// AddCategoryIDs adds the "category" edge to the Category entity by IDs.
func (bc *BookCreate) AddCategoryIDs(ids ...int) *BookCreate {
	bc.mutation.AddCategoryIDs(ids...)
	return bc
}

// AddCategory adds the "category" edges to the Category entity.
func (bc *BookCreate) AddCategory(c ...*Category) *BookCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bc.AddCategoryIDs(ids...)
}

// Mutation returns the BookMutation object of the builder.
func (bc *BookCreate) Mutation() *BookMutation {
	return bc.mutation
}

// Save creates the Book in the database.
func (bc *BookCreate) Save(ctx context.Context) (*Book, error) {
	bc.defaults()
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BookCreate) SaveX(ctx context.Context) *Book {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BookCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BookCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BookCreate) defaults() {
	if _, ok := bc.mutation.IsApproved(); !ok {
		v := book.DefaultIsApproved
		bc.mutation.SetIsApproved(v)
	}
	if _, ok := bc.mutation.NoSale(); !ok {
		v := book.DefaultNoSale
		bc.mutation.SetNoSale(v)
	}
	if _, ok := bc.mutation.NoView(); !ok {
		v := book.DefaultNoView
		bc.mutation.SetNoView(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BookCreate) check() error {
	if _, ok := bc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Book.title"`)}
	}
	if _, ok := bc.mutation.AuthorName(); !ok {
		return &ValidationError{Name: "authorName", err: errors.New(`ent: missing required field "Book.authorName"`)}
	}
	if _, ok := bc.mutation.IsApproved(); !ok {
		return &ValidationError{Name: "isApproved", err: errors.New(`ent: missing required field "Book.isApproved"`)}
	}
	if _, ok := bc.mutation.CoverPath(); !ok {
		return &ValidationError{Name: "coverPath", err: errors.New(`ent: missing required field "Book.coverPath"`)}
	}
	if _, ok := bc.mutation.PdfPath(); !ok {
		return &ValidationError{Name: "pdfPath", err: errors.New(`ent: missing required field "Book.pdfPath"`)}
	}
	if _, ok := bc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "Book.price"`)}
	}
	if _, ok := bc.mutation.NoSale(); !ok {
		return &ValidationError{Name: "noSale", err: errors.New(`ent: missing required field "Book.noSale"`)}
	}
	if _, ok := bc.mutation.NoView(); !ok {
		return &ValidationError{Name: "noView", err: errors.New(`ent: missing required field "Book.noView"`)}
	}
	if _, ok := bc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Book.description"`)}
	}
	return nil
}

func (bc *BookCreate) sqlSave(ctx context.Context) (*Book, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BookCreate) createSpec() (*Book, *sqlgraph.CreateSpec) {
	var (
		_node = &Book{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(book.Table, sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt))
	)
	if value, ok := bc.mutation.Title(); ok {
		_spec.SetField(book.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := bc.mutation.AuthorName(); ok {
		_spec.SetField(book.FieldAuthorName, field.TypeString, value)
		_node.AuthorName = value
	}
	if value, ok := bc.mutation.IsApproved(); ok {
		_spec.SetField(book.FieldIsApproved, field.TypeBool, value)
		_node.IsApproved = value
	}
	if value, ok := bc.mutation.CoverPath(); ok {
		_spec.SetField(book.FieldCoverPath, field.TypeString, value)
		_node.CoverPath = value
	}
	if value, ok := bc.mutation.PdfPath(); ok {
		_spec.SetField(book.FieldPdfPath, field.TypeString, value)
		_node.PdfPath = value
	}
	if value, ok := bc.mutation.Price(); ok {
		_spec.SetField(book.FieldPrice, field.TypeFloat64, value)
		_node.Price = value
	}
	if value, ok := bc.mutation.NoSale(); ok {
		_spec.SetField(book.FieldNoSale, field.TypeInt, value)
		_node.NoSale = value
	}
	if value, ok := bc.mutation.NoView(); ok {
		_spec.SetField(book.FieldNoView, field.TypeInt, value)
		_node.NoView = value
	}
	if value, ok := bc.mutation.Description(); ok {
		_spec.SetField(book.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := bc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   book.UserTable,
			Columns: []string{book.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CreatedBy = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   book.CategoryTable,
			Columns: book.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BookCreateBulk is the builder for creating many Book entities in bulk.
type BookCreateBulk struct {
	config
	builders []*BookCreate
}

// Save creates the Book entities in the database.
func (bcb *BookCreateBulk) Save(ctx context.Context) ([]*Book, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Book, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BookCreateBulk) SaveX(ctx context.Context) []*Book {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BookCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BookCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
