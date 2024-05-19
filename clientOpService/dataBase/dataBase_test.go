package database

import (
	"testing"

	. "github.com/EstebanForeroM/backendUserAPIV2/clientOpService/useCases"
	"github.com/EstebanForeroM/backendUserAPIV2/db"
)

func TestCreateUser(t *testing.T) {
    dataBase := DataBase { clientDb: db.NewTestEntConnection(t) }

    user := User {
        Id: "user_abc",
        Name: "abc",
        Email: "abc@gmail.com",
    } 

    err := dataBase.CreateUser(user)

    if err != nil {
        t.Error("Error creating user: ", err)
    }

    err = dataBase.CreateUser(user);

    if err == nil {
        t.Error("Should get error when adding a user with the same id: ", err)
    }

    user.Id = "user_abcd"
    user.Email = "abcd@gmail.com"

    err = dataBase.CreateUser(user)

    if err != nil {
        t.Error("Error creating user: ", err)
    }
} 

func TestDeleteUser(t *testing.T) {
    dataBase := DataBase { clientDb: db.NewTestEntConnection(t) }

    err := dataBase.DeleteUser("user_abc")

    if err == nil {
        t.Error("There should be an error when trying to delete a non existing user")
    }

    user := User {
        Id: "user_abc",
        Name: "abc",
        Email: "abc@gmail.com",
    } 

    dataBase.CreateUser(user)

    err = dataBase.DeleteUser(user.Id)

    if err != nil {
        t.Error("Error at deleting existing user: ", err)
    }

    err = dataBase.CreateUser(user)

    if err != nil {
        t.Error("Error at creating user when user was already deleted: ", err)
    }

    user.Id = "user_abcd"
    user.Email = "abcd@gmail.com"

    dataBase.CreateUser(user)
    
    err = dataBase.DeleteUser(user.Id)

    if err != nil {
        t.Error("There shouldn't be n error when deleting a existing user")
    }
}

