package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mxssl/todo/api/handlers"
	"github.com/mxssl/todo/api/restapi"
	"github.com/mxssl/todo/api/restapi/operations"
	"github.com/mxssl/todo/db"
	"github.com/mxssl/todo/store"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {

	// Load config from .env
	if err := godotenv.Load(); err != nil {
		log.Println("Connot load .env file, env variables will be used")
	}

	// Assign env variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	dbConnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	d, err := db.NewDB(dbConnString)
	if err != nil {
		log.Fatalln(err)
	}

	itemStore := store.NewItemStore(d)

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTodoAPI(swaggerSpec)
	api.ServerShutdown = func() {
		d.Close()
		log.Println("Database connection pool closed")
	}
	server := restapi.NewServer(api)
	defer server.Shutdown()
	handlers.Init(itemStore, api)

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "todo app"
	parser.LongDescription = "My example of a todo app"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
