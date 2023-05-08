package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main() {
	//定义ws升级器
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	//定义ws处理函数
	handleWebSocket := func(w http.ResponseWriter, r *http.Request) {
		//升级
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logs.Info("fail", conn)
			return
		}
		defer conn.Close()
		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("failed to read msg", err)
				break
			}
			log.Printf("Received msg: %s\n", msg)

			err = conn.WriteMessage(msgType, msg)
			if err != nil {
				log.Println("Failed to write msg", err)
				break
			}
		}
	}
	http.HandleFunc("/ws", handleWebSocket)
	//启动http服务器
	log.Println("listening on:8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
