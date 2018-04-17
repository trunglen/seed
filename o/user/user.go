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
	mongodb.Model     `bson:",inline"`
	Name              string   `bson:"name" json:"name"`
	Phone             string   `bson:"phone" json:"phone"`
	HashedPassword    string   `bson:"password" json:"-"`
	Password          Password `bson:"-" json:"password"`
	RestaurantName    string   `bson:"restaurant_name" json:"restaurant_name"`
	RestaurantID      string   `bson:"restaurant_id" json:"restaurant_id"`
	RestaurantAddress string   `bson:"restaurant_address" json:"restaurant_address"`
	Role              string   `bson:"role" json:"role"`
}

const (
	errExists           = "user exists!"
	errMisMatchUNamePwd = "username or password is incorect!"
)

func (u *User) CreateAccount() error {
	if user, _ := GetUserByPhone(u.Phone); user != nil {
		return web.BadRequest(errExists)
	}
	return userTable.Create(u)
}

func GetUserByPhone(phone string) (*User, error) {
	var user *User
	var err = userTable.FindOne(bson.M{
		"phone": phone,
	}, &user)
	return user, err
}

func (u *User) Create() error {
	var err = validator.Struct(u)
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
