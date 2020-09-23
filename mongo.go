package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	name  string
	id    string
	price int
}

func initMongo() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	fmt.Println("Start connect Mongodb")
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Start to check Mongodb connection")
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	// 指定获取要操作的数据集
	collection := client.Database("ratel").Collection("product")
	// p1 := Product{"xiaomi", "id003"}
	// p2 := Product{"huawei", "id004"}
	p1 := bson.D{
		{"name", "mx2"},
		{"id", "id003"},
	}
	insertResult, err := collection.InsertOne(context.TODO(), p1)
	if err != nil {
		log.Fatal(err)
	}
	// insertResult, err = collection.InsertOne(context.TODO(), p2)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	filter := bson.D{{"id", "id003"}}

	update := bson.D{
		{"$inc", bson.D{
			{"price", 1},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	// 创建一个Student变量用来接收查询的结果
	var result Product
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
}

func main() {
	initMongo()
}
