package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/EstebanForeroM/backendUserAPIV2/property"
	"github.com/google/uuid"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field {
        field.UUID("id", uuid.UUID{}).
            Default(uuid.New).
            StorageKey("oid"),
        field.Time("ordered_at").
            Default(time.Now),
        field.String("delivery_adress").
            NotEmpty(),
        field.Enum("status").
            GoType(property.Status("")).
            Default(string(property.Pending)),
    }
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge {
        edge.From("user", User.Type).
            Ref("orders").
            Unique(),
        edge.To("products", Product.Type),
    }
}
