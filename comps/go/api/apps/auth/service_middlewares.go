package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"regexp"
	"strings"
)

type jwtClaim struct {
	UserId string `json:"userId"`
	OrgId  string `json:"orgId"`
	jwt.StandardClaims
}

// AddPublicRouter add public routers
func AddPublicRouter(method, route string) {
	publicRouters[method] += " " + route
}

// AddAdminRouter add admin router
func AddAdminRouter(method, route string) {
	adminRouters[method] += " " + route
}

// verifyRouter check if the path macthes with the router passed by parameter
func verifyRouter(router map[string]string, method, path string) bool {
	path = strings.ReplaceAll(path, "/api", "")
	matched := strings.Contains(router[method], path)

	if matched == false {
		routers := strings.Split(router[method], ",")
		for _, value := range routers {
			if value == "" {
				continue
			}
			compile, err := regexp.Compile(value)
			if err != nil {
				continue
			}
			matched = compile.MatchString(path)
			if matched {
				return matched
			}
		}
	}
	if matched == false {
		routers := strings.Split(router[method], ",")
		for _, value := range routers {
			if value != "" && strings.Contains(path, value) {
				return true
			}
		}
	}
	return matched
}

// isPublicRouter check if a route is public or not
func isPublicRouter(method, path string) bool {
	matched := verifyRouter(publicRouters, method, path)
	return matched
}

// validateToken function to validate jwt token
func validateToken(token string) (err error) {
	claims := &jwtClaim{}
	if token != "" {
		tkn, errToken := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if errToken != nil {
			if err == jwt.ErrSignatureInvalid {
				err = fmt.Errorf("invalid Signature")
				return
			}
			err = fmt.Errorf("invalid Token")
			return
		}

		if !tkn.Valid {
			err = fmt.Errorf("invalid Token")
		}

	} else {
		err = fmt.Errorf("token JWT was not set, please set the Token in the Headers")
	}
	return
}
