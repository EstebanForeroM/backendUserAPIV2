package api

import (
	"net/http"
	"os"

	"github.com/EstebanForeroM/backendUserAPIV2/adminService/api/handlers"
	database "github.com/EstebanForeroM/backendUserAPIV2/adminService/dataBase"
	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
)

func GetMux() *http.ServeMux {
    clerkKey := os.Getenv("CLERK_KEY")
    clerk.SetKey(clerkKey)

    mux := http.NewServeMux()

    dataBase := database.NewDataBase()

    handler := handlers.AdminHandler { DataBase: dataBase }

    getOrdersHandler := http.HandlerFunc(handler.GetOrders)
    getOrdersByUserIdHandler := http.HandlerFunc(handler.GetOrdersByUserId)
    getOrdersByStatusHandler := http.HandlerFunc(handler.GetOrdersByStatus)
    updateOrderStatusHandler := http.HandlerFunc(handler.UpdateOrderState)
    getOrderByIdAndStatusHandler := http.HandlerFunc(handler.GetOrderByIdAndStatus)

    mux.Handle(
        "GET /orders",
        clerkhttp.RequireHeaderAuthorization()(getOrdersHandler),
    )

    mux.Handle(
        "GET /orders/{userId}/status/{status}",
        clerkhttp.RequireHeaderAuthorization()(getOrderByIdAndStatusHandler),
    )

    mux.Handle(
        "GET /orders/{userId}",
        clerkhttp.RequireHeaderAuthorization()(getOrdersByUserIdHandler),
    )

    mux.Handle(
        "GET /orders/status/{status}",
        clerkhttp.RequireHeaderAuthorization()(getOrdersByStatusHandler),
    )

    mux.Handle(
        "PUT /orders/{orderId}/status/{status}",
        clerkhttp.RequireHeaderAuthorization()(updateOrderStatusHandler),
    )
 
    return mux
}
