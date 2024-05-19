package productapi

import (
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
    if err := godotenv.Load("./../.env"); err != nil {
        log.Fatalln("Error loading enviroment variables")
    }   

    m.Run()
}

func TestGetProductInfo(t *testing.T) {
    UUIDProductId := uuid.MustParse("9c0b2f03-945d-493e-b785-c425d948b1bd") 
    prductInfo, err := GetProductInfo(UUIDProductId)

    if err != nil {
        t.Errorf("Error fetching product info: %s", err)
    }

    if prductInfo.Name != "Hamburguer" {    
        t.Errorf("Error fetching product info: %s", prductInfo.Name)
    }

    if prductInfo.Price != 99.7 {
        t.Errorf("Error fetching product info: %v", prductInfo.Price)
    }
}
