package db

import (
	"context"
	"log"
	"time"

	"github.com/FrancoCastro-94/Go-Mongo-practice/models"
	"go.mongodb.org/mongo-driver/bson"
)

var db = MongoCN.Database("team")
var collection = db.Collection("players")

//GetAllPlayers return all players
func GetAllPlayers() []models.Player {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var allPlayersList []models.Player

	if err := cursor.All(ctx, &allPlayersList); err != nil {
		log.Fatal(err)
	}
	return allPlayersList
}

//InsertPlayer save the player in BD
func InsertPlayer(player models.Player) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	registro := bson.M{
		"uuid":     player.UUID,
		"name":     player.Name,
		"lastName": player.LastName,
		"number":   player.Number,
		"foto":     player.Foto,
	}

	_, err := collection.InsertOne(ctx, registro)
	if err != nil {
		return false, err
	}

	return true, nil
}

//Update player
func UpdatePlayer(p models.Player) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	update := bson.M{
		"uuid":     p.UUID,
		"name":     p.Name,
		"lastName": p.LastName,
		"number":   p.Number,
		"foto":     p.Foto,
	}

	playerUdate := bson.M{
		"$set": update,
	}

	id := p.UUID

	_, err := collection.UpdateOne(ctx, bson.M{"uuid": id}, playerUdate)
	if err != nil {
		return false, err
	}
	return true, nil

}

//Delete a player
func DeletePlayer(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"uuid": id})
	if err != nil {
		return false, err
	}

	return true, nil

}

//GetOnePlayer get a player by id
func GetOnePlayer(id string) models.Player {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"uuid": id}

	var player models.Player

	err := collection.FindOne(ctx, filter).Decode(&player)
	if err != nil {
		log.Fatal(err)
	}

	return player
}
