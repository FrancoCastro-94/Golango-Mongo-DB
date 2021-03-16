package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*Player captura del Body, el mensaje que nos llega */
type Player struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	LastName string             `bson:"lastName"`
	Number   string             `bson:"number"`
	Foto     string             `bson:"foto"`
	UUID     string             `bson:"uuid,omitempty"`
}
