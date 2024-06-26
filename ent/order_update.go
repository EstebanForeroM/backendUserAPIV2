// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/EstebanForeroM/backendUserAPIV2/ent/order"
	"github.com/EstebanForeroM/backendUserAPIV2/ent/predicate"
	"github.com/EstebanForeroM/backendUserAPIV2/ent/product"
	"github.com/EstebanForeroM/backendUserAPIV2/ent/user"
	"github.com/EstebanForeroM/backendUserAPIV2/property"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetOrderedAt sets the "ordered_at" field.
func (ou *OrderUpdate) SetOrderedAt(t time.Time) *OrderUpdate {
	ou.mutation.SetOrderedAt(t)
	return ou
}

// SetNillableOrderedAt sets the "ordered_at" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableOrderedAt(t *time.Time) *OrderUpdate {
	if t != nil {
		ou.SetOrderedAt(*t)
	}
	return ou
}

// SetDeliveryAdress sets the "delivery_adress" field.
func (ou *OrderUpdate) SetDeliveryAdress(s string) *OrderUpdate {
	ou.mutation.SetDeliveryAdress(s)
	return ou
}

// SetNillableDeliveryAdress sets the "delivery_adress" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableDeliveryAdress(s *string) *OrderUpdate {
	if s != nil {
		ou.SetDeliveryAdress(*s)
	}
	return ou
}

// SetStatus sets the "status" field.
func (ou *OrderUpdate) SetStatus(pr property.Status) *OrderUpdate {
	ou.mutation.SetStatus(pr)
	return ou
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableStatus(pr *property.Status) *OrderUpdate {
	if pr != nil {
		ou.SetStatus(*pr)
	}
	return ou
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ou *OrderUpdate) SetUserID(id string) *OrderUpdate {
	ou.mutation.SetUserID(id)
	return ou
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ou *OrderUpdate) SetNillableUserID(id *string) *OrderUpdate {
	if id != nil {
		ou = ou.SetUserID(*id)
	}
	return ou
}

// SetUser sets the "user" edge to the User entity.
func (ou *OrderUpdate) SetUser(u *User) *OrderUpdate {
	return ou.SetUserID(u.ID)
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (ou *OrderUpdate) AddProductIDs(ids ...int) *OrderUpdate {
	ou.mutation.AddProductIDs(ids...)
	return ou
}

// AddProducts adds the "products" edges to the Product entity.
func (ou *OrderUpdate) AddProducts(p ...*Product) *OrderUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ou.AddProductIDs(ids...)
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ou *OrderUpdate) ClearUser() *OrderUpdate {
	ou.mutation.ClearUser()
	return ou
}

// ClearProducts clears all "products" edges to the Product entity.
func (ou *OrderUpdate) ClearProducts() *OrderUpdate {
	ou.mutation.ClearProducts()
	return ou
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (ou *OrderUpdate) RemoveProductIDs(ids ...int) *OrderUpdate {
	ou.mutation.RemoveProductIDs(ids...)
	return ou
}

// RemoveProducts removes "products" edges to Product entities.
func (ou *OrderUpdate) RemoveProducts(p ...*Product) *OrderUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ou.RemoveProductIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrderUpdate) check() error {
	if v, ok := ou.mutation.DeliveryAdress(); ok {
		if err := order.DeliveryAdressValidator(v); err != nil {
			return &ValidationError{Name: "delivery_adress", err: fmt.Errorf(`ent: validator failed for field "Order.delivery_adress": %w`, err)}
		}
	}
	if v, ok := ou.mutation.Status(); ok {
		if err := order.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Order.status": %w`, err)}
		}
	}
	return nil
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.OrderedAt(); ok {
		_spec.SetField(order.FieldOrderedAt, field.TypeTime, value)
	}
	if value, ok := ou.mutation.DeliveryAdress(); ok {
		_spec.SetField(order.FieldDeliveryAdress, field.TypeString, value)
	}
	if value, ok := ou.mutation.Status(); ok {
		_spec.SetField(order.FieldStatus, field.TypeEnum, value)
	}
	if ou.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.ProductsTable,
			Columns: []string{order.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedProductsIDs(); len(nodes) > 0 && !ou.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.ProductsTable,
			Columns: []string{order.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.ProductsTable,
			Columns: []string{order.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderMutation
}

// SetOrderedAt sets the "ordered_at" field.
func (ouo *OrderUpdateOne) SetOrderedAt(t time.Time) *OrderUpdateOne {
	ouo.mutation.SetOrderedAt(t)
	return ouo
}

// SetNillableOrderedAt sets the "ordered_at" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableOrderedAt(t *time.Time) *OrderUpdateOne {
	if t != nil {
		ouo.SetOrderedAt(*t)
	}
	return ouo
}

// SetDeliveryAdress sets the "delivery_adress" field.
func (ouo *OrderUpdateOne) SetDeliveryAdress(s string) *OrderUpdateOne {
	ouo.mutation.SetDeliveryAdress(s)
	return ouo
}

// SetNillableDeliveryAdress sets the "delivery_adress" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableDeliveryAdress(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetDeliveryAdress(*s)
	}
	return ouo
}

// SetStatus sets the "status" field.
func (ouo *OrderUpdateOne) SetStatus(pr property.Status) *OrderUpdateOne {
	ouo.mutation.SetStatus(pr)
	return ouo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableStatus(pr *property.Status) *OrderUpdateOne {
	if pr != nil {
		ouo.SetStatus(*pr)
	}
	return ouo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ouo *OrderUpdateOne) SetUserID(id string) *OrderUpdateOne {
	ouo.mutation.SetUserID(id)
	return ouo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableUserID(id *string) *OrderUpdateOne {
	if id != nil {
		ouo = ouo.SetUserID(*id)
	}
	return ouo
}

// SetUser sets the "user" edge to the User entity.
func (ouo *OrderUpdateOne) SetUser(u *User) *OrderUpdateOne {
	return ouo.SetUserID(u.ID)
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (ouo *OrderUpdateOne) AddProductIDs(ids ...int) *OrderUpdateOne {
	ouo.mutation.AddProductIDs(ids...)
	return ouo
}

// AddProducts adds the "products" edges to the Product entity.
func (ouo *OrderUpdateOne) AddProducts(p ...*Product) *OrderUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ouo.AddProductIDs(ids...)
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ouo *OrderUpdateOne) ClearUser() *OrderUpdateOne {
	ouo.mutation.ClearUser()
	return ouo
}

// ClearProducts clears all "products" edges to the Product entity.
func (ouo *OrderUpdateOne) ClearProducts() *OrderUpdateOne {
	ouo.mutation.ClearProducts()
	return ouo
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (ouo *OrderUpdateOne) RemoveProductIDs(ids ...int) *OrderUpdateOne {
	ouo.mutation.RemoveProductIDs(ids...)
	return ouo
}

// RemoveProducts removes "products" edges to Product entities.
func (ouo *OrderUpdateOne) RemoveProducts(p ...*Product) *OrderUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ouo.RemoveProductIDs(ids...)
}

// Where appends a list predicates to the OrderUpdate builder.
func (ouo *OrderUpdateOne) Where(ps ...predicate.Order) *OrderUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrderUpdateOne) check() error {
	if v, ok := ouo.mutation.DeliveryAdress(); ok {
		if err := order.DeliveryAdressValidator(v); err != nil {
			return &ValidationError{Name: "delivery_adress", err: fmt.Errorf(`ent: validator failed for field "Order.delivery_adress": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.Status(); ok {
		if err := order.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Order.status": %w`, err)}
		}
	}
	return nil
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.OrderedAt(); ok {
		_spec.SetField(order.FieldOrderedAt, field.TypeTime, value)
	}
	if value, ok := ouo.mutation.DeliveryAdress(); ok {
		_spec.SetField(order.FieldDeliveryAdress, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Status(); ok {
		_spec.SetField(order.FieldStatus, field.TypeEnum, value)
	}
	if ouo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.ProductsTable,
			Columns: []string{order.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedProductsIDs(); len(nodes) > 0 && !ouo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.ProductsTable,
			Columns: []string{order.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.ProductsTable,
			Columns: []string{order.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
