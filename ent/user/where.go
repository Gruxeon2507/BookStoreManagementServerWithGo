// Code generated by ent, DO NOT EDIT.

package user

import (
	"bookstore/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// DisplayName applies equality check predicate on the "displayName" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDisplayName, v))
}

// Dob applies equality check predicate on the "dob" field. It's identical to DobEQ.
func Dob(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDob, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// CreatedDate applies equality check predicate on the "createdDate" field. It's identical to CreatedDateEQ.
func CreatedDate(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedDate, v))
}

// AvatarPath applies equality check predicate on the "avatarPath" field. It's identical to AvatarPathEQ.
func AvatarPath(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatarPath, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUsername, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// DisplayNameEQ applies the EQ predicate on the "displayName" field.
func DisplayNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "displayName" field.
func DisplayNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "displayName" field.
func DisplayNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "displayName" field.
func DisplayNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "displayName" field.
func DisplayNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "displayName" field.
func DisplayNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "displayName" field.
func DisplayNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "displayName" field.
func DisplayNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "displayName" field.
func DisplayNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "displayName" field.
func DisplayNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "displayName" field.
func DisplayNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "displayName" field.
func DisplayNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "displayName" field.
func DisplayNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldDisplayName, v))
}

// DobEQ applies the EQ predicate on the "dob" field.
func DobEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDob, v))
}

// DobNEQ applies the NEQ predicate on the "dob" field.
func DobNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDob, v))
}

// DobIn applies the In predicate on the "dob" field.
func DobIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldDob, vs...))
}

// DobNotIn applies the NotIn predicate on the "dob" field.
func DobNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDob, vs...))
}

// DobGT applies the GT predicate on the "dob" field.
func DobGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldDob, v))
}

// DobGTE applies the GTE predicate on the "dob" field.
func DobGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldDob, v))
}

// DobLT applies the LT predicate on the "dob" field.
func DobLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldDob, v))
}

// DobLTE applies the LTE predicate on the "dob" field.
func DobLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldDob, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// CreatedDateEQ applies the EQ predicate on the "createdDate" field.
func CreatedDateEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedDate, v))
}

// CreatedDateNEQ applies the NEQ predicate on the "createdDate" field.
func CreatedDateNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedDate, v))
}

// CreatedDateIn applies the In predicate on the "createdDate" field.
func CreatedDateIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedDate, vs...))
}

// CreatedDateNotIn applies the NotIn predicate on the "createdDate" field.
func CreatedDateNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedDate, vs...))
}

// CreatedDateGT applies the GT predicate on the "createdDate" field.
func CreatedDateGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedDate, v))
}

// CreatedDateGTE applies the GTE predicate on the "createdDate" field.
func CreatedDateGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedDate, v))
}

// CreatedDateLT applies the LT predicate on the "createdDate" field.
func CreatedDateLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedDate, v))
}

// CreatedDateLTE applies the LTE predicate on the "createdDate" field.
func CreatedDateLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedDate, v))
}

// AvatarPathEQ applies the EQ predicate on the "avatarPath" field.
func AvatarPathEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatarPath, v))
}

// AvatarPathNEQ applies the NEQ predicate on the "avatarPath" field.
func AvatarPathNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAvatarPath, v))
}

// AvatarPathIn applies the In predicate on the "avatarPath" field.
func AvatarPathIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldAvatarPath, vs...))
}

// AvatarPathNotIn applies the NotIn predicate on the "avatarPath" field.
func AvatarPathNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAvatarPath, vs...))
}

// AvatarPathGT applies the GT predicate on the "avatarPath" field.
func AvatarPathGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldAvatarPath, v))
}

// AvatarPathGTE applies the GTE predicate on the "avatarPath" field.
func AvatarPathGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAvatarPath, v))
}

// AvatarPathLT applies the LT predicate on the "avatarPath" field.
func AvatarPathLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldAvatarPath, v))
}

// AvatarPathLTE applies the LTE predicate on the "avatarPath" field.
func AvatarPathLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAvatarPath, v))
}

// AvatarPathContains applies the Contains predicate on the "avatarPath" field.
func AvatarPathContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldAvatarPath, v))
}

// AvatarPathHasPrefix applies the HasPrefix predicate on the "avatarPath" field.
func AvatarPathHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldAvatarPath, v))
}

// AvatarPathHasSuffix applies the HasSuffix predicate on the "avatarPath" field.
func AvatarPathHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldAvatarPath, v))
}

// AvatarPathEqualFold applies the EqualFold predicate on the "avatarPath" field.
func AvatarPathEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldAvatarPath, v))
}

// AvatarPathContainsFold applies the ContainsFold predicate on the "avatarPath" field.
func AvatarPathContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldAvatarPath, v))
}

// HasCreatedBy applies the HasEdge predicate on the "createdBy" edge.
func HasCreatedBy() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CreatedByTable, CreatedByColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatedByWith applies the HasEdge predicate on the "createdBy" edge with a given conditions (other predicates).
func HasCreatedByWith(preds ...predicate.Book) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newCreatedByStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRole applies the HasEdge predicate on the "role" edge.
func HasRole() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RoleTable, RolePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoleWith applies the HasEdge predicate on the "role" edge with a given conditions (other predicates).
func HasRoleWith(preds ...predicate.Role) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newRoleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
