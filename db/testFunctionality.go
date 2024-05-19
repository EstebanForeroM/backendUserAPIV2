package db

import (
	"context"

	"github.com/EstebanForeroM/backendUserAPIV2/ent"
)

func AddTestUser(userId string, d *ent.Client) {
    d.User.Create().SetID(userId).
        SetName("Esteban").
        SetEmail("estebanmff@gmail.com").
        SaveX(context.Background())
}
