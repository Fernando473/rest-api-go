package main

import (
	"context"
	"os"

	"github.com/Fernando473/api-rest-go/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/gomongo"))

	if err != nil {
		panic(err)
	}

	// Static files from client/dist in React JS
	app.Use(cors.New())
	app.Static("/", "./client/dist")

	app.Post("/users", func(c *fiber.Ctx) error {

		var user models.User

		c.BodyParser(&user)

		coll := client.Database("gomongo").Collection("users")
		result, err := coll.InsertOne(context.TODO(), bson.D{
			{Key: "name", Value: user.Name},
		})

		if err != nil {
			panic(err)
		}

		return c.JSON(&fiber.Map{
			"data": result,
		})
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		var users []models.User

		coll := client.Database("gomongo").Collection("users")
		results, error := coll.Find(context.TODO(), bson.M{})

		if error != nil {
			panic(error)
		}
		for results.Next(context.TODO()) {
			var user models.User
			results.Decode(&user)
			users = append(users, user)
		}

		return c.JSON(&fiber.Map{
			"data": users,
		})

	})

	// Server is working on http://localhost:3000
	app.Listen(":" + port)
}
