package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/middle", MiddlewareFunc(http.HandlerFunc(Ping)))
	http.HandleFunc("/svc", Ping)

	err := http.ListenAndServe(":6666", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func MiddlewareFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		next.ServeHTTP(writer, request)
		fmt.Println("net/http middleware")
	})
}

// Ping 定义handle处理函数，只要该 ping 被调用，就会写入 pong
func Ping(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "pong")
}
