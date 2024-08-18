package schema

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/hook"
	gen "github.com/Milkado/api-challenge-jornada-milhas/ent"
)

// Destinies holds the schema definition for the Destinies entity.
type Destinies struct {
	ent.Schema
}

// Fields of the Destinies.
func (Destinies) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Float("price"),
		field.String("meta").MaxLen(160),
		field.Text("description").Optional().Nillable(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the Destinies.
func (Destinies) Edges() []ent.Edge {
	return nil
}

// Hooks of the Destinies.
func (Destinies) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.DestiniesFunc(func(ctx context.Context, m *gen.DestiniesMutation) (ent.Value, error) {
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
