package productapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func GetProductInfo(productId uuid.UUID) (ProductInfo, error) {
    baseUrl := os.Getenv("PRODUCT_API_URL")
    url := baseUrl + "/api/dish/" + productId.String()
    response, err := http.Get(url)

    if response.StatusCode != http.StatusOK {
        return ProductInfo{}, fmt.Errorf("Product not found: %v", response)
    }

    if err != nil {
        return ProductInfo{}, err
    }

    defer response.Body.Close()

    var productInfo ProductInfo

    err = json.NewDecoder(response.Body).Decode(&productInfo)

    if err != nil {
        return ProductInfo{}, err
    }

    return productInfo, nil
}

type ProductInfo struct {
    Name string `json:"name"`
    Price float32 `json:"price"`
}
