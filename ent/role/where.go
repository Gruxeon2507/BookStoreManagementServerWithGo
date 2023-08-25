// Code generated by ent, DO NOT EDIT.

package role

import (
	"bookstore/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldID, id))
}

// RoleName applies equality check predicate on the "roleName" field. It's identical to RoleNameEQ.
func RoleName(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldRoleName, v))
}

// RoleNameEQ applies the EQ predicate on the "roleName" field.
func RoleNameEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldRoleName, v))
}

// RoleNameNEQ applies the NEQ predicate on the "roleName" field.
func RoleNameNEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldRoleName, v))
}

// RoleNameIn applies the In predicate on the "roleName" field.
func RoleNameIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldRoleName, vs...))
}

// RoleNameNotIn applies the NotIn predicate on the "roleName" field.
func RoleNameNotIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldRoleName, vs...))
}

// RoleNameGT applies the GT predicate on the "roleName" field.
func RoleNameGT(v string) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldRoleName, v))
}

// RoleNameGTE applies the GTE predicate on the "roleName" field.
func RoleNameGTE(v string) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldRoleName, v))
}

// RoleNameLT applies the LT predicate on the "roleName" field.
func RoleNameLT(v string) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldRoleName, v))
}

// RoleNameLTE applies the LTE predicate on the "roleName" field.
func RoleNameLTE(v string) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldRoleName, v))
}

// RoleNameContains applies the Contains predicate on the "roleName" field.
func RoleNameContains(v string) predicate.Role {
	return predicate.Role(sql.FieldContains(FieldRoleName, v))
}

// RoleNameHasPrefix applies the HasPrefix predicate on the "roleName" field.
func RoleNameHasPrefix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasPrefix(FieldRoleName, v))
}

// RoleNameHasSuffix applies the HasSuffix predicate on the "roleName" field.
func RoleNameHasSuffix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasSuffix(FieldRoleName, v))
}

// RoleNameEqualFold applies the EqualFold predicate on the "roleName" field.
func RoleNameEqualFold(v string) predicate.Role {
	return predicate.Role(sql.FieldEqualFold(FieldRoleName, v))
}

// RoleNameContainsFold applies the ContainsFold predicate on the "roleName" field.
func RoleNameContainsFold(v string) predicate.Role {
	return predicate.Role(sql.FieldContainsFold(FieldRoleName, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFeature applies the HasEdge predicate on the "feature" edge.
func HasFeature() predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, FeatureTable, FeaturePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFeatureWith applies the HasEdge predicate on the "feature" edge with a given conditions (other predicates).
func HasFeatureWith(preds ...predicate.Feature) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := newFeatureStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Role) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Role) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
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
func Not(p predicate.Role) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		p(s.Not())
	})
}
