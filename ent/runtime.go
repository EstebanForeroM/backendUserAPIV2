// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/EstebanForeroM/backendUserAPIV2/ent/cart"
	"github.com/EstebanForeroM/backendUserAPIV2/ent/order"
	"github.com/EstebanForeroM/backendUserAPIV2/ent/product"
	"github.com/EstebanForeroM/backendUserAPIV2/ent/schema"
	"github.com/EstebanForeroM/backendUserAPIV2/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cartFields := schema.Cart{}.Fields()
	_ = cartFields
	// cartDescTotalPrice is the schema descriptor for total_price field.
	cartDescTotalPrice := cartFields[1].Descriptor()
	// cart.DefaultTotalPrice holds the default value on creation for the total_price field.
	cart.DefaultTotalPrice = cartDescTotalPrice.Default.(float32)
	// cartDescID is the schema descriptor for id field.
	cartDescID := cartFields[0].Descriptor()
	// cart.DefaultID holds the default value on creation for the id field.
	cart.DefaultID = cartDescID.Default.(func() uuid.UUID)
	orderFields := schema.Order{}.Fields()
	_ = orderFields
	// orderDescOrderedAt is the schema descriptor for ordered_at field.
	orderDescOrderedAt := orderFields[1].Descriptor()
	// order.DefaultOrderedAt holds the default value on creation for the ordered_at field.
	order.DefaultOrderedAt = orderDescOrderedAt.Default.(func() time.Time)
	// orderDescDeliveryAdress is the schema descriptor for delivery_adress field.
	orderDescDeliveryAdress := orderFields[2].Descriptor()
	// order.DeliveryAdressValidator is a validator for the "delivery_adress" field. It is called by the builders before save.
	order.DeliveryAdressValidator = orderDescDeliveryAdress.Validators[0].(func(string) error)
	// orderDescID is the schema descriptor for id field.
	orderDescID := orderFields[0].Descriptor()
	// order.DefaultID holds the default value on creation for the id field.
	order.DefaultID = orderDescID.Default.(func() uuid.UUID)
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescQuantity is the schema descriptor for quantity field.
	productDescQuantity := productFields[1].Descriptor()
	// product.DefaultQuantity holds the default value on creation for the quantity field.
	product.DefaultQuantity = productDescQuantity.Default.(uint8)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = func() func(string) error {
		validators := userDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = func() func(string) error {
		validators := userDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
