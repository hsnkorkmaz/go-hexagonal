package middleware

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
)

type azureMiddleware struct {
	keySetUrl     string
	accessControl IAccessControl
}

func NewAzureMiddleware(keysetUrl string, ac IAccessControl) *azureMiddleware {
	return &azureMiddleware{
		keySetUrl:     keysetUrl,
		accessControl: ac,
	}
}

func (az *azureMiddleware) AzureMiddleWare(next http.HandlerFunc, reqRoles []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := verifyToken(r, az.keySetUrl)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "token is not valid")
			return
		}

		claims, err := getClaims(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "could not get claims")
			return
		}
		
		roles := claims["roles"].([]interface{})
		isPermitted := az.accessControl.RBAC(roles, reqRoles)
		if !isPermitted {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "user does not have the required roles")
			return
		}
		next(w, r)
	}
}

func extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header not found")
	}
	values := strings.Split(authHeader, " ")
	if len(values) != 2 || values[0] != "Bearer" {
		return "", errors.New("invalid authorization header")
	}
	return values[1], nil
}

func verifyToken(r *http.Request, keySetUrl string) (*jwt.Token, error) {
	accessToken, err := extractToken(r)
	if err != nil {
		return nil, err
	}

	keyset, err := jwk.Fetch(r.Context(), keySetUrl)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {

		if token.Method.Alg() != jwa.RS256.String() {
			return nil, errors.New("signing method is not valid")
		}

		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("kid not found")
		}

		keys, isOk := keyset.LookupKeyID(kid)
		if !isOk {
			return nil, errors.New("key not found")
		}

		publickey := &rsa.PublicKey{}
		err = keys.Raw(publickey)
		if err != nil {
			return nil, errors.New("could not parse public key")
		}

		return publickey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func getClaims(token *jwt.Token) (map[string]interface{}, error) {
	claims, isOk := token.Claims.(jwt.MapClaims)
	if !isOk {
		return nil, errors.New("claims not found")
	}
	return claims, nil
}
