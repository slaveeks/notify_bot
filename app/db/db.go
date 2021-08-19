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

func InitDB(url string){
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		logger.Fatal(fmt.Sprint(err))
	}
	err = client.Connect(context.TODO())
	if err != nil {
		logger.Fatal(fmt.Sprint(err))
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logger.Fatal(fmt.Sprint(err))
	}
	collection = client.Database("default").Collection("notify")
	logger.Info("Connected to MongoDB!")
}

func GetChatID(token string) int64{
	var result DataFromDb
	filter := bson.D{{ "token" , token }}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		logger.Warn(fmt.Sprint(err))
	}
	return result.ChatID
}

func AddNewChat(ChatID int64, token string){
	newChat := DataFromDb{token, ChatID}
	insertResult, err := collection.InsertOne(context.TODO(), newChat)
	if err != nil {
		logger.Warn(fmt.Sprint(err))
	}
	logger.Info(fmt.Sprint("Inserted a single document: ", insertResult.InsertedID))
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
		logger.Warn(fmt.Sprint(err))
	}
	return result.Token
}
