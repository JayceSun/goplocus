package main


import (
	"fmt"
	"golang.org/x/net/websocket"

	"log"
	"net/http"
)



func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("来自浏览器的的消息： " + reply)
		msg := "（自动回复）:  " + reply
		fmt.Println("服务器的消息： " + msg)

	if err = websocket.Message.Send(ws, msg); err != nil {
		fmt.Println("Can't send")
		break        }    }
 }




func main() {
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe("127.0.0.1:1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	} }





