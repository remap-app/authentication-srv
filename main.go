package main

import (
	"fmt"
	"net/http"
	"context"
	"log"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func initializeAppWithServiceAccount() *firebase.App {
	opt := option.WithCredentialsFile("secret/service_account_key_staging.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app
}

func main() {
	app := initializeAppWithServiceAccount()
	fmt.Print(app)
	if err != nil {
		panic(err)
	}
	opt := option.WithCredentialsFile("secret/service_account_key_staging.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	
	auth, err := app.Auth(context.Background())
	if err != nil {
			panic(err)
	}

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		t, err := auth.VerifyIDToken()
		if err != nil {
			panic(err)
		}
	})
	
	http.ListenAndServe(":8080", nil)
}
