package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	usecases "github.com/EstebanForeroM/backendUserAPIV2/clientOpService/useCases"
)

type deleteData struct {
    Data struct {
        Id string `json:"id"`
    } `json:"data"`
}

func (h *UserOpHandler) DeleteUser(w http.ResponseWriter, r *http.Request) { 

    userId, err := getIdFronRequest(r)

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    err = usecases.DeleteUser(&h.DataBase, userId)

    if err != nil {
        log.Printf("Error deleting user with id %v: %v", userId, err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    log.Println("user deleted succesfully")
}

func getIdFronRequest(r *http.Request) (string, error) {
    body, err := io.ReadAll(r.Body)

    if err != nil {
        log.Println("Error reading json data: ", err)
        return "", err
    }
    defer r.Body.Close()

    var delData deleteData

    err = json.Unmarshal(body, &delData)

    if err != nil {
        log.Println("Error unmarshaling data: ", err)
        return "", err
    }

    return delData.Data.Id, nil
}
