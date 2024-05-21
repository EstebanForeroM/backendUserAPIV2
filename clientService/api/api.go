package api

import (
	"net/http"
	"os"

	"github.com/EstebanForeroM/backendUserAPIV2/clientService/api/handlers"
	database "github.com/EstebanForeroM/backendUserAPIV2/clientService/dataBase"
	"github.com/clerk/clerk-sdk-go/v2"
    clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
)

func GetMux() *http.ServeMux {
    clerkKey := os.Getenv("CLERK_KEY")
    clerk.SetKey(clerkKey)

    mux := http.NewServeMux()

    dataBase := database.NewDataBase()

    handler := handlers.ClientHandler { DataBase: dataBase }

    addProductToCartHandler := http.HandlerFunc(handler.AddProductToCart)
    deleteProductFromCartHandler := http.HandlerFunc(handler.DeleteProductFromCart)
    getCartWithUserIdHandler := http.HandlerFunc(handler.GetCartWithUserId)
    createOrderHandler := http.HandlerFunc(handler.CreateOrder)
    getOrdersHandler := http.HandlerFunc(handler.GetOrders)
    getProductInfoHandler := http.HandlerFunc(handler.GetProductInfo)

    mux.Handle(
        "POST /cart/products/{pid}",
        clerkhttp.RequireHeaderAuthorization()(addProductToCartHandler),
    )

    mux.Handle(
        "DELETE /cart/products/{pid}",
        clerkhttp.RequireHeaderAuthorization()(deleteProductFromCartHandler),
    )

    mux.Handle(
        "GET /cart",
        clerkhttp.RequireHeaderAuthorization()(getCartWithUserIdHandler),
    ) 

    mux.Handle(
        "POST /orders/{deliveryAdress}",
        clerkhttp.RequireHeaderAuthorization()(createOrderHandler),
    )

    mux.Handle(
        "GET /orders",
        clerkhttp.RequireHeaderAuthorization()(getOrdersHandler),
    )

    mux.Handle(
        "GET /products/{pid}",
        clerkhttp.RequireHeaderAuthorization()(getProductInfoHandler),
    )
 
    return mux
}
