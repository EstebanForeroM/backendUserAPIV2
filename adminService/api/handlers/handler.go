package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	database "github.com/EstebanForeroM/backendUserAPIV2/adminService/dataBase"
	usecases "github.com/EstebanForeroM/backendUserAPIV2/adminService/useCases"
	"github.com/EstebanForeroM/backendUserAPIV2/property"
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/google/uuid"
)

type AdminHandler struct {
    DataBase database.DataBase
}

func (handler *AdminHandler) UpdateOrderState(w http.ResponseWriter, r *http.Request) {
    claims, _ := clerk.SessionClaimsFromContext(r.Context())

    status := r.PathValue("status")
    orderId := r.PathValue("orderId")

    orderIdUUID, err := uuid.Parse(orderId)

    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    userId := claims.Subject

    statusEnum, err := property.StringToStatus(status)

    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    err = usecases.UpdateOrderStatus(&handler.DataBase, orderIdUUID, statusEnum, userId)

    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
}

func (handler *AdminHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
    claims, _ := clerk.SessionClaimsFromContext(r.Context())

    userId := claims.Subject

    orders, err := usecases.GetOrders(&handler.DataBase, userId)

    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(orders)
}

func (handler *AdminHandler) GetOrdersByUserId(w http.ResponseWriter, r *http.Request) {
    userId := r.PathValue("userId")

    claims, _ := clerk.SessionClaimsFromContext(r.Context())

    orders, err := usecases.GetOrderByUserId(&handler.DataBase, claims.Subject, userId)

    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(orders)
}

func (handler *AdminHandler) GetOrderByIdAndStatus(w http.ResponseWriter, r *http.Request) {
    status := r.PathValue("status")
    userId := r.PathValue("userId")

    claims, _ := clerk.SessionClaimsFromContext(r.Context())

    statusEnum, err := property.StringToStatus(status)

    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    orders, err := usecases.GetOrderByUserId(&handler.DataBase, claims.Subject, userId)

    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    var userOrders []usecases.Order

    userOrders, err = usecases.FilterOrdersByStatus(orders, statusEnum)

    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(userOrders)
}

func (handler *AdminHandler) GetOrdersByStatus(w http.ResponseWriter, r *http.Request) {
    status := r.PathValue("status")

    claims, _ := clerk.SessionClaimsFromContext(r.Context())

    statusEnum, err := property.StringToStatus(status)

    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    orders, err := usecases.GetOrderByStatus(&handler.DataBase, claims.Subject, statusEnum)

    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(orders)
}
