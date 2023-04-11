package mongo

import (
	"context"
	"fmt"
	"serverMonitor/pkg/constant"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NewMongoCollection mongo.Collection
type NewMongoClient mongo.Client

type NewMongo struct {
	database   string
	Collection string
}

var NewClient *mongo.Client

func init() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(constant.MongoUrl))
	if err != nil {
		fmt.Printf("failed to init mongo client, err:%+v\n", err)
		panic("error mongo client")
	}
	NewClient = client
}

func NewMgo(collection string) *NewMongo {
	return &NewMongo{
		database:   constant.MongoDB,
		Collection: collection,
	}
}

func (c *NewMongo) FindOne(filter interface{}) *mongo.SingleResult {
	coll := NewClient.Database(c.database).Collection(c.Collection)
	return coll.FindOne(context.Background(), filter)
}

func (c *NewMongo) UpdateOne(filter interface{}, value interface{}) error {
	coll := NewClient.Database(c.database).Collection(c.Collection)
	updateInfo := bson.M{"$set": value}
	_, err := coll.UpdateOne(context.Background(), filter, updateInfo)
	return err
}

func (c *NewMongo) InsertOne(value interface{}) error {
	coll := NewClient.Database(c.database).Collection(c.Collection)
	_, err := coll.InsertOne(context.Background(), value)
	return err
}

func (c *NewMongo) DeleteOne(filter interface{}) error {
	coll := NewClient.Database(c.database).Collection(c.Collection)
	_, err := coll.DeleteOne(context.Background(), filter, nil)
	return err
}

func (c *NewMongo) List() *mongo.Cursor {
	coll := NewClient.Database(c.database).Collection(c.Collection)
	result, _ := coll.Find(context.Background(), bson.M{})
	return result
}
