package database

import (
	"context"

	usecases "github.com/EstebanForeroM/backendUserAPIV2/adminService/useCases"
	"github.com/EstebanForeroM/backendUserAPIV2/db"
	"github.com/EstebanForeroM/backendUserAPIV2/ent"
	"github.com/EstebanForeroM/backendUserAPIV2/ent/user"
	"github.com/EstebanForeroM/backendUserAPIV2/property"
	"github.com/google/uuid"
)

type DataBase struct {
    clientDb *ent.Client
}

func NewDataBase() DataBase {
    return DataBase{
        clientDb: db.NewEntConnection(),
    }
}

func (dataBase *DataBase) UpdateOrderStatus(orderId uuid.UUID, status property.Status) error {
    _, err := dataBase.clientDb.Order.UpdateOneID(orderId).SetStatus(status).Save(context.Background())
    if err != nil {
        return err
    }

    return nil
}

func (dataBase *DataBase) GetUser(userId string) (usecases.User, error) {
    user, err := dataBase.clientDb.User.Query().Where(user.ID(userId)).Only(context.Background())
    if err != nil {
        return usecases.User{}, err
    }

    return usecases.User{
        UserName: user.Name,
        UserRole: string(user.Role),
    }, nil
}

func (dataBase *DataBase) GetOrders() ([]usecases.Order, error) {
    orders, err := dataBase.clientDb.Order.Query().All(context.Background())
    if err != nil {
        return nil, err
    }

    var userOrders []usecases.Order
    for _, order := range orders {

        productsEnt, err := order.QueryProducts().All(context.Background())

        products := entProductsToUseCaseProducts(productsEnt)

        if err != nil {
            return nil, err
        }

        userOrders = append(userOrders, usecases.Order{
            OrderId: order.ID,
            UserId: order.Edges.User.ID,
            Status: order.Status,
            DeliveryAdress: order.DeliveryAdress,
            Products: products,
        })
    }

    return userOrders, nil
}

func entProductsToUseCaseProducts(products []*ent.Product) []usecases.Product {
    var useCaseProducts []usecases.Product
    for _, product := range products {
        useCaseProducts = append(useCaseProducts, usecases.Product{
            ProductId: product.Pid,
            ProductQuantity: uint(product.Quantity),
        })
    }

    return useCaseProducts
}
