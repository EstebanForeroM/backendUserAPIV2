package db

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/EstebanForeroM/backendUserAPIV2/ent"
	"github.com/EstebanForeroM/backendUserAPIV2/ent/enttest"

    _ "github.com/mattn/go-sqlite3"
    _ "github.com/lib/pq"
)

const (
    host     = "viaduct.proxy.rlwy.net"
    port     = 38039
    userDB   = "postgres"
    password = "ZdOYymQqhVUcWYkOAGcrMpfVghtTLkaJ"
    dbname   = "railway"
)

type DataBase struct {
    Client *ent.Client
}

func NewEntConnection() *ent.Client {
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, userDB, password, dbname)

    client, err := ent.Open("postgres", psqlconn)

    if err != nil {
        log.Fatalf("failed opening connection to postgres: %v", err)
    }
    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

    return client
}

func NewTestEntConnection(t *testing.T) *ent.Client {
    return enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
}
