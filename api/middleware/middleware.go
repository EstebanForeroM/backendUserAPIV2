package middleware

import (
	"log"
	"net/http"
)

func CorsPolicy(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("Applying CORS policy")
        w.Header().Add("Access-Control-Allow-Origin", "*") // Allow all domains
        w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Add("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
        w.Header().Add("Access-Control-Expose-Headers", "Authorization")
        w.Header().Add("Access-Control-Allow-Credentials", "true")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}
