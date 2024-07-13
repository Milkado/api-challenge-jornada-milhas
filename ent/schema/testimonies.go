package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Testmonies holds the schema definition for the Testmonies entity.
type Testimonies struct {
	ent.Schema
}

// Fields of the Testmonies.
func (Testimonies) Fields() []ent.Field {
	return []ent.Field{
		field.Text("testimony"),
		field.String("name"),
		field.String("picture").Unique(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the Testmonies.
func (Testimonies) Edges() []ent.Edge {
	return nil
}
