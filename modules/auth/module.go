package auth

import (
	"context"
	"log"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

func Register(app *fiber.App) *auth.Client {
	log.Println("REGISTER MODULE: Auth")

	serviceAccountKeyFilePath, err := filepath.Abs("./service-account.json")
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	//Firebase admin SDK initialization
	firebaseApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}

	//Firebase Auth
	auth, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Panic(err)
		panic("Firebase load error")
	}

	return auth
}
