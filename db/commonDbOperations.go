package db

import (
	"context"

	. "github.com/EstebanForeroM/backendUserAPIV2/ent"
	"github.com/EstebanForeroM/backendUserAPIV2/ent/cart"
	"github.com/EstebanForeroM/backendUserAPIV2/ent/order"
	"github.com/google/uuid"
)

func GetProductsOfCart(dataBase *Client, cartId uuid.UUID) ([]*Product, error) {
    ctx := context.Background()

    return dataBase.Cart.Query().
        Where(cart.ID(cartId)).
        QueryProducts().
        All(ctx)
}

func GetProductsOfOrder(dataBase *Client, orderId uuid.UUID) ([]*Product, error) {
    ctx := context.Background()

    return dataBase.Order.Query().
        Where(order.ID(orderId)).
        QueryProducts().
        All(ctx)
}
