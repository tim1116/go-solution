package main

import (
	"fmt"
	"log"
	"message/global"
	"message/model"
	"message/route"
	"net/http"
)

func init() {
	global.InitDb()
	global.Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(&model.Message{})
}

func main() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

	http.HandleFunc("/", route.MessageList)
	http.HandleFunc("/create", route.MessageCreate)
	http.HandleFunc("/del", route.MessageDel)

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "hello")
	})
	// 默认使用 DefaultServeMux
	log.Fatal(http.ListenAndServe("127.0.0.1:8085", nil))
}
