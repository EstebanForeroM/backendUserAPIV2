package database

import (
	"testing"

	"github.com/EstebanForeroM/backendUserAPIV2/db"
	"github.com/google/uuid"
)

func TestGetProductInfo(t *testing.T) {
    d := DataBase{db.NewTestEntConnection(t)}

    db.AddTestUser("user_1", d.clientDb)

    d.AddCartToUser("user_1")


    productID := uuid.New()

    d.AddNewProductToCart("user_1", productID, 13)

    if productInfo, err := d.GetProductInfo(productID); err != nil {
        t.Error("Error getting product info: ", err)
        return
    } else if productInfo.ProductQuantity != 1 {
        t.Error("Error getting product info, invalid product quantity: ", productInfo)
        return
    }

    d.AddProductToCart("user_1", productID, 13)

    if productInfo, err := d.GetProductInfo(productID); err != nil {
        t.Error("Error getting product info: ", err)
        return
    } else if productInfo.ProductQuantity != 2 {
        t.Error("Error getting product info, invalid product quantity: ", productInfo)
        return
    }
}

func TestAddCartToUser(t *testing.T) {
    d := DataBase{db.NewTestEntConnection(t)} 
    
    db.AddTestUser("user_1", d.clientDb)

    err := d.AddCartToUser("user_1")

    if err != nil {
        t.Error("Error adding cart to user: ", err)
        return
    }

    if res, err := d.UserHasCart("user_1"); err != nil {
        t.Error("Error adding cart to user: ", err, res)
        return
    } else if !res {
        t.Error("Error adding cart to user: ", res)
        return
    }
}

func TestComplexAddCartToUser(t *testing.T) {
    d := DataBase{db.NewTestEntConnection(t)} 
    
    db.AddTestUser("user_1", d.clientDb)

    err := d.AddCartToUser("user_1")

    if err != nil {
        t.Error("Error adding cart to user: ", err)
        return
    }

    productIDs := []uuid.UUID{uuid.New(), uuid.New()}

    err = d.AddNewProductToCart("user_1", productIDs[0], 10)

    if err != nil {
        t.Error("Error adding new product to cart: ", err)
        return
    }

    err = d.AddProductToCart("user_1", productIDs[0], 10)

    if err != nil {
        t.Error("Error adding product to cart: ", err)
        return
    }

    err = d.AddNewProductToCart("user_1", productIDs[1], 10)

    if err != nil {
        t.Error("Error adding new product to cart: ", err)
        return
    }

    if res, err := d.CartHasProduct("user_1", productIDs[0]); err != nil {
        t.Error("Error adding new product to cart: ", err, res)
        return
    } else if !res {
        t.Error("Error adding new product to cart: ", res)
        return
    }

    err = d.AddOrder("user_1", "myhouse")

    if err != nil {
        t.Error("Error adding order: ", err)
        return
    }

    err = d.DeleteCart("user_1")

    if err != nil {
        t.Error("Error deleting cart: ", err)
        return
    }

    if res, err := d.UserHasCart("user_1"); err != nil {
        t.Error("Error adding order: ", err, res)
        return
    } else if res {
        t.Error("user shouldn't have cart: ", res)
        return
    }

    err = d.AddCartToUser("user_1")

    if err != nil {
        t.Error("Error adding cart to user: ", err)
        return
    }

    err = d.AddNewProductToCart("user_1", productIDs[0], 10)

    if err != nil {
        t.Error("Error adding new product to cart: ", err)
        return
    }

    err = d.AddProductToCart("user_1", productIDs[0], 10)

    if err != nil {
        t.Error("Error adding product to cart: ", err)
        return
    }

    err = d.AddNewProductToCart("user_1", productIDs[1], 10)

    if err != nil {
        t.Error("Error adding new product to cart: ", err)
        return
    }

    if res, err := d.CartHasProduct("user_1", productIDs[0]); err != nil {
        t.Error("Error adding new product to cart: ", err, res)
        return
    } else if !res {
        t.Error("Error adding new product to cart: ", res)
        return
    }

}

func TestAddNewProductToCart(t *testing.T) {
    d := DataBase{db.NewTestEntConnection(t)} 
    
    db.AddTestUser("user_1", d.clientDb)

    d.AddCartToUser("user_1") 

    productIDs := []uuid.UUID{uuid.New()}

    err := d.AddNewProductToCart("user_1", productIDs[0], 10)

    if err != nil {
        t.Error("Error adding new product to cart: ", err)
        return
    }

    if res, err := d.CartHasProduct("user_1", productIDs[0]); err != nil {
        t.Error("Error adding new product to cart: ", err, res)
        return
    } else if !res {
        t.Error("Error adding new product to cart: ", res)
        return
    }
}

func TestDeleteCart(t *testing.T) {
    d := DataBase{db.NewTestEntConnection(t)}

    db.AddTestUser("user_1", d.clientDb)

    d.AddCartToUser("user_1")

    err := d.DeleteCart("user_1")

    if err != nil {
        t.Error("Error deleting cart: ", err)
        return
    }

    if res, err := d.UserHasCart("user_1"); err != nil {
        t.Error("Error deleting cart: ", err, res)
        return
    } else if res {
        t.Error("Error deleting cart: ", res)
        return
    }
}

func TestAddProductToCart(t *testing.T) {
    d := DataBase{db.NewTestEntConnection(t)}

    db.AddTestUser("user_1", d.clientDb)

    d.AddCartToUser("user_1")

    productIDs := []uuid.UUID{uuid.New()}

    d.AddNewProductToCart("user_1", productIDs[0], 10)

    err := d.AddProductToCart("user_1", productIDs[0], 10)

    if err != nil {
        t.Error("Error adding product to cart: ", err)
        return
    }

    cart, _ := d.GetCart("user_1") 

    if cart.Products[0].ProductQuantity != 2 {
        t.Error("Error adding product to cart: ", cart.Products[0])
        return
    }
}

func TestGetCart(t *testing.T) {
    d := DataBase{db.NewTestEntConnection(t)}

    db.AddTestUser("user_1", d.clientDb)

    d.AddCartToUser("user_1")

    if _, err := d.GetCart("user_1"); err != nil {
        t.Error("Error getting cart: ", err)
        return
    }

    productUUID := uuid.New()

    d.AddNewProductToCart("user_1",productUUID, 10)

    cart, err := d.GetCart("user_1")

    if err != nil {
        t.Error("Error getting cart: ", err)
        return
    }

    if len(cart.Products) != 1 {
        t.Error("Error getting cart: ", cart.Products)
        return
    }

    if cart.Products[0].ProductId != productUUID {
        t.Error("Error getting cart: ", cart.Products)
        return
    }
}

func TestCartHasProduct(t *testing.T) {
    d := DataBase{db.NewTestEntConnection(t)}

    db.AddTestUser("user_1", d.clientDb)

    d.AddCartToUser("user_1")

    productIDs := []uuid.UUID{uuid.New()}

    d.AddNewProductToCart("user_1", productIDs[0], 10)

    if res, err := d.CartHasProduct("user_1", productIDs[0]); err != nil {

        t.Error("Error checking if cart has product: ", err, res)
        return
    } else if !res {
        t.Error("Error checking if cart has product: ", res)
        return
    }

    if res, err := d.CartHasProduct("user_1", uuid.New()); err != nil {

        t.Error("Error checking if cart has product: ", err, res)
        return
    } else if res {
        t.Error("Error checking if cart has product, cart shouldn't have product: ", res)
        return
    }
}

func TestAddOrder(t *testing.T) {
    d := DataBase{db.NewTestEntConnection(t)}

    db.AddTestUser("user_1", d.clientDb)

    d.AddCartToUser("user_1")

    productIDs := []uuid.UUID{uuid.New()}

    d.AddNewProductToCart("user_1", productIDs[0], 10)

    err := d.AddOrder("user_1", "myhouse")

    if err != nil {
        t.Error("Error adding order: ", err)
        return
    }

    if orders, err := d.GetOrders("user_1"); err != nil {
        t.Error("Error adding order: ", err)
        return
    } else if len(orders) != 1 {
        t.Error("There should be only one order: ", orders)
        return
    } else if orders[0].DeliveryAdress != "myhouse" {
        t.Error("Wrong delivery Adress: ", orders)
        return
    } else if len(orders[0].Products) != 1 {
        t.Error("Wrong product quantity: ", orders)
        return
    }

    d.DeleteCart("user_1")

    d.AddCartToUser("user_1")

    d.AddNewProductToCart("user_1", productIDs[0], 10)

    err = d.AddOrder("user_1", "myhouse2")

    if err != nil {
        t.Error("Error adding second order: ", err)
        return
    }
}

func TestGetOrders(t *testing.T) {
    d := DataBase{db.NewTestEntConnection(t)}

    db.AddTestUser("user_1", d.clientDb)

    d.AddCartToUser("user_1")

    productIDs := []uuid.UUID{uuid.New(), uuid.New(), uuid.New()}

    d.AddNewProductToCart("user_1", productIDs[0], 10)

    d.AddNewProductToCart("user_1", productIDs[1], 10)

    d.AddOrder("user_1", "myhouse")

    d.DeleteCart("user_1")

    d.AddCartToUser("user_1")

    d.AddNewProductToCart("user_1", productIDs[2], 10)

    d.AddProductToCart("user_1", productIDs[2], 10)

    d.AddOrder("user_1", "myhouse2")

    if orders, err := d.GetOrders("user_1"); err != nil {
        t.Error("Error getting orders: ", err)
        return
    } else if len(orders) != 2 {
        t.Error("Error getting all the orders: ", orders)
        return
    } else if len(orders[0].Products) != 2 {
        t.Errorf(`Error getting products in first order: %v
Expected %v | %v`, orders[0].Products, productIDs[0], productIDs[1])
        return
    } else if len(orders[1].Products) != 1 {
        t.Error("Error getting products in second order: ", orders[1].Products)
        return
    } else if orders[0].DeliveryAdress != "myhouse" {
        t.Error("Error getting delivery address in first order: ", orders[0].DeliveryAdress)
        return
    } else if orders[1].DeliveryAdress != "myhouse2" {
        t.Error("Error getting delivery address in second order: ", orders[1].DeliveryAdress)
        return
    } 
}
