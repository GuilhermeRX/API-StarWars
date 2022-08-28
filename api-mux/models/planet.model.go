package models

import (
	"context"
	"errors"
	"fmt"

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

func FindByID(id int) (Planet, error) {
	filter := bson.D{{Key: "id", Value: id}}
	result := Planet{}
	err := Db().FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		return result, errors.New("ID does not exist")
	}
	return result, nil
}

func FindByName(name string) (Planet, error) {
	filter := bson.D{{Key: "name", Value: name}}

	result := Planet{}
	err := Db().FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		return result, errors.New("There is no planet with that name")
	}
	return result, nil
}

func Create(planet Planet) (Planet, error) {
	_, err := Db().InsertOne(context.Background(), planet)

	if err != nil {
		return planet, errors.New("An error occurred while inserting")
	}

	return planet, nil
}

func Delete(id int) {
	res, err := Db().DeleteOne(context.TODO(), bson.D{{Key: "id", Value: id}})
	if err != nil {
		panic(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
}
