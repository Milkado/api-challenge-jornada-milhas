package schema

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/hook"
	gen "github.com/Milkado/api-challenge-jornada-milhas/ent"
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
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the Testmonies.
func (Testimonies) Edges() []ent.Edge {
	return nil
}

func (Testimonies) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.TestimoniesFunc(func(ctx context.Context, m *gen.TestimoniesMutation) (ent.Value, error) {
					if anyChanged := len(m.Fields()) > 0; anyChanged {
						m.SetUpdatedAt(time.Now())
					}

					return next.Mutate(ctx, m)
				})
			},
			ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
