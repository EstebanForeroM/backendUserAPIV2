package db

import (
	"context"
	"slices"
	"testing"

	"github.com/EstebanForeroM/backendUserAPIV2/ent"
	"github.com/google/uuid"
)

func TestGetProductsForCart(t *testing.T) {
    dataBase := NewTestEntConnection(t) 

    ctx := context.Background()

    cartId := NewUUID(t)

    err := dataBase.Cart.Create().
        SetID(cartId).
        SetTotalPrice(10).
        Exec(ctx)

    if err != nil {
        t.Error("Error creating cart: ", err)
    }

    productsIds := createCartTestProducts(t, *dataBase, cartId)  

    products, err := GetProductsOfCart(dataBase, cartId)

    if err != nil {
        t.Error("Error trying to get carts: ", err)
    }

    for _, prod := range products {
        if !slices.Contains(productsIds, prod.Pid) {
            t.Error("Unvalid uuid found: ", prod.Pid, productsIds)
        }
    }
}

func TestGetProductsForOrder(t *testing.T) {
    dataBase := NewTestEntConnection(t) 

    ctx := context.Background()

    orderId := NewUUID(t)

    err := dataBase.Order.Create().
        SetID(orderId).
        SetDeliveryAdress("idk").
        Exec(ctx)

    if err != nil {
        t.Error("Error creating order: ", err)
    }

    productsIds := createOrderTestProducts(t, *dataBase, orderId)  

    products, err := GetProductsOfOrder(dataBase, orderId)

    if err != nil {
        t.Error("Error trying to get products from Order: ", err)
    }

    for _, prod := range products {
        if !slices.Contains(productsIds, prod.Pid) {
            t.Error("Unvalid uuid found: ", prod.Pid, productsIds)
        }
    }
}

func createCartTestProducts(t *testing.T, dataBase ent.Client, cartId uuid.UUID) []uuid.UUID {
    productsIds := []uuid.UUID {
        NewUUID(t),
        NewUUID(t),
        NewUUID(t),
    }

    ctx := context.Background()

    _, err := dataBase.Product.CreateBulk(
        dataBase.Product.Create().SetCartsID(cartId).SetPid(productsIds[0]),
        dataBase.Product.Create().SetCartsID(cartId).SetPid(productsIds[1]),
        dataBase.Product.Create().SetCartsID(cartId).SetPid(productsIds[2]),
    ).Save(ctx)

    if err != nil {
        t.Error("Unexpected error creating products: ", err)
    }

    return productsIds
}

func createOrderTestProducts(t *testing.T, dataBase ent.Client, orderId uuid.UUID) []uuid.UUID {
    productsIds := []uuid.UUID {
        NewUUID(t),
        NewUUID(t),
        NewUUID(t),
    }

    ctx := context.Background()

    _, err := dataBase.Product.CreateBulk(
        dataBase.Product.Create().SetOrdersID(orderId).SetPid(productsIds[0]),
        dataBase.Product.Create().SetOrdersID(orderId).SetPid(productsIds[1]),
        dataBase.Product.Create().SetOrdersID(orderId).SetPid(productsIds[2]),
    ).Save(ctx)

    if err != nil {
        t.Error("Unexpected error creating products: ", err)
    }

    return productsIds
}

func NewUUID(t *testing.T) uuid.UUID {
    cartUuid, err := uuid.NewUUID()

    if err != nil {
        t.Error("Error generating the new uuid: ", err)
    }

    return cartUuid
}
