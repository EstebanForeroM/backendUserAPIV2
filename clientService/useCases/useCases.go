package usecases

import (
	"errors"
	"log"

	productapi "github.com/EstebanForeroM/backendUserAPIV2/productAPI"
	"github.com/EstebanForeroM/backendUserAPIV2/property"
	"github.com/google/uuid"
)

type ClientDb interface {
    GetCart(userId string) (Cart, error)
    GetUser(userId string) (User, error)
    GetOrders(userId string) ([]Order, error)
    AddOrder(userId string, deliveryAdress string) (error)
    AddProductToCart(userId string, productId uuid.UUID, price float32) (error)
    AddNewProductToCart(userId string, productId uuid.UUID, price float32) (error)
    AddCartToUser(userId string) (error)
    DeleteProductFromCart(userId string, productId uuid.UUID, price float32) (error)
    UserHasCart(userId string) (bool, error)
    CartIsEmpty(userId string) (bool, error)
    DeleteCart(userId string) (error)
    CartHasProduct(userId string, productId uuid.UUID) (bool, error)
    GetProductInfo(productId uuid.UUID) (Product, error)
}

func GetOrders(db ClientDb, userId string) ([]Order, error) {
    return db.GetOrders(userId)
}

func CreateOrder(db ClientDb, userId string, deliveryAdress string) error {
    if res, err := db.UserHasCart(userId); err != nil {
        return err
    } else if !res {
        return errors.New("User has no cart")
    }

    cart, err := db.GetCart(userId)

    if err != nil {
        return err
    }

    if len(cart.Products) == 0 {
        return nil
    }

    if err := db.AddOrder(userId, deliveryAdress); err != nil {
        return err
    }

    return db.DeleteCart(userId)
}

func GetCart(db ClientDb, userId string) (Cart, error) {
    return db.GetCart(userId)
}

func AddProductToCart(db ClientDb, userId string, productId uuid.UUID) error {

    productInfo, err := productapi.GetProductInfo(productId)

    if err != nil {
        return err 
    }

    log.Printf("Adding product to cart: %s %s %f", userId, productId, productInfo.Price)
    if res, err := db.UserHasCart(userId); err != nil {
        return err
    } else if !res { 
        db.AddCartToUser(userId)
    }

    if res, err := db.CartIsEmpty(userId); err != nil {
        return err
    } else if res {
        db.AddNewProductToCart(userId, productId, productInfo.Price)
    } else {
        db.AddProductToCart(userId, productId, productInfo.Price)
    }

    return nil
}

func DeleteProductFromCart(db ClientDb, userId string, productId uuid.UUID) error {
    productInfo, err := productapi.GetProductInfo(productId)

    if err != nil {
        return err
    }

    log.Printf("Deleting product from cart: %s %s", userId, productId)
    if res, err := db.CartHasProduct(userId, productId); err != nil {
        return err
    } else if !res {
        return errors.New("Product not in cart")
    } 

    return db.DeleteProductFromCart(userId, productId, productInfo.Price)
}

func GetProductInfo(db ClientDb, productId uuid.UUID) (Product, error) {
    return db.GetProductInfo(productId)
}

type User struct {
    UserName string
}

type Cart struct {
    Products []Product
    TotalPrice float32
}

type Order struct {
    OrderId uuid.UUID
    Status property.Status
    DeliveryAdress string
    Products []Product
}

type Product struct {
    ProductId uuid.UUID
    ProductQuantity uint
}
