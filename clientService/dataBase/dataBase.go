package database

import (
	"context"

	usecases "github.com/EstebanForeroM/backendUserAPIV2/clientService/useCases"
	"github.com/EstebanForeroM/backendUserAPIV2/db"
	"github.com/EstebanForeroM/backendUserAPIV2/ent"
	"github.com/EstebanForeroM/backendUserAPIV2/ent/product"
	"github.com/EstebanForeroM/backendUserAPIV2/ent/user"
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

func (d *DataBase) GetProductInfo(productId uuid.UUID) (usecases.Product, error) {
    ctx := context.Background()

    productEnt, err := d.clientDb.Product.Query().Where(product.Pid(productId)).Only(ctx)

    if err != nil {
        return usecases.Product{}, err
    }

    return usecases.Product{
        ProductId: productEnt.Pid,
        ProductQuantity: uint(productEnt.Quantity),
    }, nil
}

func (d *DataBase) AddCartToUser(userId string) (error) {
    ctx := context.Background()

    user, err := d.clientDb.User.Query().Where(user.ID(userId)).Only(ctx)

    if err != nil {
        return err
    }

    cart, err := d.clientDb.Cart.Create().Save(ctx)

    if err != nil {
        return err
    }

    _, err = user.Update().SetCart(cart).Save(ctx)

    return err
}

func (d *DataBase) CartHasProduct(userId string, productId uuid.UUID) (bool, error) {
    ctx := context.Background()

    return d.clientDb.User.Query().Where(user.ID(userId)).
        QueryCart().QueryProducts().Where(product.Pid(productId)).Exist(ctx)
}

func (d *DataBase) UserHasCart(userId string) (bool, error) {
    ctx := context.Background()

    return d.clientDb.User.Query().Where(user.ID(userId)).
        QueryCart().Exist(ctx)
}
    

func (d *DataBase) AddNewProductToCart(userId string, productId uuid.UUID, price float32) (error) {
    ctx := context.Background()

    cart, err := d.clientDb.User.Query().Where(user.ID(userId)).QueryCart().Only(ctx)

    if err != nil {
        return err
    }

    product, err := d.clientDb.Product.Create().
        SetPid(productId).Save(ctx)

    if err != nil {
        return err
    }

    _, err = cart.Update().AddProducts(product).SetTotalPrice(cart.TotalPrice + price).Save(ctx)

    return err 
}

func (d *DataBase) DeleteProductFromCart(userId string, productId uuid.UUID, price float32) (error) {
    ctx := context.Background()

    cart, err := d.clientDb.User.Query().Where(user.ID(userId)).QueryCart().Only(ctx)

    if err != nil {
        return err
    }

    product, err := cart.QueryProducts().Where(product.Pid(productId)).Only(ctx)

    if err != nil {
        return err 
    }

    if product.Quantity == 1 {
        err = d.clientDb.Product.DeleteOne(product).Exec(ctx)

        if err != nil { return err }

        err = cart.Update().SetTotalPrice(0).Exec(ctx)

        if err != nil { return err }
    } else {
        _, err = product.Update().SetQuantity(product.Quantity - 1).Save(ctx)
        
        if err != nil { return err }
    }

    cart.Update().SetTotalPrice(cart.TotalPrice - price).Save(ctx)

    return nil
}

func (d *DataBase) AddProductToCart(userId string, productId uuid.UUID, price float32) (error) {
    ctx := context.Background()

    cart, err := d.clientDb.User.Query().Where(user.ID(userId)).QueryCart().Only(ctx)

    if err != nil {
        return err 
    }

    product, err := cart.QueryProducts().Where(product.Pid(productId)).Only(ctx)

    if err != nil {
        return err
    }

    _, err = cart.Update().SetTotalPrice(cart.TotalPrice + price).Save(ctx)

    if err != nil {
        return err
    }

    _, err = product.Update().SetQuantity(product.Quantity + 1).Save(ctx)

    return err
}

func (d *DataBase) GetCart(userId string) (usecases.Cart, error) {
    ctx := context.Background()

    cart, err := d.clientDb.User.Query().Where(user.ID(userId)).QueryCart().Only(ctx)

    if err != nil {
        return usecases.Cart{}, err
    }

    productsEnt, err := cart.QueryProducts().All(ctx)

    if err != nil {
        return usecases.Cart{}, err
    }

    var products []usecases.Product

    for _, product := range productsEnt {
        products = append(products, usecases.Product{
            ProductId: product.Pid,
            ProductQuantity: uint(product.Quantity),
        })
    }

    return usecases.Cart{
        Products: products,
        TotalPrice: cart.TotalPrice,
    }, nil
}

func (d *DataBase) GetUser(userId string) (usecases.User, error) {
    ctx := context.Background()

    userEnt, err := d.clientDb.User.Query().Where(user.ID(userId)).Only(ctx)

    if err != nil {
        return usecases.User{}, err
    }

    return usecases.User{
        UserName: userEnt.Name,
    }, nil
}

func (d *DataBase) AddOrder(userId string, deliveryAddress string) (error) {
    ctx := context.Background()

    user, err := d.clientDb.User.Query().Where(user.ID(userId)).Only(ctx)

    if err != nil {
        return err
    }

    products, err := user.QueryCart().QueryProducts().All(ctx)

    if err != nil {
        return err
    }

    order, err := d.clientDb.Order.Create().
        AddProducts(products...).SetDeliveryAdress(deliveryAddress).Save(ctx)

    if err != nil {
        return err
    }

    _, err = user.Update().AddOrders(order).Save(ctx)

    return err
}

func (d *DataBase) CartIsEmpty(userId string) (bool, error) {
    ctx := context.Background()

    cart, err := d.clientDb.User.Query().Where(user.ID(userId)).QueryCart().Only(ctx)

    if err != nil {
        return false, err
    }

    return cart.TotalPrice == 0, nil
}

func (d *DataBase) GetOrders(userId string) ([]usecases.Order, error) {
    ctx := context.Background()

    user, err := d.clientDb.User.Query().Where(user.ID(userId)).Only(ctx)

    if err != nil {
        return nil, err
    }

    ordersEnt, err := user.QueryOrders().All(ctx)

    if err != nil {
        return nil, err
    }

    var orders []usecases.Order

    for _, orderEnt := range ordersEnt {
        order, err := getUseCaseOrder(orderEnt)
        if err != nil { return nil, err }
        orders = append(orders, order)
    }

    return orders, nil
}

func getUseCaseOrder(orderEnt *ent.Order) (usecases.Order, error) {
    var products []usecases.Product

    productsEnt, err := orderEnt.QueryProducts().All(context.Background())

    if err != nil {
        return usecases.Order{}, err
    }

    for _, productEnt := range productsEnt {
        products = append(products, usecases.Product{
            ProductId: productEnt.Pid,
            ProductQuantity: uint(productEnt.Quantity),
        })
    }

    return usecases.Order{
        OrderId: orderEnt.ID,
        Status: orderEnt.Status,
        Products: products,
        DeliveryAdress: orderEnt.DeliveryAdress,
    }, nil
}

func (d *DataBase) DeleteCart(userId string) (error) {
    ctx := context.Background()

    user, err := d.clientDb.User.Query().Where(user.ID(userId)).Only(ctx)

    if err != nil {
        return err
    }

    cart, err := user.QueryCart().Only(ctx)

    if err != nil {
        return err
    }

    d.clientDb.Cart.DeleteOne(cart).Exec(ctx)

    return err
}
