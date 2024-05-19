package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field {
        field.UUID("pid", uuid.UUID{}).
            StorageKey("pid"),
        field.Uint8("quantity").
            Default(1),
    }
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge {
        edge.From("orders", Order.Type).
            Ref("products").
            Unique(),
        edge.From("carts", Cart.Type).
            Ref("products").
            Unique(),
    }
}
