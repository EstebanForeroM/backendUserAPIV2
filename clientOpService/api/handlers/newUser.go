package handlers

import (
	"log"
	"net/http"

	usecases "github.com/EstebanForeroM/backendUserAPIV2/clientOpService/useCases"
)

func (h *UserOpHandler) NewUser(w http.ResponseWriter, r *http.Request) {

    user, err := getUser(r)

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        log.Println("Error creating new user: ", err)
        return
    }

    err = usecases.NewUser(&h.DataBase, user)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        log.Println("Error creating user: ", err)
        return
    }

    log.Println("user created succesfully")
}


