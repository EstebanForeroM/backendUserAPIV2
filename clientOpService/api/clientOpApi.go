package api

import (
	"net/http"

	"github.com/EstebanForeroM/backendUserAPIV2/clientOpService/api/handlers"
	database "github.com/EstebanForeroM/backendUserAPIV2/clientOpService/dataBase"
)

const prefix string = "clientOp"

func GetMux() *http.ServeMux {
    mux := http.NewServeMux()

    dataBase := database.NewDataBase()

    handler := handlers.UserOpHandler { DataBase: dataBase }

    mux.HandleFunc("POST /createUser", handler.NewUser)
    mux.HandleFunc("POST /deleteUser", handler.DeleteUser)
 
    return mux
}


