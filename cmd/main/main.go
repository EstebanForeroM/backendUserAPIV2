package main

import (
	"log"
	"os"

	"github.com/EstebanForeroM/backendUserAPIV2/api"
	"github.com/joho/godotenv"
    clientOpServiceApi "github.com/EstebanForeroM/backendUserAPIV2/clientOpService/api"
    clientServiceApi "github.com/EstebanForeroM/backendUserAPIV2/clientService/api"
    adminServiceApi "github.com/EstebanForeroM/backendUserAPIV2/adminService/api"
)

func main() {
    loadEnvVariables()
    setUpServer()
}

func setUpServer() {
    port := os.Getenv("PORT")
    server := api.NewServer(port) 

    server.InitializeServer(
        api.RouteData{Prefix: "clientOp", Mux: clientOpServiceApi.GetMux()},
        api.RouteData{Prefix: "clientService", Mux: clientServiceApi.GetMux()},
        api.RouteData{Prefix: "adminService", Mux: adminServiceApi.GetMux()},
        api.RouteData{Prefix: "test", Mux: api.TestServeMux()},
    )
}

func loadEnvVariables() {
    if err := godotenv.Load(); err != nil {
        log.Println("Error loading enviroment variables")
    }   
}
