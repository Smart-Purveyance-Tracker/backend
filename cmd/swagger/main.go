package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Smart-Purveyance-Tracker/backend/api/rest-swagger/apiimpl"
	"github.com/Smart-Purveyance-Tracker/backend/api/rest-swagger/restapi"
	"github.com/Smart-Purveyance-Tracker/backend/api/rest-swagger/restapi/operations"
	"github.com/Smart-Purveyance-Tracker/backend/repository"
	"github.com/Smart-Purveyance-Tracker/backend/service"
	"github.com/Smart-Purveyance-Tracker/backend/service/auth"
)

type envs struct {
	MongoURI  string `envconfig:"MONGO_URI" required:"true"`
	JWTSecret string `envconfig:"JWT_SECRET" required:"true"`
}

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewSwaggerAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "test"
	parser.LongDescription = "test"
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

	var e envs
	err = envconfig.Process("", &e)
	if err != nil {
		log.Fatal(err.Error())
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(e.MongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	userRepo := repository.NewUserMongoDB(client)
	userSvc := service.NewUserImpl(userRepo)
	jwtSvc := auth.NewJWTService([]byte(e.JWTSecret))
	server.SetHandler(apiimpl.ConfigureAPI(api, apiimpl.NewServer(userSvc, jwtSvc, service.NewProductImpl(repository.NewProductMongoDB(client)))))

	if err := server.Serve(); err != nil {
		log.Panicln(err)
	}
}
