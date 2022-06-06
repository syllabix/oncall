package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Shift holds the schema definition for the Shift entity.
type Shift struct {
	ent.Schema
}

// Fields of the Shift.
func (Shift) Fields() []ent.Field {
	return Fields{
		field.Int("sequence_id").
			Optional().
			SchemaType(
				map[string]string{
					dialect.Postgres: "serial",
				},
			),
		field.Enum("status").
			Values("active", "override").
			Optional(),
		field.Int("user_id"),
		field.Int("schedule_id"),
		field.Time("started_at").
			Default(time.Now).
			Optional(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Shift) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sequence_id"),
		index.Fields("user_id", "schedule_id").
			Unique(),
	}
}

// Edges of the Shift.
func (Shift) Edges() []ent.Edge {
	return Edges{
		edge.From("user", User.Type).
			Ref("shifts").
			Field("user_id").
			Unique().
			Required().
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.From("schedule", Schedule.Type).
			Ref("shifts").
			Field("schedule_id").
			Unique().
			Required().
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
