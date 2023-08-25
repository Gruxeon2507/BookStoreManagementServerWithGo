// Code generated by ent, DO NOT EDIT.

package ent

import (
	"bookstore/ent/book"
	"bookstore/ent/predicate"
	"bookstore/ent/role"
	"bookstore/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetDisplayName sets the "displayName" field.
func (uu *UserUpdate) SetDisplayName(s string) *UserUpdate {
	uu.mutation.SetDisplayName(s)
	return uu
}

// SetDob sets the "dob" field.
func (uu *UserUpdate) SetDob(t time.Time) *UserUpdate {
	uu.mutation.SetDob(t)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetCreatedDate sets the "createdDate" field.
func (uu *UserUpdate) SetCreatedDate(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedDate(t)
	return uu
}

// SetAvatarPath sets the "avatarPath" field.
func (uu *UserUpdate) SetAvatarPath(s string) *UserUpdate {
	uu.mutation.SetAvatarPath(s)
	return uu
}

// AddCreatedByIDs adds the "createdBy" edge to the Book entity by IDs.
func (uu *UserUpdate) AddCreatedByIDs(ids ...int) *UserUpdate {
	uu.mutation.AddCreatedByIDs(ids...)
	return uu
}

// AddCreatedBy adds the "createdBy" edges to the Book entity.
func (uu *UserUpdate) AddCreatedBy(b ...*Book) *UserUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uu.AddCreatedByIDs(ids...)
}

// AddRoleIDs adds the "role" edge to the Role entity by IDs.
func (uu *UserUpdate) AddRoleIDs(ids ...int) *UserUpdate {
	uu.mutation.AddRoleIDs(ids...)
	return uu
}

// AddRole adds the "role" edges to the Role entity.
func (uu *UserUpdate) AddRole(r ...*Role) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.AddRoleIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearCreatedBy clears all "createdBy" edges to the Book entity.
func (uu *UserUpdate) ClearCreatedBy() *UserUpdate {
	uu.mutation.ClearCreatedBy()
	return uu
}

// RemoveCreatedByIDs removes the "createdBy" edge to Book entities by IDs.
func (uu *UserUpdate) RemoveCreatedByIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveCreatedByIDs(ids...)
	return uu
}

// RemoveCreatedBy removes "createdBy" edges to Book entities.
func (uu *UserUpdate) RemoveCreatedBy(b ...*Book) *UserUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uu.RemoveCreatedByIDs(ids...)
}

// ClearRole clears all "role" edges to the Role entity.
func (uu *UserUpdate) ClearRole() *UserUpdate {
	uu.mutation.ClearRole()
	return uu
}

// RemoveRoleIDs removes the "role" edge to Role entities by IDs.
func (uu *UserUpdate) RemoveRoleIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveRoleIDs(ids...)
	return uu
}

// RemoveRole removes "role" edges to Role entities.
func (uu *UserUpdate) RemoveRole(r ...*Role) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.RemoveRoleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.DisplayName(); ok {
		_spec.SetField(user.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Dob(); ok {
		_spec.SetField(user.FieldDob, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.CreatedDate(); ok {
		_spec.SetField(user.FieldCreatedDate, field.TypeTime, value)
	}
	if value, ok := uu.mutation.AvatarPath(); ok {
		_spec.SetField(user.FieldAvatarPath, field.TypeString, value)
	}
	if uu.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CreatedByTable,
			Columns: []string{user.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedCreatedByIDs(); len(nodes) > 0 && !uu.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CreatedByTable,
			Columns: []string{user.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CreatedByTable,
			Columns: []string{user.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.RoleTable,
			Columns: user.RolePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedRoleIDs(); len(nodes) > 0 && !uu.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.RoleTable,
			Columns: user.RolePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.RoleTable,
			Columns: user.RolePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetDisplayName sets the "displayName" field.
func (uuo *UserUpdateOne) SetDisplayName(s string) *UserUpdateOne {
	uuo.mutation.SetDisplayName(s)
	return uuo
}

// SetDob sets the "dob" field.
func (uuo *UserUpdateOne) SetDob(t time.Time) *UserUpdateOne {
	uuo.mutation.SetDob(t)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetCreatedDate sets the "createdDate" field.
func (uuo *UserUpdateOne) SetCreatedDate(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedDate(t)
	return uuo
}

// SetAvatarPath sets the "avatarPath" field.
func (uuo *UserUpdateOne) SetAvatarPath(s string) *UserUpdateOne {
	uuo.mutation.SetAvatarPath(s)
	return uuo
}

// AddCreatedByIDs adds the "createdBy" edge to the Book entity by IDs.
func (uuo *UserUpdateOne) AddCreatedByIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddCreatedByIDs(ids...)
	return uuo
}

// AddCreatedBy adds the "createdBy" edges to the Book entity.
func (uuo *UserUpdateOne) AddCreatedBy(b ...*Book) *UserUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uuo.AddCreatedByIDs(ids...)
}

// AddRoleIDs adds the "role" edge to the Role entity by IDs.
func (uuo *UserUpdateOne) AddRoleIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddRoleIDs(ids...)
	return uuo
}

// AddRole adds the "role" edges to the Role entity.
func (uuo *UserUpdateOne) AddRole(r ...*Role) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.AddRoleIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearCreatedBy clears all "createdBy" edges to the Book entity.
func (uuo *UserUpdateOne) ClearCreatedBy() *UserUpdateOne {
	uuo.mutation.ClearCreatedBy()
	return uuo
}

// RemoveCreatedByIDs removes the "createdBy" edge to Book entities by IDs.
func (uuo *UserUpdateOne) RemoveCreatedByIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveCreatedByIDs(ids...)
	return uuo
}

// RemoveCreatedBy removes "createdBy" edges to Book entities.
func (uuo *UserUpdateOne) RemoveCreatedBy(b ...*Book) *UserUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uuo.RemoveCreatedByIDs(ids...)
}

// ClearRole clears all "role" edges to the Role entity.
func (uuo *UserUpdateOne) ClearRole() *UserUpdateOne {
	uuo.mutation.ClearRole()
	return uuo
}

// RemoveRoleIDs removes the "role" edge to Role entities by IDs.
func (uuo *UserUpdateOne) RemoveRoleIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveRoleIDs(ids...)
	return uuo
}

// RemoveRole removes "role" edges to Role entities.
func (uuo *UserUpdateOne) RemoveRole(r ...*Role) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.RemoveRoleIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.DisplayName(); ok {
		_spec.SetField(user.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Dob(); ok {
		_spec.SetField(user.FieldDob, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.CreatedDate(); ok {
		_spec.SetField(user.FieldCreatedDate, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.AvatarPath(); ok {
		_spec.SetField(user.FieldAvatarPath, field.TypeString, value)
	}
	if uuo.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CreatedByTable,
			Columns: []string{user.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedCreatedByIDs(); len(nodes) > 0 && !uuo.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CreatedByTable,
			Columns: []string{user.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CreatedByTable,
			Columns: []string{user.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.RoleTable,
			Columns: user.RolePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedRoleIDs(); len(nodes) > 0 && !uuo.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.RoleTable,
			Columns: user.RolePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.RoleTable,
			Columns: user.RolePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
