package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	routes "github.com/drawflow_app/controllers"
	"github.com/drawflow_app/database"
)

func main() {
	port := "9000"

	fmt.Println("The server is running on port", port)

	conn, err := grpc.Dial("localhost:9080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	database.DgraphClient = *database.NewDatabaseClient(conn)
	defer conn.Close()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.URLFormat)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	c := cors.New(cors.Options{

		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(c.Handler)

	r.Mount("/v1/programs", routes.DiagramsResource{}.Routes())

	log.Fatal(http.ListenAndServe(":"+port, r))
}
