package schema

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	gen "github.com/Milkado/api-challenge-jornada-milhas/ent"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/hook"
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
		field.Int("destiny_id"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the Testmonies.
func (Testimonies) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("destinies", Destinies.Type).
			Ref("testimonies").
			Field("destiny_id").
			Unique().
			Required(),
	}
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
