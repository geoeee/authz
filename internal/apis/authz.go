package apis

import (

	"authz/openapi/gen/authzservice/server/operations"

	"authz/openapi/gen/authzservice/server/operations/authorization"
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	jwt "github.com/dgrijalva/jwt-go"

)
type AuthzAPI struct {

}


func (api *AuthzAPI) Health(params operations.GetHealthParams) middleware.Responder {


	fmt.Println("authz service health")
	return operations.NewGetHealthOK()
}


func (api *AuthzAPI) Auth(params authorization.AuthorizeParams) middleware.Responder {


	fmt.Println("authz service authorize")

	fmt.Println(params.Body.Token)


	cl, err := getPayloadFromToken(params.Body.Token)
	if err != nil {
		return authorization.NewAuthorizeInternalServerError()
	}
	fmt.Println(cl)

	roles := cl["user_roles"].([]interface{})

fmt.Println(roles)

for _, r := range roles {
	sr := fmt.Sprintf("%s", r) 
	if sr  == "admin" {
		return authorization.NewAuthorizeOK()
	}
}
	return authorization.NewAuthorizeUnauthorized()
}

func getPayloadFromToken(tokenStr string) (jwt.MapClaims, error) {
	parser := jwt.Parser{SkipClaimsValidation: true}
	token, _, err := parser.ParseUnverified(tokenStr, jwt.MapClaims{})
	return token.Claims.(jwt.MapClaims), err
}