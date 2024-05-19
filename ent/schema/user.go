package schema

import (
	"regexp"

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
	return []ent.Field {
        field.String("id").
            Match(regexp.MustCompile("^user_[a-zA-Z0-9]+$")).
            NotEmpty().
            Unique(),
        field.String("name").
            MaxLen(25).
            NotEmpty(),
        field.String("email").
            NotEmpty().
            Unique(),
        field.Enum("role").
            Values("user", "admin").
            Default("user"),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge {
        edge.To("orders", Order.Type),
        edge.To("cart", Cart.Type).
            Unique(),
    }
}
