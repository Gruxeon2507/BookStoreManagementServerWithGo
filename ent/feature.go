// Code generated by ent, DO NOT EDIT.

package ent

import (
	"bookstore/ent/feature"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Feature is the model entity for the Feature schema.
type Feature struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FeatureName holds the value of the "featureName" field.
	FeatureName string `json:"featureName,omitempty"`
	// FeatureUrl holds the value of the "featureUrl" field.
	FeatureUrl string `json:"featureUrl,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FeatureQuery when eager-loading is set.
	Edges        FeatureEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FeatureEdges holds the relations/edges for other nodes in the graph.
type FeatureEdges struct {
	// Role holds the value of the role edge.
	Role []*Role `json:"role,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading.
func (e FeatureEdges) RoleOrErr() ([]*Role, error) {
	if e.loadedTypes[0] {
		return e.Role, nil
	}
	return nil, &NotLoadedError{edge: "role"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Feature) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case feature.FieldID:
			values[i] = new(sql.NullInt64)
		case feature.FieldFeatureName, feature.FieldFeatureUrl:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Feature fields.
func (f *Feature) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case feature.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case feature.FieldFeatureName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field featureName", values[i])
			} else if value.Valid {
				f.FeatureName = value.String
			}
		case feature.FieldFeatureUrl:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field featureUrl", values[i])
			} else if value.Valid {
				f.FeatureUrl = value.String
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Feature.
// This includes values selected through modifiers, order, etc.
func (f *Feature) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// QueryRole queries the "role" edge of the Feature entity.
func (f *Feature) QueryRole() *RoleQuery {
	return NewFeatureClient(f.config).QueryRole(f)
}

// Update returns a builder for updating this Feature.
// Note that you need to call Feature.Unwrap() before calling this method if this Feature
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Feature) Update() *FeatureUpdateOne {
	return NewFeatureClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Feature entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Feature) Unwrap() *Feature {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Feature is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Feature) String() string {
	var builder strings.Builder
	builder.WriteString("Feature(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("featureName=")
	builder.WriteString(f.FeatureName)
	builder.WriteString(", ")
	builder.WriteString("featureUrl=")
	builder.WriteString(f.FeatureUrl)
	builder.WriteByte(')')
	return builder.String()
}

// Features is a parsable slice of Feature.
type Features []*Feature
