package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"notify_bot/logger"
)

var collection *mongo.Collection

type DataFromDb struct {
	Token string
	ChatID int64
}

func InitDB(){
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		logger.Fatal("")
	}
	err = client.Connect(context.TODO())
	if err != nil {
		logger.Fatal("")
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logger.Fatal("")
	}
	collection = client.Database("default").Collection("notify")
}

func GetChatID(token string) int64{
	var result DataFromDb
	filter := bson.D{{ "token" , token }}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		logger.Fatal("")
	}
	return result.ChatID
}

func AddNewChat(ChatID int64, token string){
	newChat := DataFromDb{token, ChatID}
	insertResult, err := collection.InsertOne(context.TODO(), newChat)
	if err != nil {
		logger.Fatal("")
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func IsChatIDInDB(ChatID int64) bool{
	var result DataFromDb
	filter := bson.D{{ "chatid" , ChatID }}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments{
			return false
		}
	}
	return true
}

func IsTokenInDB(token string) bool {
	var result DataFromDb
	filter := bson.D{{ "token" , token }}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments{
			return false
		}
	}
	return true
}

func GetToken(ChatID int64) string{
	var result DataFromDb
	filter := bson.D{{ "chatid" , ChatID }}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
	}
	return result.Token
}
