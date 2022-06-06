package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Schedule holds the schema definition for the Schedule entity.
type Schedule struct {
	ent.Schema
}

// Fields of the Schedule.
func (Schedule) Fields() []ent.Field {
	return Fields{
		field.String("slack_channel_id").
			Unique(),
		field.String("team_slack_id"),
		field.String("name"),
		field.Enum("interval").
			Values("daily", "weekly", "bi-weekly", "monthly"),
		field.Bool("is_enabled").Default(true),
		field.Time("end_time"),
		field.Time("start_time"),
		field.Bool("weekdays_only"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			Optional(),
	}
}

func (Schedule) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("is_enabled"),
		index.Fields("team_slack_id"),
		index.Fields("slack_channel_id"),
	}
}

// Edges of the Schedule.
func (Schedule) Edges() []ent.Edge {
	return Edges{
		edge.To("shifts", Shift.Type),
	}
}
