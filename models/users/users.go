package users

import (
	database "alshashiguchi/quiz_gem/db/mysql"
	"alshashiguchi/quiz_gem/graph/model"
	"strconv"

	log "github.com/sirupsen/logrus"
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

//GetAll - Get all users permission ADMIN
func GetAll() ([]model.User, error) {
	stmt, err := database.Db.Prepare("SELECT ID, Username, Name, Email, Access, Situation FROM Users")

	if err != nil {
		logWarningUser("Database prepare", "Error Get All User", err)
		return []model.User{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		logWarningUser("Commmand error sql", "Error Get All User", err)
		return []model.User{}, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.Email, &user.Access, &user.Situation)
		if err != nil {
			logWarningUser("Error scan user", "Error Get All User", err)
			return []model.User{}, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		logWarningUser("Rows Error", "Error Get All User", err)
		return []model.User{}, err
	}
	return users, nil
}

//Create - create new user
func (user *User) Create() (model.User, error) {
	statement, err := database.Db.Prepare("INSERT INTO Users(Username, Name, Email, Access, Situation, Password) VALUES(?,?,?,?,?,?)")

	if err != nil {
		logWarningUser("Database prepare", "Error Create User", err)
		return model.User{}, err
	}
	// hashedPassword, err := HashPassword(user.Password)
	result, err := statement.Exec(user.Username, user.Name, user.Email, user.Access, user.Situation, user.Password)
	if err != nil {
		logWarningUser("Fields create user", "Error Create User", err)
		return model.User{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return model.User{}, err
	}

	// fmt.Println(reflect.TypeOf().Kind())

	return user.convertToGraphModelUser(id), nil
}

func logWarningUser(title, msg string, err error) {
	log.WithField(title, log.Fields{
		"package":  "users",
		"function": "model.User.Create",
		"error":    err,
	}).Warning(msg)
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
