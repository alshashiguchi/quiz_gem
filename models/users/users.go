package users

import (
	database "alshashiguchi/quiz_gem/db/mysql"
	"alshashiguchi/quiz_gem/graph/model"
	"strconv"
)

//User - Struct User
type User struct {
	ID        string           `json:"id,omitempty"`
	Name      string           `json:"name,omitempty"`
	Email     string           `json:"email,omitempty"`
	Username  string           `json:"username,omitempty"`
	Access    model.Access     `json:"access,omitempty"`
	Situation model.UserStatus `json:"situation,omitempty"`
	Password  string           `json:"password,omitempty"`
}

//Create - create new user
func (user *User) Create() (model.User, error) {
	statement, err := database.Db.Prepare("INSERT INTO Users(Username, Name, Email, Access, Situation, Password) VALUES(?,?,?,?,?,?)")

	if err != nil {
		return model.User{}, err
	}
	// hashedPassword, err := HashPassword(user.Password)
	result, err := statement.Exec(user.Username, user.Name, user.Email, user.Access, user.Situation, user.Password)
	if err != nil {
		return model.User{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return model.User{}, err
	}

	// fmt.Println(reflect.TypeOf().Kind())

	return user.convertToGraphModelUser(id), nil
}

func (user *User) convertToGraphModelUser(id int64) model.User {
	var userModel model.User

	userModel.ID = strconv.Itoa(int(id))
	userModel.Name = user.Name
	userModel.Username = user.Username
	userModel.Email = user.Email
	userModel.Access = user.Access

	return userModel
}
