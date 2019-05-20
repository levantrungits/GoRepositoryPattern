package main

import (
	"fmt"
	"go-repo-modules/driver"
	"go-repo-modules/repository/repoimpl"
	models "go-repo-modules/model"
)

// host, port, user, password, dbname
// gorm.Open("postgres", "host=0.0.0.0 port=5432 user=default password=secret dbname=gorm_db sslmode=disable")
const (
	host = "0.0.0.0"
	port = "5432"
	user = "default"
	password = "secret"
	dbname = "go_repo_db"
)

func main()  {
	db := driver.Connect(host, port, user, password, dbname)
	err := db.SQL.Ping()
	if err != nil {
		panic(err)
	}

	// init Repo
	userRepo := repoimpl.NewUserRepo(db.SQL)

	// init 2 users
	uhp := models.User{
		ID: 1,
		Name: "Ung Hoang Phuc",
		Gender: "Male",
		Email: "uhp@gmail.com",
	}
	dt := models.User{
		ID: 2,
		Name: "Dan Truong",
		Gender: "Male",
		Email: "dt@gmail.com",
	}

	// insert data
	userRepo.Insert(uhp)
	userRepo.Insert(dt)

	// get all users
	users, _ := userRepo.Select()

	for i:= range users {
		fmt.Println(users[i])
	}
}