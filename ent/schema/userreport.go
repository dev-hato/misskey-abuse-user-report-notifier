package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// UserReport holds the schema definition for the UserReport entity.
type UserReport struct {
	ent.Schema
}

// Fields of the UserReport.
func (UserReport) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("id"),
		field.Time("created_at").Immutable().Default(time.Now),
	}
}

// Edges of the UserReport.
func (UserReport) Edges() []ent.Edge {
	return nil
}
