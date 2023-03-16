package database

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
	"log"
)

func InitRealtime() *db.Ref {
	conf := &firebase.Config{
		DatabaseURL: "https://forgrxxm-default-rtdb.firebaseio.com/",
	}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile("database/forgrxxm-firebase-adminsdk-52ku1-9e3fa61eb9.json")

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err := app.Database(context.Background())
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	// As an admin, the app has access to read and write all data, regradless of Security Rules
	ref := client.NewRef("products")

	return ref
}

func SetRealtime(path string, data *Product) {
	if err := InitRealtime().Child(path).Set(context.Background(), data); err != nil {
		log.Fatalln("Error setting value:", err)
	}

}
func GetRealtime() map[string]Product {
	var data map[string]Product
	if err := InitRealtime().Get(context.Background(), &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	return data
}
func DelRealtime(path string) {
	if err := InitRealtime().Child(path).Delete(context.Background()); err != nil {
		log.Fatalln("Error deleting value:", err)
		log.Fatalln("Error deleting value:", err)
	}
}
