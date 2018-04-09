package user

import (
	"g/x/web"
	"gopkg.in/mgo.v2/bson"
)

var ErrMismatchedHashAndPassword = "crypto/bcrypt: hashedPassword is not the hash of the given password"

func GetUsers() ([]*User, error) {
	var users []*User
	err := userTable.FindWhere(bson.M{
		"dtime": bson.M{
			"$ne": 0,
		},
	}, &users)
	return users, err
}

func Login(uname, pwd string) (*User, error) {
	var user *User
	var query = bson.M{"uname": uname}
	err := userTable.FindOne(query, &user)
	if err != nil {
		if err.Error() == "not found" {
			return nil, web.BadRequest("Sai tên đăng nhập hoặc mật khẩu")
		}
		return nil, err
	}
	if err := Password(pwd).ComparePassword(user.HashedPassword); err != nil {
		if err.Error() == ErrMismatchedHashAndPassword {
			userLog.Error(err)
			return nil, web.BadRequest("Sai tên đăng nhập hoặc mật khẩu")
		}
		return nil, err
	}
	return user, nil
}
