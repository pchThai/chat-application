package main

import (
	"flag"
	"chat-application/chat"
)

var (
	port = flag.String("p", ":8080", "set port")
)
func init(){
	flag.Parse()
}

func main() {
	chat.Start(*port)
}
