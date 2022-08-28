package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type PlanetId struct {
	ID int
}

type Planet struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name"`
	Clima   string `json:"clima"`
	Terreno string `json:"terreno"`
}

func FindAll() []Planet {
	cur, err := Db().Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}

	defer cur.Close(context.Background())
	var planets []Planet
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()

		err := cur.All(context.Background(), &planets)

		if err != nil {
			panic(err)
		}
	}
	return planets
}

func Count() int {
	return len(FindAll())
}

func FindByID(id int) Planet {
	filter := bson.D{{Key: "id", Value: id}}
	result := Planet{}
	err := Db().FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func FindByName(name string) []Planet {
	filter := bson.D{{Key: "name", Value: name}}

	cur, err := Db().Find(context.Background(), filter)
	if err != nil {
		panic(err)
	}

	defer cur.Close(context.Background())
	var planets []Planet
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()

		err := cur.All(context.Background(), &planets)

		if err != nil {
			panic(err)
		}
	}
	return planets
}

func Create(planet Planet) Planet {
	_, err := Db().InsertOne(context.Background(), planet)
	if err != nil {
		panic(err)
	}

	return planet
}

// func Update() {

// }

// func Delete() {
// 	return
// }
