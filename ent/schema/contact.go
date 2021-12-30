package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Contact holds the schema definition for the Contact entity.
type Contact struct {
	ent.Schema
}

// Fields of the Contact.
func (Contact) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Unique(),
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New),
		field.String("name"),
		field.String("phone"),
		field.Time("create_time").
			Default(time.Now),
		field.Time("update_time").
			Default(time.Now),
	}
}

// Edges of the Contact.
func (Contact) Edges() []ent.Edge {
	return nil
}
