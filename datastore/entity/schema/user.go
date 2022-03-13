package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return Fields{
		field.String("slack_id").Unique(),
		field.String("slack_handle").Unique(),
		field.String("email"),
		field.String("first_name"),
		field.String("last_name"),
		field.String("avatar_url"),
		field.String("display_name"),
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

// Edges of the User.
func (User) Edges() []ent.Edge {
	return Edges{
		edge.To("shifts", Shift.Type),
	}
}
