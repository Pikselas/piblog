package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogData struct {
	Tag     string
	Content string
}

type BlogPost struct {
	Id       string
	Title    string
	Contents []BlogData
}

type BlogDescription struct {
	Id          string
	Title       string
	Description string
	Tags        []string
}

const db_name = "piblogdata"

var server_api_options = options.ServerAPI(options.ServerAPIVersion1)

var db_client *mongo.Client

func init_connection() {
	db_url := ENV["DB_URL"]
	opts := options.Client().ApplyURI(db_url).SetServerAPIOptions(server_api_options)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	db_client = client
}

func GetDbCollection(name string) (*mongo.Collection, error) {
	collection := db_client.Database(db_name).Collection(name)
	return collection, nil
}

func AddTagIfNotExists(tag string) {
	collection, err := GetDbCollection("tags")
	if err != nil {
		panic(err)
	}
	res := collection.FindOne(context.Background(), bson.M{"TAG_NAME": tag})
	if res.Err() == mongo.ErrNoDocuments {
		_, err = collection.InsertOne(context.Background(), bson.M{"TAG_NAME": tag})
		if err != nil {
			panic(err)
		}
	}
}

func GetTags(tag_name string) []string {
	collection, err := GetDbCollection("tags")
	if err != nil {
		panic(err)
	}
	cursor, err := collection.Find(context.Background(), bson.M{"TAG_NAME": bson.M{"$regex": tag_name, "$options": "i"}})
	if err != nil {
		panic(err)
	}
	var tag_result []struct{ TAG_NAME string }
	if err = cursor.All(context.Background(), &tag_result); err != nil {
		panic(err)
	}
	var result []string
	for _, tag := range tag_result {
		result = append(result, tag.TAG_NAME)
	}
	return result
}

func SearchBlog(bson_format bson.M) []BlogDescription {
	collection, err := GetDbCollection("blogs_desc")
	if err != nil {
		panic(err)
	}
	cursor, err := collection.Find(context.Background(), bson_format)
	if err != nil {
		panic(err)
	}
	var result []BlogDescription
	if err = cursor.All(context.Background(), &result); err != nil {
		panic(err)
	}
	return result
}

func SearchBlogByTags(tags []string) []BlogDescription {
	return SearchBlog(bson.M{"tags": bson.M{"$in": tags}})
}

func SearchBlogByTitle(title string) []BlogDescription {
	return SearchBlog(bson.M{"title": bson.M{"$regex": title, "$options": "i"}})
}

func InsertBlog(blog BlogPost, desc BlogDescription) {
	collection, _ := GetDbCollection("blogs")
	_, err := collection.InsertOne(context.Background(), blog)
	if err != nil {
		panic(err)
	}
	for _, tag := range desc.Tags {
		AddTagIfNotExists(tag)
	}
	collection, _ = GetDbCollection("blogs_desc")
	_, err = collection.InsertOne(context.Background(), desc)
	if err != nil {
		panic(err)
	}
}

func FetchBlog(ID string) BlogPost {
	collection, _ := GetDbCollection("blogs")
	var result BlogPost
	err := collection.FindOne(context.Background(), bson.M{"id": ID}).Decode(&result)
	if err != nil {
		panic(err)
	}
	return result
}
