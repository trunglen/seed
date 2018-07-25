package auth

import (
	"gopkg.in/mgo.v2/bson"
	"rocky-springs-86767/g/x/math"
	"rocky-springs-86767/x/mongodb"
)

var authTable = mongodb.NewTable("auth", "auth")

type Auth struct {
	mongodb.Model `bson:",inline"`
	UserID        string `bson:"user_id" json:"user_id"`
	Role          string `bson:"role" json:"role"`
	Revoked       bool   `bson:"revoked" json:"revoked"`
}

func Create(userID, role string) *Auth {
	var auth = Auth{
		UserID:  userID,
		Role:    role,
		Revoked: false,
	}
	auth.SetID(math.RandString("auth", 80))
	authTable.Upsert(bson.M{
		"user_id": userID,
		"role":    role,
	}, auth)
	return &auth
}

func GetByID(id string) (*Auth, error) {
	var auth *Auth
	return auth, authTable.FindID(id, &auth)
}
