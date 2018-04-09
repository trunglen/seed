package user

import (
	"g/x/web"
	"gopkg.in/mgo.v2/bson"
	"seed/x/logger"
	"seed/x/mongodb"
	"seed/x/validator"
)

var userTable = mongodb.NewTable("user", "usr")
var userLog = logger.NewLogger("user")

const (
	GUEST = "guest"
	ADMIN = "admin"
)

type User struct {
	mongodb.Model  `bson:",inline"`
	Name           string   `bson:"name" json:"name"`
	UName          string   `bson:"uname" json:"uname"`
	HashedPassword string   `bson:"password" json:"-"`
	Password       Password `bson:"-" json:"password"`
	Email          string   `bson:"email" json:"email"`
	Role           string   `bson:"role" json:"role"`
}

const (
	errExists           = "user exists!"
	errMisMatchUNamePwd = "username or password is incorect!"
)

func (u *User) Create() error {
	var err = validator.Struct(u)
	if u.Role == GUEST {
		if user, _ := GetGuestByEmail(u.Email); user != nil {
			return web.BadRequest(errExists)
		}
	} else {
		if user, _ := GetAdmin(u.UName, "admin"); user != nil {
			return web.BadRequest(errExists)
		}
	}

	hashed, _ := u.Password.Hash()
	u.HashedPassword = hashed
	if err != nil {
		userLog.Error(err)
		return web.WrapBadRequest(err, "")
	}
	return userTable.Create(u)
}

func GetAdmin(uname string, role string) (*User, error) {
	var user *User
	var err = userTable.FindOne(bson.M{
		"uname": uname,
		"role":  role,
	}, &user)
	return user, err
}

func GetGuestByEmail(email string) (*User, error) {
	var user *User
	var err = userTable.FindOne(bson.M{
		"email": email,
		"role":  GUEST,
	}, &user)
	return user, err
}

func DeleteUserByID(id string) error {
	return userTable.DeleteByID(id)
}
