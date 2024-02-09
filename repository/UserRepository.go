package repository

import (
	entity "Clinic_System/entity/User"
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// make the interface that will be implemented by the struct
type UserRepository interface {
	GetAll() ([]entity.User, error)
	GetAllDoctors() ([]entity.User, error)
	SignUp(user entity.User) (entity.User, error)
	LogIn(email string, password string) (entity.User, error)
}

// make the struct that implements the interface
type userRepository struct {
	db []entity.User
}

// New creates a new user repository object
func New() UserRepository {
	// call the FillDB function to fill the db with the data in the database
	return &userRepository{}
}

// implement the methods of the interface
func (r userRepository) GetAll() ([]entity.User, error) {

	//call the FillDB function to fill the db with the data in the database
	r.db, _ = FillDB() // return the db
	return r.db, nil

}

// function that gets all doctors stored in the mysql db
func (r userRepository) GetAllDoctors() ([]entity.User, error) {
	// call the FillDB function to fill the db with the data in the database
	r.db, _ = FillDB()
	// make a slice of users
	var doctors []entity.User
	// loop through the results
	for _, user := range r.db {
		if user.Type == "Doctor" {
			doctors = append(doctors, user)
		}
	}
	// return the slice of doctors
	return doctors, nil
}

// singup function with email that is the ID and password and name
func (r userRepository) SignUp(user entity.User) (entity.User, error) {
	// call the FillDB function to fill the db with the data in the database
	r.db, _ = FillDB() // check if the email already exists
	for _, user1 := range r.db {
		if user1.Email == user.Email {
			return entity.User{}, errors.New("email already exists")
		}
	}
	// call the NewSignUp function to write the new user to the database
	NewSignUp(user)

	return user, nil
}

func (r userRepository) LogIn(email string, password string) (entity.User, error) {
	// call the FillDB function to fill the db with the data in the database
	r.db, _ = FillDB()
	// check if the email and password are correct
	for _, user := range r.db {
		if user.Email == email && user.Password == password {
			return user, nil
		}
	}
	return entity.User{}, errors.New("Invalid Credentials")
}

// open a connection to the mysql database and fill the db with the data in the database
func FillDB() ([]entity.User, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("root:root@tcp(%s:%s)/clinic_system", dbHost, dbPort)
	database, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	// make the query
	results, err := database.Query("SELECT email, password, type, fullName FROM user")
	if err != nil {
		panic(err.Error())
	}
	// make a slice of users
	var users []entity.User
	// loop through the results
	for results.Next() {
		var user entity.User
		// for each row, scan the result into our user composite object
		err = results.Scan(&user.Email, &user.Password, &user.Type, &user.FullName)
		if err != nil {
			panic(err.Error())
		}
		// and then print out the user's Name attribute
		users = append(users, user)
	}
	// close the database
	defer database.Close()
	// return the slice of users
	return users, nil
}

// write the New Sign Up function to the database
func NewSignUp(user entity.User) (entity.User, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("root:root@tcp(%s:%s)/clinic_system", dbHost, dbPort)
	database, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	// make the query
	insert, err := database.Query("INSERT INTO user VALUES (?,?,?,?)", user.Email, user.Password, user.Type, user.FullName)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	defer database.Close()
	return user, nil
}
