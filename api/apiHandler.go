package api

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	myRouter.HandleFunc("/send-email", SendEmail).Methods("POST")
	myRouter.HandleFunc("/deploy-backend", HandleBackendWebhook).Methods("POST")
	myRouter.HandleFunc("/deploy-website", HandleWebsiteWebhook).Methods("POST")

	fmt.Println("Listening for APIs on: 5000")
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(header, methods, origins)(myRouter)))
}
