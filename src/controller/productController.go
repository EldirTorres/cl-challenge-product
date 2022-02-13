package controller

import (
	"cl.falabella.product/src/model"
	configRedis "cl.falabella.product/src/service/redis"
	"encoding/json"
	"fmt"
	"net/http"
)

func Saveproduct(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var productData model.Product

	err := decoder.Decode(&productData)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	redis, _ := configRedis.RedisInit()
	data, err := json.Marshal(productData)

	if err != nil {
		fmt.Println(err)
	}

	key := fmt.Sprintf("id[%d]", productData.Sku)
	result := redis.SetKey(key, data, 1000)

	if result == true {
		fmt.Println("Saved in Redis")
		objectResponse(w, 200, productData)
	} else {
		fmt.Println("Couldn't save in Redis")
		objectResponseError(w, 500)
	}
}

func objectResponse(w http.ResponseWriter, status int, results model.Product) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func objectResponseError(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode("An error occurred while saving the data")
}
