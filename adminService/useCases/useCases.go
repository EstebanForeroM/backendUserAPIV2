package usecases

import (
	"errors"

	"github.com/EstebanForeroM/backendUserAPIV2/property"
	"github.com/google/uuid"
)

type AdminDb interface {
    GetUser(userId string) (User, error)
    GetOrders() ([]Order, error)
    UpdateOrderStatus(orderId uuid.UUID, status property.Status) error
}

func UpdateOrderStatus(db AdminDb, orderId uuid.UUID, status property.Status, userId string) error {
    authorized, err := authorizeAdmin(db, userId)
    if err != nil {
        return err
    }
    if !authorized {
        return errors.New("unauthorized") 
    }

    return db.UpdateOrderStatus(orderId, status)
}

func GetOrders(db AdminDb, adminId string) ([]Order, error) {
    authorized, err := authorizeAdmin(db, adminId)
    if err != nil {
        return nil, err
    }
    if !authorized {
        return nil, errors.New("unauthorized") 
    }

    return db.GetOrders()
}

func GetOrderByUserId(db AdminDb, adminId string, userId string) ([]Order, error) {
    authorized, err := authorizeAdmin(db, adminId)
    if err != nil {
        return nil, err
    }
    if !authorized {
        return nil, errors.New("unauthorized") 
    }

    orders, err := db.GetOrders()
    if err != nil {
        return nil, err
    }

    var userOrders []Order
    for _, order := range orders {
        if order.UserId == userId {
            userOrders = append(userOrders, order)
        }
    }

    return userOrders, nil
}

func GetOrderByStatus(db AdminDb, adminId string, status property.Status) ([]Order, error) {
    authorized, err := authorizeAdmin(db, adminId)
    if err != nil {
        return nil, err
    }
    if !authorized {
        return nil, errors.New("unauthorized") 
    }

    orders, err := db.GetOrders()
    if err != nil {
        return nil, err
    }

    var statusOrders []Order
    for _, order := range orders {
        if order.Status == status {
            statusOrders = append(statusOrders, order)
        }
    }

    return statusOrders, nil
}

func authorizeAdmin(db AdminDb, adminId string) (bool, error) {
    user, err := db.GetUser(adminId)
    if err != nil {
        return false, err
    }
    if user.UserRole != "admin" {
        return false, nil
    }
    return true, nil
}

type User struct {
    UserName string
    UserRole string
}

type Cart struct {
    Products []Product
    TotalPrice float32
}

type Order struct {
    OrderId uuid.UUID
    UserId string
    Status property.Status
    DeliveryAdress string
    Products []Product
}

type Product struct {
    ProductId uuid.UUID
    ProductQuantity uint
}
