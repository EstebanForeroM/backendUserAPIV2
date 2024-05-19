// Code generated by ent, DO NOT EDIT.

package order

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/EstebanForeroM/backendUserAPIV2/ent/predicate"
	"github.com/EstebanForeroM/backendUserAPIV2/property"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldID, id))
}

// OrderedAt applies equality check predicate on the "ordered_at" field. It's identical to OrderedAtEQ.
func OrderedAt(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldOrderedAt, v))
}

// DeliveryAdress applies equality check predicate on the "delivery_adress" field. It's identical to DeliveryAdressEQ.
func DeliveryAdress(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldDeliveryAdress, v))
}

// OrderedAtEQ applies the EQ predicate on the "ordered_at" field.
func OrderedAtEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldOrderedAt, v))
}

// OrderedAtNEQ applies the NEQ predicate on the "ordered_at" field.
func OrderedAtNEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldOrderedAt, v))
}

// OrderedAtIn applies the In predicate on the "ordered_at" field.
func OrderedAtIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldOrderedAt, vs...))
}

// OrderedAtNotIn applies the NotIn predicate on the "ordered_at" field.
func OrderedAtNotIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldOrderedAt, vs...))
}

// OrderedAtGT applies the GT predicate on the "ordered_at" field.
func OrderedAtGT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldOrderedAt, v))
}

// OrderedAtGTE applies the GTE predicate on the "ordered_at" field.
func OrderedAtGTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldOrderedAt, v))
}

// OrderedAtLT applies the LT predicate on the "ordered_at" field.
func OrderedAtLT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldOrderedAt, v))
}

// OrderedAtLTE applies the LTE predicate on the "ordered_at" field.
func OrderedAtLTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldOrderedAt, v))
}

// DeliveryAdressEQ applies the EQ predicate on the "delivery_adress" field.
func DeliveryAdressEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldDeliveryAdress, v))
}

// DeliveryAdressNEQ applies the NEQ predicate on the "delivery_adress" field.
func DeliveryAdressNEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldDeliveryAdress, v))
}

// DeliveryAdressIn applies the In predicate on the "delivery_adress" field.
func DeliveryAdressIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldDeliveryAdress, vs...))
}

// DeliveryAdressNotIn applies the NotIn predicate on the "delivery_adress" field.
func DeliveryAdressNotIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldDeliveryAdress, vs...))
}

// DeliveryAdressGT applies the GT predicate on the "delivery_adress" field.
func DeliveryAdressGT(v string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldDeliveryAdress, v))
}

// DeliveryAdressGTE applies the GTE predicate on the "delivery_adress" field.
func DeliveryAdressGTE(v string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldDeliveryAdress, v))
}

// DeliveryAdressLT applies the LT predicate on the "delivery_adress" field.
func DeliveryAdressLT(v string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldDeliveryAdress, v))
}

// DeliveryAdressLTE applies the LTE predicate on the "delivery_adress" field.
func DeliveryAdressLTE(v string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldDeliveryAdress, v))
}

// DeliveryAdressContains applies the Contains predicate on the "delivery_adress" field.
func DeliveryAdressContains(v string) predicate.Order {
	return predicate.Order(sql.FieldContains(FieldDeliveryAdress, v))
}

// DeliveryAdressHasPrefix applies the HasPrefix predicate on the "delivery_adress" field.
func DeliveryAdressHasPrefix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasPrefix(FieldDeliveryAdress, v))
}

// DeliveryAdressHasSuffix applies the HasSuffix predicate on the "delivery_adress" field.
func DeliveryAdressHasSuffix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasSuffix(FieldDeliveryAdress, v))
}

// DeliveryAdressEqualFold applies the EqualFold predicate on the "delivery_adress" field.
func DeliveryAdressEqualFold(v string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldDeliveryAdress, v))
}

// DeliveryAdressContainsFold applies the ContainsFold predicate on the "delivery_adress" field.
func DeliveryAdressContainsFold(v string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldDeliveryAdress, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v property.Status) predicate.Order {
	vc := v
	return predicate.Order(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v property.Status) predicate.Order {
	vc := v
	return predicate.Order(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...property.Status) predicate.Order {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...property.Status) predicate.Order {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(sql.FieldNotIn(FieldStatus, v...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProducts applies the HasEdge predicate on the "products" edge.
func HasProducts() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductsTable, ProductsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductsWith applies the HasEdge predicate on the "products" edge with a given conditions (other predicates).
func HasProductsWith(preds ...predicate.Product) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newProductsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Order) predicate.Order {
	return predicate.Order(sql.NotPredicates(p))
}
