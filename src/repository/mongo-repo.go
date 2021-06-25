package repository

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repoMongo struct{}

type IMongoRepository interface {
	FindOne(colName string, doc interface{}) *mongo.SingleResult
	FindAll(colName string, doc interface{}) (*mongo.Cursor, error)
	FindList(colName string, doc interface{}) (*mongo.Cursor, error)
	Create(colName string, doc interface{}) (*mongo.InsertOneResult, error)
	CreateMany(colName string, doc []interface{}) (*mongo.InsertManyResult, error)
	UpdateById(colName string, doc interface{}) (*mongo.UpdateResult, error)
	UpdateOne(colName string, doc interface{}) (*mongo.UpdateResult, error)
	UpdateMany(colName string, doc interface{}) (*mongo.UpdateResult, error)
	DeleteDocument(colName string, doc interface{}) (*mongo.DeleteResult, error)

	Findtemp(colName string, doc interface{}) *mongo.SingleResult
}

func ObjIMongoRepository() IMongoRepository {
	return &repoMongo{}
}

func (*repoMongo) FindOne(colName string, doc interface{}) *mongo.SingleResult {
	database, client := connectDB()
	defer closeConnection(client)
	collection := database.Collection(colName)

	kmap, err := trimObject(doc)
	if err != nil {
		log.Println(err)
	}
	filter := formateFilter(kmap)
	result := collection.FindOne(context.TODO(), filter)

	return result
}

/**
* Never Use this function in production application
* curser.All will load whole collection in memory
* suppose we have millions record then this function is going to fetch all records
**/
func (*repoMongo) FindAll(colName string, doc interface{}) (*mongo.Cursor, error) {
	database, client := connectDB()
	defer closeConnection(client)
	collection := database.Collection(colName)
	kmap, err := trimObject(doc)
	if err != nil {
		log.Println(err)
	}
	filter := formateFilter(kmap)
	cursor, err := collection.Find(context.TODO(), filter)

	return cursor, err
}

/**
* This function will provide List of objects in form of cursor which needs to decode
* If filter is null here it will fetch whole database
**/
func (*repoMongo) FindList(colName string, doc interface{}) (*mongo.Cursor, error) {
	database, client := connectDB()
	defer closeConnection(client)
	collection := database.Collection(colName)
	findOptions := options.Find()
	kmap, err := trimObject(doc)
	if err != nil {
		log.Println(err)
	}
	filter := formateFilter(kmap)
	cursor, err := collection.Find(context.TODO(), filter, findOptions)

	return cursor, err
}

func (*repoMongo) Create(colName string, doc interface{}) (*mongo.InsertOneResult, error) {
	database, client := connectDB()
	defer closeConnection(client)
	collection := database.Collection(colName)
	result, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println(err)
	}
	return result, err
}
func (*repoMongo) CreateMany(colName string, data []interface{}) (*mongo.InsertManyResult, error) {
	database, client := connectDB()
	defer closeConnection(client)
	collection := database.Collection(colName)
	result, err := collection.InsertMany(context.TODO(), data)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

/* In this method update object will hve specific settings*/
func (*repoMongo) UpdateById(colName string, doc interface{}) (*mongo.UpdateResult, error) {
	database, client := connectDB()
	defer closeConnection(client)
	collection := database.Collection(colName)
	kmap, err := trimObject(doc)
	if err != nil {
		log.Println(err)
	}
	filter, update := formateUpdate(kmap)
	result, err := collection.UpdateByID(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

func (*repoMongo) UpdateOne(colName string, doc interface{}) (*mongo.UpdateResult, error) {
	database, client := connectDB()
	defer closeConnection(client)
	collection := database.Collection(colName)
	kmap, err := trimObject(doc)
	if err != nil {
		log.Println(err)
	}
	filter, update := formateUpdate(kmap)
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

func (*repoMongo) UpdateMany(colName string, doc interface{}) (*mongo.UpdateResult, error) {
	database, client := connectDB()
	defer closeConnection(client)
	collection := database.Collection(colName)
	kmap, err := trimObject(doc)
	if err != nil {
		log.Println(err)
	}
	filter, update := formateUpdate(kmap)
	result, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

func (*repoMongo) DeleteDocument(colName string, doc interface{}) (*mongo.DeleteResult, error) {
	database, client := connectDB()
	defer closeConnection(client)
	collection := database.Collection(colName)
	kmap, err := trimObject(doc)
	if err != nil {
		log.Println(err)
	}
	filter := formateFilter(kmap)
	deleteResult, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Deleted %v documents in the collection\n", deleteResult.DeletedCount)
	return deleteResult, err
}

/**
* This function will provide connection with the DB
**/
func connectDB() (*mongo.Database, *mongo.Client) {
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
	}
	database := client.Database("Users")
	fmt.Println("Connected to MongoDB!")
	return database, client
}
func closeConnection(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Connection to MongoDB instance is closed.")
}

func formateUpdate(kmap map[string]interface{}) (primitive.M, []primitive.M) {
	var filter primitive.M
	var update []primitive.M
	i := 0
	for key, value := range kmap {
		if i == 0 {
			zmap := map[string]interface{}{key: value}
			filter = formateFilter(zmap)
			i = 1
			continue
		}
		zmap := map[string]interface{}{key: value}
		var temp primitive.M
		keyRecursion(&temp, &zmap)
		final := bson.M{"$set": temp} //$set could be replaced by other methods
		update = append(update, final)
	}
	return filter, update
}

func formateFilter(kmap map[string]interface{}) primitive.M {
	var filter primitive.M
	keyRecursion(&filter, &kmap)
	log.Println("filter", filter)
	return filter
}

func keyRecursion(filter *primitive.M, kmap *map[string]interface{}, tempKey ...string) {
	for key, value := range *kmap {
		if reflect.TypeOf(value).Kind() != reflect.Map {
			if tempKey != nil {
				key = tempKey[0] + "." + key
			}
			*filter = bson.M{key: value}
		} else {
			tempKey := key
			tempMap := value.(map[string]interface{})
			keyRecursion(filter, &tempMap, tempKey)
		}
	}
}

func trimObject(doc interface{}) (map[string]interface{}, error) {
	var kmap map[string]interface{}
	data, err := bson.Marshal(doc)
	if err != nil {
		log.Println(err)
	}
	err = bson.Unmarshal(data, &kmap)
	return kmap, err
}

func (*repoMongo) Findtemp(colName string, doc interface{}) *mongo.SingleResult {
	database, client := connectDB()
	defer closeConnection(client)
	collection := database.Collection(colName)

	kmap, err := trimObject(doc)
	if err != nil {
		log.Println(err)
	}
	filter := formateFilter(kmap)
	log.Println("filter", filter)
	result := collection.FindOne(context.TODO(), filter)

	return result
}
