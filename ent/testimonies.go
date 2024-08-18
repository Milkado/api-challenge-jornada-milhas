// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/testimonies"
)

// Testimonies is the model entity for the Testimonies schema.
type Testimonies struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Testimony holds the value of the "testimony" field.
	Testimony string `json:"testimony,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Picture holds the value of the "picture" field.
	Picture string `json:"picture,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Testimonies) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case testimonies.FieldID:
			values[i] = new(sql.NullInt64)
		case testimonies.FieldTestimony, testimonies.FieldName, testimonies.FieldPicture:
			values[i] = new(sql.NullString)
		case testimonies.FieldCreatedAt, testimonies.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Testimonies fields.
func (t *Testimonies) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case testimonies.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case testimonies.FieldTestimony:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field testimony", values[i])
			} else if value.Valid {
				t.Testimony = value.String
			}
		case testimonies.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case testimonies.FieldPicture:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field picture", values[i])
			} else if value.Valid {
				t.Picture = value.String
			}
		case testimonies.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case testimonies.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Testimonies.
// This includes values selected through modifiers, order, etc.
func (t *Testimonies) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// Update returns a builder for updating this Testimonies.
// Note that you need to call Testimonies.Unwrap() before calling this method if this Testimonies
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Testimonies) Update() *TestimoniesUpdateOne {
	return NewTestimoniesClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Testimonies entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Testimonies) Unwrap() *Testimonies {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Testimonies is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Testimonies) String() string {
	var builder strings.Builder
	builder.WriteString("Testimonies(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("testimony=")
	builder.WriteString(t.Testimony)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("picture=")
	builder.WriteString(t.Picture)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TestimoniesSlice is a parsable slice of Testimonies.
type TestimoniesSlice []*Testimonies
