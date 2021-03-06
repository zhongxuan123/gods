package main

//import (
//	"github.com/gorilla/websocket"
//	"net/http"
//	log "github.com/golang/glog"
//)

//websocket server

//func CheckOrigin(r *http.Request) bool {
//	// allow all connections by default
//	return true
//}
//
//var upgrader = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//	CheckOrigin:CheckOrigin,
//}
//
//func ServeWebsocket(w http.ResponseWriter, r *http.Request) {
//	conn, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		log.Error("upgrade err:", err)
//		return
//	}
//	conn.SetReadLimit(64*1024)
//	conn.SetPongHandler(func(string) error {
//		log.Info("brower websocket pong...")
//		return nil
//	})
//
//	client := NewClient(conn)
//	client.Run()
//}
//
//func StartWSServer(address string, tls_address string, cert_file string, key_file string) {
//	mux := http.NewServeMux()
//	mux.HandleFunc("/ws", ServeWebsocket)
//
//	if tls_address != "" && cert_file != "" && key_file != "" {
//		go func() {
//			log.Infof("EngineIO Serving TLS at %s...", tls_address)
//			err := http.ListenAndServeTLS(tls_address, cert_file, key_file, mux)
//			if err != nil {
//				log.Fatalf("listen err:%s", err)
//			}
//		}()
//	}
//	err := http.ListenAndServe(address, mux)
//	if err != nil {
//		log.Fatalf("listen err:%s", err)
//	}
//}
//
//
//func ReadWebsocketMessage(conn *websocket.Conn) *Message {
//	messageType, p, err := conn.ReadMessage()
//	if err != nil {
//		log.Info("read websocket err:", err)
//		return nil
//	}
//	if messageType == websocket.BinaryMessage {
//		return ReadBinaryMesage(p)
//	} else {
//		log.Error("invalid websocket message type:", messageType)
//		return nil
//	}
//}
//
//func SendWebsocketBinaryMessage(conn *websocket.Conn, msg *Message) error {
//	w, err := conn.NextWriter(websocket.BinaryMessage)
//	if err != nil {
//		log.Info("get next writer fail")
//		return err
//	}
//	err = SendMessage(w, msg)
//	if err != nil {
//		return err
//	}
//	err = w.Close()
//	return err
//}
//
