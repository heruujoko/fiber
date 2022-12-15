package auth

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	apperror "fiberapp/modules/errors"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

func GetFirebaseAuth() (error, *auth.Client) {
	serviceAccountKeyFilePath, err := filepath.Abs("./service-account.json")
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
		return err, nil
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	//Firebase admin SDK initialization
	firebaseApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err, nil
	}

	//Firebase Auth
	auth, err := firebaseApp.Auth(context.Background())
	if err != nil {
		return err, nil
	}

	return nil, auth
}

func AuthMiddlware(c *fiber.Ctx) error {
	// Set a custom header on all responses:
	c.Set("X-Custom-Header", "Hello, World")
	headers := c.GetReqHeaders()
	authorizationToken := headers["Authorization"]
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
	if len(idToken) == 0 {
		errorCode, errorContext := apperror.TokenRequiredError()
		c.SendStatus(errorCode)
		return c.JSON(errorContext)
	}

	err, firebaseAuth := GetFirebaseAuth()
	if err != nil {
		errorCode, errorContext := apperror.FirebaseLibraryError()
		c.SendStatus(errorCode)
		return c.JSON(errorContext)
	}

	//verify token
	verifiedToken, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		errorCode, errorContext := apperror.ForbiddenError()
		c.SendStatus(errorCode)
		return c.JSON(errorContext)
	}
	// Go to next middleware:
	fmt.Println(verifiedToken)
	return c.Next()
}
