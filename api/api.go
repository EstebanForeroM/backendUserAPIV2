package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EstebanForeroM/backendUserAPIV2/api/middleware"
)

type Server struct {
    port string
    httpServer *http.Server
}

func NewServer(port string) *Server {
    return &Server{
        port: port,
    }
}

func (a *Server)CloseServer() {
    a.httpServer.Close()
}

func (a *Server)InitializeServer(routes... RouteData) error {
    mux := initializeMux(routes...) 
    a.createServer(mux)
    return a.startServer()
}

func (a *Server) createServer(mux *http.ServeMux) {
    a.httpServer = &http.Server{
        Addr: ":" + a.port,
        Handler: middleware.CorsPolicy(mux),
    }
}

func (a *Server) startServer() error {
    log.Println("Starting server in port: ", a.port)
    if err := a.httpServer.ListenAndServe(); err != nil {
        log.Println("Error starting the server: ", err)
        return err
    } 

    return nil
}

func initializeMux(routes... RouteData) *http.ServeMux {
    mux := http.NewServeMux()

    for _, route := range routes {
        mux.Handle("/" + route.Prefix + "/", http.StripPrefix("/" + route.Prefix, route.Mux))
    }

    return mux
}

type RouteData struct {
    Prefix string
    Mux *http.ServeMux
}

func TestServeMux() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(`server working`))
    })

    return mux
}

func AddPrefix(mux *http.ServeMux, prefix string) *http.ServeMux {
    prefix = fmt.Sprintf("/%s/", prefix)
    userOpMux := http.NewServeMux()
    userOpMux.Handle(prefix, http.StripPrefix(prefix, mux))
    fmt.Println("user op mux is: ", userOpMux)
    return userOpMux
}
