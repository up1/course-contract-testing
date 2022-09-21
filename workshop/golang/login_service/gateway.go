package login_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var (
	ErrNotFound    = errors.New("not found")
	ErrUnavailable = errors.New("api unavailable")
)

type User struct {
	User string `json:"user" pact:"example=somkiat"`
	Name string `json:"name" pact:"example=Somkiat Pui"`
	Role string `json:"role" pact:"example=admin,regex=^(admin|user|guest)$"`
}

type Gateway struct {
	BaseURL *url.URL
}

func (c *Gateway) GetUser(id int) (*User, error) {
	requestURL := fmt.Sprintf("%v/users/%v", c.BaseURL, id)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	var user User
	if res != nil {
		switch res.StatusCode {
		case http.StatusNotFound:
			return nil, ErrNotFound
		}
	}

	if err != nil {
		return nil, ErrUnavailable
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &user)

	return &user, err

}
