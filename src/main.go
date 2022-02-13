package main

import (
	configRoute "cl.challenge.product/src/config"
	"log"
	"net/http"
	//"cl.challenge.product/src/main/go/util"
)

func main() {
	router := configRoute.NewRouter()
	//redis, _ := config.RedisInit()
	//router.Use(util.RedisConnection(redis))

	server := http.ListenAndServe(configRoute.Config("PORT"), router)
	log.Fatal(server)

}
