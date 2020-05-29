package users

import (
	database "alshashiguchi/quiz_gem/db/mysql"
	"alshashiguchi/quiz_gem/graph/model"
	"log"
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
func (user *User) Create() (int64, error) {
	statement, err := database.Db.Prepare("INSERT INTO Users(Username, Name, Email, Access, Situation, Password) VALUES(?,?,?,?,?,?)")

	if err != nil {
		log.Fatal(err)
	}
	// hashedPassword, err := HashPassword(user.Password)
	result, err := statement.Exec(user.Username, user.Name, user.Email, user.Access, user.Situation, user.Password)
	if err != nil {
		log.Fatal(err)
	}

	lastId, _ := result.LastInsertId()

	// fmt.Println(reflect.TypeOf().Kind())

	return lastId, nil
}

//ConvertToGraphModelUser - convert graphql model user to user
func (user *User) ConvertToGraphModelUser() model.User {
	var userModel model.User

	userModel.Name = user.Name
	userModel.Username = user.Username
	userModel.Email = user.Email
	userModel.Access = user.Access

	return userModel
}
