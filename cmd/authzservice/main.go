package main

import (
	"authz/internal/apis"
	"authz/openapi/gen/authzservice/server"
	"authz/openapi/gen/authzservice/server/operations"
	authorization "authz/openapi/gen/authzservice/server/operations/authorization"
	"flag"
	"log"

	"github.com/go-openapi/loads"
)

var portFlag = flag.Int("port", 8080, "Port to run this service on")

func main() {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(server.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewAuthzServiceAPI(swaggerSpec)
	s := server.NewServer(api)
	defer func() {
		_ = s.Shutdown()
	}()

	// parse flags
	flag.Parse()
	s.Port = *portFlag
	authzAPI := apis.AuthzAPI{}

	// api.GetAPIV1CompaniesHandler = operations.GetAPIV1CompaniesHandlerFunc(comAPI.List)
	// api.GetAPIV1CompaniesCompanyIDHandler = operations.GetAPIV1CompaniesCompanyIDHandlerFunc(comAPI.Get)
	// api.PostAPIV1CompaniesHandler = operations.PostAPIV1CompaniesHandlerFunc(comAPI.Create)
	api.AuthorizationAuthorizeHandler = authorization.AuthorizeHandlerFunc(authzAPI.Auth)
	// GetHealthHandler sets the operation handler for the get health operation
	api.GetHealthHandler = operations.GetHealthHandlerFunc(authzAPI.Health)


	// serve API
	if err := s.Serve(); err != nil {
		log.Fatalln(err)
	}
}