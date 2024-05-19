package database

import (
	"context"

	. "github.com/EstebanForeroM/backendUserAPIV2/clientOpService/useCases"
	"github.com/EstebanForeroM/backendUserAPIV2/db"
	"github.com/EstebanForeroM/backendUserAPIV2/ent"
)

type DataBase struct {
    clientDb *ent.Client
}

func NewDataBase() DataBase {
    return DataBase{
        clientDb: db.NewEntConnection(),
    }
}

func (d *DataBase) CreateUser(user User) error {
    ctx := context.Background()

    err := d.clientDb.User.Create().
        SetID(user.Id).
        SetName(user.Name).
        SetEmail(user.Email).
        Exec(ctx)

    return err
}

func (d *DataBase) DeleteUser(userId string) error {
    ctx := context.Background()

    err := d.clientDb.User.DeleteOneID(userId).Exec(ctx)

    return err
}

