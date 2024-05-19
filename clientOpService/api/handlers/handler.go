package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	database "github.com/EstebanForeroM/backendUserAPIV2/clientOpService/dataBase"
	usecases "github.com/EstebanForeroM/backendUserAPIV2/clientOpService/useCases"
)

type UserData struct {
    FirstName      string  `json:"first_name"`
    EmailAddresses []Email `json:"email_addresses"` 
    ID             string  `json:"id"`
}

type Email struct {
    EmailAddress string `json:"email_address"`
}

type UserOpHandler struct {
    DataBase database.DataBase
}

func parseUserData(r *http.Request) (UserData, error) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Printf("Error reading body: %v", err)
        return UserData{}, fmt.Errorf("can't read body")
    }
    defer r.Body.Close()

    var data struct {
        Data UserData `json:"data"`
    }

    if err := json.Unmarshal(body, &data); err != nil {
        log.Printf("Error parsing JSON body: %v", err)

        return UserData{}, fmt.Errorf("error parsing JSON body")
    }

    log.Printf("The user data is this %+v", data.Data)

    return data.Data, nil
}

func getUser(r *http.Request) (user usecases.User, err error) {
    userData, err := parseUserData(r)

    if err != nil {
        return user, err
    }

    user.Name = userData.FirstName
    user.Email = userData.EmailAddresses[0].EmailAddress
    user.Id = userData.ID

    return user, nil
}

