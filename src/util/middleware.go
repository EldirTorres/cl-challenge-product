package util

import (
	"cl.challenge.product/src/service/redis"
	"net/http"
)

func RedisConnection(redis redis.Conn, f func(redis redis.Conn, w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { f(redis, w, r) })
}
