package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var ctx context.Context

func main() {
	ctx = context.Background()

	log.Println("Starting server...")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("You must set your 'PORT' environment variable.")
	}

	mongodb_uri := os.Getenv("MONGODB_URI")
	if mongodb_uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	log.Println("Connecting to MongoDB using URI: " + mongodb_uri)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		panic(err)
	}

	mongodb_dbname := os.Getenv("MONGODB_DBNAME")
	if mongodb_dbname == "" {
		log.Fatal("You must set your 'MONGODB_DBNAME' environmental variable.")
	}
	db = client.Database(mongodb_dbname)

	server := &http.Server{
		Addr: ":" + port,
	}
	cors_config := cors.New(cors.Options{
		AllowedOrigins: []string{
			"https://cordbleibaum.com",
			"https://www.cordbleibaum.com",
			"https://cordbleibaum.de",
			"https://www.cordbleibaum.de",
			"http://localhost:3000",
		},
		AllowCredentials: true,
		Debug:            false,
	})

	handler := cors_config.Handler(http.HandlerFunc(serve))
	server.Handler = handler

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}

		log.Println("Disconnecting from MongoDB")
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	log.Println("Shutting down server")

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
