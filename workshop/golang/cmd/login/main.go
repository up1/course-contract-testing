package main

import (
	"demo/login_service"
	"log"
	"net/url"
	"time"
)

var token = time.Now().Format("2006-01-02T15:04")

func main() {
	u, _ := url.Parse("http://localhost:4000")
	client := &login_service.Gateway{
		BaseURL: u,
	}

	users, err := client.GetUser(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(users)
}
