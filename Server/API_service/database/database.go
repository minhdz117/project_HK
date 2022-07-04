package database

import (
	"context" //use import withs " char
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Username string
	Password string
	Role     string
}

func ConnectMongo() *mongo.Client {
	mongoURL := "mongodb+srv://minhdz117:ZsjxTwn9i6GqgBA0@minhdz117.p9omz.mongodb.net/?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetUserByUsername(username string) (User, error) {
	var user User
	client := ConnectMongo()
	col := client.Database("project_hk").Collection("users")
	err := col.FindOne(context.TODO(), bson.M{"Username": username}).Decode(&user)
	return user, err
}

func CreateUser(user User) (bool, error) {
	client := ConnectMongo()
	col := client.Database("project_hk").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := col.InsertOne(ctx, user)
	if err != nil {
		return false, err
	}
	return true, err
}
