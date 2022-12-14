package db

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var Db *firestore.Client
var Bu *storage.BucketHandle

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file detected trying to access os variabels...")
	}

	if len(os.Getenv("PROJECT_ID")) == 0 {
		log.Fatal("PROJECT_ID not set.")
	}

	ctx := context.Background()
	var app *firebase.App
	if os.Getenv("GIN_MODE") == "release" {
		conf := &firebase.Config{
			ProjectID:     os.Getenv("PROJECT_ID"),
			StorageBucket: os.Getenv("PROJECT_ID") + ".appspot.com",
		}
		app, err = firebase.NewApp(ctx, conf)
	} else {
		opt := option.WithCredentialsFile("secrets/firebase.json")
		conf := &firebase.Config{
			ProjectID:     os.Getenv("PROJECT_ID"),
			StorageBucket: os.Getenv("PROJECT_ID") + ".appspot.com",
		}
		app, err = firebase.NewApp(context.Background(), conf, opt)

	}

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	Db = client

	s, err := app.Storage(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	b, err := s.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}
	Bu = b
}
