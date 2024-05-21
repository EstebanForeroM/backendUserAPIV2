package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	database "github.com/EstebanForeroM/backendUserAPIV2/clientService/dataBase"
	usecases "github.com/EstebanForeroM/backendUserAPIV2/clientService/useCases"
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/google/uuid"
)

type ClientHandler struct {
    DataBase database.DataBase
}

func (handler *ClientHandler) GetProductInfo(w http.ResponseWriter, r *http.Request) {
    idString := r.PathValue("pid")

    claims, _ := clerk.SessionClaimsFromContext(r.Context())

    userId := claims.Subject


    puuid, err := uuid.Parse(idString)

    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    product, err := usecases.GetProductInfo(&handler.DataBase, userId, puuid)

    //if err != nil {
        //return
        //w.WriteHeader(http.StatusInternalServerError)
        //w.Write([]byte("Error getting product info"))
        //log.Println(err, " ", puuid)
    //}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(product)
}

func (handler *ClientHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
    claims, _ := clerk.SessionClaimsFromContext(r.Context())

    userId := claims.Subject

    orders, err := usecases.GetOrders(&handler.DataBase, userId)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Error getting orders"))
        log.Println(err, " ", userId)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(orders)
}

func (handler *ClientHandler) AddProductToCart(w http.ResponseWriter, r *http.Request) {
    idString := r.PathValue("pid") 
    puuid, err := uuid.Parse(idString)

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Invalid product id"))
        log.Println(err)
        return
    }

    claims, _ := clerk.SessionClaimsFromContext(r.Context())

    userId := claims.Subject 

    if err := usecases.AddProductToCart(&handler.DataBase, userId, puuid); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Error adding product to cart"))
        log.Println(err, " ", userId)
        return
    }
}

func (handler *ClientHandler) DeleteProductFromCart(w http.ResponseWriter, r *http.Request) {
    idString := r.PathValue("pid")
    puuid, err := uuid.Parse(idString)

    claims, _ := clerk.SessionClaimsFromContext(r.Context())

    userId := claims.Subject

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Invalid product id"))
        return
    }

    if err := usecases.DeleteProductFromCart(&handler.DataBase, userId, puuid); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Error deleting product from cart"))
        log.Println(err, " ", userId)
        return
    }
}

func (handler *ClientHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
    claims, _ := clerk.SessionClaimsFromContext(r.Context())
    deliveryAdress := r.PathValue("deliveryAdress")

    userId := claims.Subject

    log.Println("Creating order for user: ", userId, " with delivery adress: ", deliveryAdress)

    if err := usecases.CreateOrder(&handler.DataBase, userId, deliveryAdress); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Error creating order"))
        log.Println(err, " ", userId)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Order created"))
}

func (handler *ClientHandler) GetCartWithUserId(w http.ResponseWriter, r *http.Request) {
    claims, _ := clerk.SessionClaimsFromContext(r.Context())

    userId := claims.Subject

    cart, err := usecases.GetCart(&handler.DataBase, userId)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Error getting cart"))
        log.Println(err, " ", userId)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(cart)
}
