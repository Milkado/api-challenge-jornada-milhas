package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DestinyPictures holds the schema definition for the DestinyPictures entity.
type DestinyPictures struct {
	ent.Schema
}

// Fields of the DestinyPictures.
func (DestinyPictures) Fields() []ent.Field {
	return []ent.Field{
		field.String("picture"),
		field.String("path"),
		field.Int("destiny_id"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the DestinyPictures.
func (DestinyPictures) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("destinies", Destinies.Type).
			Ref("destiny_pictures").
			Field("destiny_id").
			Unique().
			Required(),
	}
}
