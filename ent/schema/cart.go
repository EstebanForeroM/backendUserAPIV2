package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Cart holds the schema definition for the Cart entity.
type Cart struct {
	ent.Schema
}

// Fields of the Cart.
func (Cart) Fields() []ent.Field {
	return []ent.Field {
        field.UUID("id", uuid.UUID{}).
            Default(uuid.New).
            StorageKey("oid"),
        field.Float32("total_price").
            Default(0),

    }
}

// Edges of the Cart.
func (Cart) Edges() []ent.Edge {
	return []ent.Edge {
        edge.From("user", User.Type).
            Ref("cart").
            Unique(),
        edge.To("products", Product.Type),
    }
}
