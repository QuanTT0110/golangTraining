package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"quanlyhoso/config"
	"time"
)

var db *mongo.Database

func Connect() {
	env := config.GetEnv()

	cl, err := mongo.NewClient(options.Client().ApplyURI(env.Database.URI))
	if err != nil {
		log.Println(err)
		log.Fatal("Cannot connect to database:", env.Database.URI)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = cl.Connect(ctx)
	if err != nil {
		log.Println(err)
	}
	db = cl.Database(env.Database.Name)
	fmt.Println("Database Connected to", env.Database.Name)
}
