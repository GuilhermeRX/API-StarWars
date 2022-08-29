package models

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

type PlanetId struct {
	ID int
}

type Planet struct {
	ID               int    `json:"id,omitempty" valid:"optional"`
	Name             string `json:"name" valid:"type(string)"`
	Climate          string `json:"climate" valid:"type(string)"`
	Terrain          string `json:"terrain" valid:"type(string)"`
	MovieAppearances int    `json:"movieAppearances" valid:"optional"`
}

func FindAll() []Planet {
	cur, err := Db().Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}

	defer cur.Close(context.Background())
	var planets []Planet
	for cur.Next(context.Background()) {

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

func FindByName(name string) ([]Planet, error) {
	filter := bson.D{{Key: "name", Value: bson.D{{Key: "$regex", Value: name}}}}
	cur, err := Db().Find(context.Background(), filter)

	if err != nil {
		return []Planet{}, err
	}

	defer cur.Close(context.Background())
	var planets []Planet
	for cur.Next(context.Background()) {

		err := cur.All(context.Background(), &planets)

		if err != nil {
			return []Planet{}, err
		}
	}
	return planets, nil
}

func Create(planet Planet) (Planet, error) {
	_, err := Db().InsertOne(context.Background(), planet)

	if err != nil {
		return planet, errors.New("An error occurred while inserting")
	}

	return planet, nil
}

func Delete(id int) (int64, error) {
	res, err := Db().DeleteOne(context.TODO(), bson.D{{Key: "id", Value: id}})
	if err != nil {
		return 0, errors.New("Error when deleting a planet")
	}
	return res.DeletedCount, nil
}
